// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudbenchTaskConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCloudbenchTaskConfigResponseBody
	GetCode() *string
	SetData(v *DescribeCloudbenchTaskConfigResponseBodyData) *DescribeCloudbenchTaskConfigResponseBody
	GetData() *DescribeCloudbenchTaskConfigResponseBodyData
	SetMessage(v string) *DescribeCloudbenchTaskConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCloudbenchTaskConfigResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeCloudbenchTaskConfigResponseBody
	GetSuccess() *string
}

type DescribeCloudbenchTaskConfigResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information, including the error codes and the number of entries that are returned.
	Data *DescribeCloudbenchTaskConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudbenchTaskConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudbenchTaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudbenchTaskConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCloudbenchTaskConfigResponseBody) GetData() *DescribeCloudbenchTaskConfigResponseBodyData {
	return s.Data
}

func (s *DescribeCloudbenchTaskConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCloudbenchTaskConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudbenchTaskConfigResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeCloudbenchTaskConfigResponseBody) SetCode(v string) *DescribeCloudbenchTaskConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBody) SetData(v *DescribeCloudbenchTaskConfigResponseBodyData) *DescribeCloudbenchTaskConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBody) SetMessage(v string) *DescribeCloudbenchTaskConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBody) SetRequestId(v string) *DescribeCloudbenchTaskConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBody) SetSuccess(v string) *DescribeCloudbenchTaskConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudbenchTaskConfigResponseBodyData struct {
	// The path in which the files are archived.
	//
	// example:
	//
	// /tmp/das/cloudbench/archive-sqls/
	ArchiveFolder *string `json:"ArchiveFolder,omitempty" xml:"ArchiveFolder,omitempty"`
	// The command that was run to start the stress testing task.
	//
	// example:
	//
	// java -jar /tmp/das/cloudbench/CloudBenchClient.jar --bench --rocksdb /tmp/das/cloudbench/rocksdb --meta /tmp/das/cloudbench/cl-1621353601000-360****.meta --task_name 2777bba9-a836-49e6-9f70-1c3822fc9239 --result_file /tmp/das/cloudbench/null.result --user cloudb***	- --pwd \\"cloudbench@****\\" --host rm-bp1j5f8s5x26kq79216****.mysql.rds.aliyuncs.com --port 3306 --charset utf8mb4 --interval 1 --bench_time 3600 --rate_factor 1.0 --start_time 1621353601 --rt > /tmp/das/cloudbench/null.log
	BenchCmd *string `json:"BenchCmd,omitempty" xml:"BenchCmd,omitempty"`
	// The path to the JAR file that is used for stress testing.
	//
	// example:
	//
	// /tmp/das/cloudbench/CloudBenchClient.jar
	ClientJarPath *string `json:"ClientJarPath,omitempty" xml:"ClientJarPath,omitempty"`
	// The path to the JAR file that is stored in OSS. The JAR file is used for stress testing.
	//
	// example:
	//
	// https://cloudbench-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/CloudBenchClient.jar?OSSAccessKeyId=LTAI5tKj8B4wikkVtupK****&Expires=1622441372&Signature=28p%2BCe4tNHpr9VPOcHc3Si9iOb****
	JarOnOss *string `json:"JarOnOss,omitempty" xml:"JarOnOss,omitempty"`
	// The command that was run to preload the file that stores the analysis result of full SQL statistics.
	//
	// example:
	//
	// java -jar /tmp/das/cloudbench/CloudBenchClient.jar --load --out /tmp/das/cloudbench/cl-1621353601000-360****.sc --meta /tmp/das/cloudbench/cl-1621353601000-360****.meta --task_name 2777bba9-****-49e6-9f70-1c3822fc***	- --rocksdb /tmp/das/cloudbench/rocksdb
	LoadCmd *string `json:"LoadCmd,omitempty" xml:"LoadCmd,omitempty"`
	// The name of the metadata file.
	//
	// example:
	//
	// cl-1621353601000-360****.meta
	MetaFileName *string `json:"MetaFileName,omitempty" xml:"MetaFileName,omitempty"`
	// The name of the metadata file stored in Object Storage Service (OSS).
	//
	// example:
	//
	// "https://cb-rm-bp1w9g06h560l****.oss-cn-hangzhou.aliyuncs.com/cl-1621353601000-360****.meta?OSSAccessKeyId=LTAI5tKj8B4wikkVtupK****&Expires=1622441372&Signature=Qsehg3tzeA57M%2BIixAbWPWAtvl****
	MetaFileOnOss *string `json:"MetaFileOnOss,omitempty" xml:"MetaFileOnOss,omitempty"`
	// The path to the metadata file.
	//
	// example:
	//
	// /tmp/das/cloudbench/cl-1621353601000-360****.meta
	MetaFilePath *string `json:"MetaFilePath,omitempty" xml:"MetaFilePath,omitempty"`
	// The command that was run to parse the file that stores the analysis result of full SQL statistics.
	//
	// example:
	//
	// cd /tmp/das/cloudbench && java -jar CloudBenchClient.jar --parse --threads 32 --file /tmp/das/cloudbench/2777bba9-a836-49e6-9f70-1c3822fc9239.archiveSql --meta /tmp/das/cloudbench/cl-1621353601000-360****.meta --out /tmp/das/cloudbench/cl-1621353601000-360****.sc --parent_patmp/das/cloudbench --source RDS --h /thost rm-bp1j5f8s5x266****.mysql.rds.aliyuncs.com --port 3306 --user cloudb***	- --pwd \\"cloudbench@****\\" --cutSqlLen 8192 --db_black_list=information_schema,test,unknow,null
	ParseCmd *string `json:"ParseCmd,omitempty" xml:"ParseCmd,omitempty"`
	// The path to the file that is parsed. The file stores the analysis result of full SQL statistics.
	//
	// example:
	//
	// /tmp/das/cloudbench/2777bba9-a836-49e6-9f70-1c3822fc****.archiveSql
	ParseFilePath *string `json:"ParseFilePath,omitempty" xml:"ParseFilePath,omitempty"`
	// The location where the RocksDB storage system is deployed in the stress testing client.
	//
	// example:
	//
	// /tmp/das/cloudbench/rocksdb
	RocksDbPath *string `json:"RocksDbPath,omitempty" xml:"RocksDbPath,omitempty"`
	// The name of the file that stores the analysis result of full SQL statistics.
	//
	// example:
	//
	// cl-1621353601000-360****.sc
	SqlFileName *string `json:"SqlFileName,omitempty" xml:"SqlFileName,omitempty"`
	// The name of the file that stores the analysis result of full SQL statistics and that is stored in OSS.
	//
	// example:
	//
	// https://cb-rm-bp1w9g06h560l****.oss-cn-hangzhou.aliyuncs.com/cl-1621353601000-360****.sc?OSSAccessKeyId=LTAI5tKj8B4wikkVtupK****&Expires=1622441372&Signature=LYMADwo%2BRrJeqR3e4d8OlIkVmw****
	SqlFileOnOss *string `json:"SqlFileOnOss,omitempty" xml:"SqlFileOnOss,omitempty"`
	// The path to the file that stores the analysis result of full SQL statistics.
	//
	// example:
	//
	// /tmp/das/cloudbench/cl-1621353601000-360****.sc
	SqlFilePath *string `json:"SqlFilePath,omitempty" xml:"SqlFilePath,omitempty"`
	// The task ID.
	//
	// example:
	//
	// e5cec704-0518-430f-8263-76f4dcds****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 1091411816252****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The path of the temporary directory that is generated for stress testing.
	//
	// example:
	//
	// /tmp/bench/
	WorkDir *string `json:"WorkDir,omitempty" xml:"WorkDir,omitempty"`
}

func (s DescribeCloudbenchTaskConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudbenchTaskConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) GetArchiveFolder() *string {
	return s.ArchiveFolder
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) GetBenchCmd() *string {
	return s.BenchCmd
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) GetClientJarPath() *string {
	return s.ClientJarPath
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) GetJarOnOss() *string {
	return s.JarOnOss
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) GetLoadCmd() *string {
	return s.LoadCmd
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) GetMetaFileName() *string {
	return s.MetaFileName
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) GetMetaFileOnOss() *string {
	return s.MetaFileOnOss
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) GetMetaFilePath() *string {
	return s.MetaFilePath
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) GetParseCmd() *string {
	return s.ParseCmd
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) GetParseFilePath() *string {
	return s.ParseFilePath
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) GetRocksDbPath() *string {
	return s.RocksDbPath
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) GetSqlFileName() *string {
	return s.SqlFileName
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) GetSqlFileOnOss() *string {
	return s.SqlFileOnOss
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) GetSqlFilePath() *string {
	return s.SqlFilePath
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) GetWorkDir() *string {
	return s.WorkDir
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetArchiveFolder(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.ArchiveFolder = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetBenchCmd(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.BenchCmd = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetClientJarPath(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.ClientJarPath = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetJarOnOss(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.JarOnOss = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetLoadCmd(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.LoadCmd = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetMetaFileName(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.MetaFileName = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetMetaFileOnOss(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.MetaFileOnOss = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetMetaFilePath(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.MetaFilePath = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetParseCmd(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.ParseCmd = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetParseFilePath(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.ParseFilePath = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetRocksDbPath(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.RocksDbPath = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetSqlFileName(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.SqlFileName = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetSqlFileOnOss(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.SqlFileOnOss = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetSqlFilePath(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.SqlFilePath = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetTaskId(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetUserId(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.UserId = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetWorkDir(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.WorkDir = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
