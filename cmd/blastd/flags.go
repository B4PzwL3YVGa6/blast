// Copyright (c) 2018 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"

	"github.com/mosuka/blast/index"
	"github.com/mosuka/blast/raft"
	"github.com/mosuka/blast/store"
	"github.com/urfave/cli"
)

var (
	flRaftAddr = cli.StringFlag{
		Name:   "raft-addr",
		Value:  ":10000",
		Usage:  "Raft address to listen on for peer traffic",
		EnvVar: "BLAST_RAFT_ADDR",
	}
	flGRPCAddr = cli.StringFlag{
		Name:   "grpc-addr",
		Value:  ":10001",
		Usage:  "gRPC address to listen on for client traffic",
		EnvVar: "BLAST_GRPC_ADDR",
	}
	flHTTPAddr = cli.StringFlag{
		Name:   "http-addr",
		Value:  ":10002",
		Usage:  "HTTP address to listen on for client traffic",
		EnvVar: "BLAST_HTTP_ADDR",
	}

	flRaftNodeID = cli.StringFlag{
		Name:   "raft-node-id",
		Value:  raft.DefaultNodeID,
		Usage:  "Raft node ID",
		EnvVar: "BLAST_RAFT_NODE_ID",
	}
	flRaftDir = cli.StringFlag{
		Name:   "raft-dir",
		Value:  raft.DefaultDir,
		Usage:  "Raft data directory",
		EnvVar: "BLAST_RAFT_DIR",
	}
	flRaftSnapshotCount = cli.IntFlag{
		Name:   "raft-snapshot-count",
		Value:  raft.DefaultSnapshotCount,
		Usage:  "Raft snapshot count",
		EnvVar: "BLAST_RAFT_SNAPSHOT_COUNT",
	}
	flRaftTimeout = cli.StringFlag{
		Name:   "raft-timeout",
		Value:  raft.DefaultTimeout,
		Usage:  "Raft timeout",
		EnvVar: "BLAST_RAFT_TIMEOUT",
	}

	flStoreDir = cli.StringFlag{
		Name:   "store-dir",
		Value:  store.DefaultDir,
		Usage:  "Store data directory",
		EnvVar: "BLAST_STORE_DIR",
	}

	flIndexDir = cli.StringFlag{
		Name:   "index-dir",
		Value:  index.DefaultDir,
		Usage:  "Index data directory",
		EnvVar: "BLAST_INDEX_DIR",
	}
	flIndexMappingFile = cli.StringFlag{
		Name:   "index-mapping-file",
		Value:  index.DefaultIndexMappingFile,
		Usage:  "Index mapping file",
		EnvVar: "BLAST_INDEX_MAPPING_FILE",
	}
	flIndexType = cli.StringFlag{
		Name:   "index-type",
		Value:  index.DefaultIndexType,
		Usage:  "Index type",
		EnvVar: "BLAST_INDEX_TYPE",
	}
	flIndexKvstore = cli.StringFlag{
		Name:   "index-kvstore",
		Value:  index.DefaultKvstore,
		Usage:  "Index Key-Value store",
		EnvVar: "BLAST_INDEX_KVSTORE",
	}

	flPeerGRPCAddr = cli.StringFlag{
		Name:   "peer-grpc-addr",
		Usage:  "Peer gRPC address to connect on for join the cluster",
		EnvVar: "BLAST_PEER_GRPC_ADDR",
	}

	flLogLevel = cli.StringFlag{
		Name:   "log-level",
		Value:  "INFO",
		Usage:  "Log level",
		EnvVar: "BLAST_LOG_LEVEL",
	}
	flLogFile = cli.StringFlag{
		Name:   "log-file",
		Value:  os.Stdout.Name(),
		Usage:  "Log file",
		EnvVar: "BLAST_LOG_FILE",
	}
	flLogMaxSize = cli.IntFlag{
		Name:   "log-max-size",
		Value:  500,
		Usage:  "Max size of a log file (megabytes)",
		EnvVar: "BLAST_LOG_MAX_SIZE",
	}
	flLogMaxBackups = cli.IntFlag{
		Name:   "log-max-backups",
		Value:  3,
		Usage:  "Max backup count of log files",
		EnvVar: "BLAST_LOG_MAX_BACKUPS",
	}
	flLogMaxAge = cli.IntFlag{
		Name:   "log-max-age",
		Value:  30,
		Usage:  "Max age of a log file (days)",
		EnvVar: "BLAST_LOG_MAX_AGE",
	}
	flLogCompress = cli.BoolFlag{
		Name:   "log-compress",
		Usage:  "Compress a log file",
		EnvVar: "BLAST_LOG_COMPRESS",
	}

	flHTTPAccessLogFile = cli.StringFlag{
		Name:   "http-access-log-file",
		Value:  os.Stdout.Name(),
		Usage:  "HTTP access log file",
		EnvVar: "BLAST_HTTP_ACCESS_LOG_FILE",
	}
	flHTTPAccessLogMaxSize = cli.IntFlag{
		Name:   "http-access-log-max-size",
		Value:  500,
		Usage:  "Max size of a HTTP access log file (megabytes)",
		EnvVar: "BLAST_HTTP_ACCESS_LOG_MAX_SIZE",
	}
	flHTTPAccessLogMaxBackups = cli.IntFlag{
		Name:   "http-access-log-max-backups",
		Value:  3,
		Usage:  "Max backup count of HTTP access log files",
		EnvVar: "BLAST_HTTP_ACCESS_LOG_MAX_BACKUPS",
	}
	flHTTPAccessLogMaxAge = cli.IntFlag{
		Name:   "http-access-log-max-age",
		Value:  30,
		Usage:  "Max age of a HTTP access log file (days)",
		EnvVar: "BLAST_HTTP_ACCESS_LOG_MAX_AGE",
	}
	flHTTPAccessLogCompress = cli.BoolFlag{
		Name:   "http-access-log-compress",
		Usage:  "Compress a HTTP access log",
		EnvVar: "BLAST_HTTP_ACCESS_LOG_COMPRESS",
	}
)
