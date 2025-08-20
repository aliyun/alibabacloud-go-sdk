// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSparkAppLogRootPathRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *SetSparkAppLogRootPathRequest
	GetDBClusterId() *string
	SetOssLogPath(v string) *SetSparkAppLogRootPathRequest
	GetOssLogPath() *string
	SetUseDefaultOss(v bool) *SetSparkAppLogRootPathRequest
	GetUseDefaultOss() *bool
}

type SetSparkAppLogRootPathRequest struct {
	// The ID of the AnalyticDB for MySQL cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-dbclusterid
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The path of Object Storage Service (OSS) logs.
	//
	// example:
	//
	// oss://path/to/log
	OssLogPath *string `json:"OssLogPath,omitempty" xml:"OssLogPath,omitempty"`
	// Specifies whether to use the default OSS log path.
	//
	// example:
	//
	// true
	UseDefaultOss *bool `json:"UseDefaultOss,omitempty" xml:"UseDefaultOss,omitempty"`
}

func (s SetSparkAppLogRootPathRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSparkAppLogRootPathRequest) GoString() string {
	return s.String()
}

func (s *SetSparkAppLogRootPathRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *SetSparkAppLogRootPathRequest) GetOssLogPath() *string {
	return s.OssLogPath
}

func (s *SetSparkAppLogRootPathRequest) GetUseDefaultOss() *bool {
	return s.UseDefaultOss
}

func (s *SetSparkAppLogRootPathRequest) SetDBClusterId(v string) *SetSparkAppLogRootPathRequest {
	s.DBClusterId = &v
	return s
}

func (s *SetSparkAppLogRootPathRequest) SetOssLogPath(v string) *SetSparkAppLogRootPathRequest {
	s.OssLogPath = &v
	return s
}

func (s *SetSparkAppLogRootPathRequest) SetUseDefaultOss(v bool) *SetSparkAppLogRootPathRequest {
	s.UseDefaultOss = &v
	return s
}

func (s *SetSparkAppLogRootPathRequest) Validate() error {
	return dara.Validate(s)
}
