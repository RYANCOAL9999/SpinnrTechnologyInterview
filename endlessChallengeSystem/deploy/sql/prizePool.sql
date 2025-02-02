-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

-- MySQL Script generated by MySQL Workbench
-- Sat Jul  27 16:09:21 2024
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering;

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema SpinnrTechnology
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema SpinnrTechnology
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `SpinnrTechnology` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
USE `SpinnrTechnology` ;

-- -----------------------------------------------------
-- Table `SpinnrTechnology`.`PrizePool`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `SpinnrTechnology`.`PrizePool` (
    `ID` INT AUTO_INCREMENT PRIMARY KEY,
    `Amount` FLOAT NOT NULL)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

-- -----------------------------------------------------
-- Insert the single record into PrizePool
-- -----------------------------------------------------
INSERT INTO PrizePool (amount)
SELECT 0
WHERE NOT EXISTS (SELECT 1 FROM PrizePool);

-- -----------------------------------------------------
-- Create a trigger to prevent additional inserts
-- -----------------------------------------------------
DELIMITER //
CREATE TRIGGER prevent_prizepool_insert
BEFORE INSERT ON PrizePool
FOR EACH ROW
BEGIN
    DECLARE record_count INT;
    SELECT COUNT(*) INTO record_count FROM PrizePool;
    IF record_count > 0 THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Only one record is allowed in the PrizePool table';
    END IF;
END;//
DELIMITER ;

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

-- -----------------------------------------------------
-- Table `SpinnrTechnology`.`PrizePool`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `SpinnrTechnology`.`PrizePool` ;
-- -----------------------------------------------------
-- Schema SpinnrTechnology
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `SpinnrTechnology` ;