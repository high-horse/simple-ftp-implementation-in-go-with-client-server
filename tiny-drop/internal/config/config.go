package config

import "time"


const (
	// TempDir  = "storage/temp"
	// FinalDir = "storage/uploads"
	// DBPath = "storage/files.db"

	TempDir  = "/home/opc/tinydrop/storage/temp"
    FinalDir = "/home/opc/tinydrop/storage/uploads"
    DBPath  = "/home/opc/tinydrop/storage/files.db"

	MinFreeSpace uint64 = 10 * 1024 *1024 * 1024 // 10 GIGS

	CleanupTime = 3 * time.Minute
	DeleteAfterTIme = 60 * time.Minute
	UploadsTable = "uploads"
	Port = ":9090"
)