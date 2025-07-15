// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogInfos(v *DescribeLogConfigResponseBodyLogInfos) *DescribeLogConfigResponseBody
	GetLogInfos() *DescribeLogConfigResponseBodyLogInfos
	SetRequestId(v string) *DescribeLogConfigResponseBody
	GetRequestId() *string
}

type DescribeLogConfigResponseBody struct {
	// Info of the log config.
	LogInfos *DescribeLogConfigResponseBodyLogInfos `json:"LogInfos,omitempty" xml:"LogInfos,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E3BC2706-ABDB-5B64-A12F-08DFD9E3F339
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogConfigResponseBody) GetLogInfos() *DescribeLogConfigResponseBodyLogInfos {
	return s.LogInfos
}

func (s *DescribeLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogConfigResponseBody) SetLogInfos(v *DescribeLogConfigResponseBodyLogInfos) *DescribeLogConfigResponseBody {
	s.LogInfos = v
	return s
}

func (s *DescribeLogConfigResponseBody) SetRequestId(v string) *DescribeLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLogConfigResponseBodyLogInfos struct {
	LogInfo []*DescribeLogConfigResponseBodyLogInfosLogInfo `json:"LogInfo,omitempty" xml:"LogInfo,omitempty" type:"Repeated"`
}

func (s DescribeLogConfigResponseBodyLogInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogConfigResponseBodyLogInfos) GoString() string {
	return s.String()
}

func (s *DescribeLogConfigResponseBodyLogInfos) GetLogInfo() []*DescribeLogConfigResponseBodyLogInfosLogInfo {
	return s.LogInfo
}

func (s *DescribeLogConfigResponseBodyLogInfos) SetLogInfo(v []*DescribeLogConfigResponseBodyLogInfosLogInfo) *DescribeLogConfigResponseBodyLogInfos {
	s.LogInfo = v
	return s
}

func (s *DescribeLogConfigResponseBodyLogInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeLogConfigResponseBodyLogInfosLogInfo struct {
	// The log type.
	//
	// example:
	//
	// PROVIDER
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The region ID of the Logstore.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the Logstore in Log Service.
	//
	// example:
	//
	// slsstore
	SlsLogStore *string `json:"SlsLogStore,omitempty" xml:"SlsLogStore,omitempty"`
	// The name of the Log Service project.
	//
	// example:
	//
	// slsproject
	SlsProject *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
}

func (s DescribeLogConfigResponseBodyLogInfosLogInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogConfigResponseBodyLogInfosLogInfo) GoString() string {
	return s.String()
}

func (s *DescribeLogConfigResponseBodyLogInfosLogInfo) GetLogType() *string {
	return s.LogType
}

func (s *DescribeLogConfigResponseBodyLogInfosLogInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLogConfigResponseBodyLogInfosLogInfo) GetSlsLogStore() *string {
	return s.SlsLogStore
}

func (s *DescribeLogConfigResponseBodyLogInfosLogInfo) GetSlsProject() *string {
	return s.SlsProject
}

func (s *DescribeLogConfigResponseBodyLogInfosLogInfo) SetLogType(v string) *DescribeLogConfigResponseBodyLogInfosLogInfo {
	s.LogType = &v
	return s
}

func (s *DescribeLogConfigResponseBodyLogInfosLogInfo) SetRegionId(v string) *DescribeLogConfigResponseBodyLogInfosLogInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeLogConfigResponseBodyLogInfosLogInfo) SetSlsLogStore(v string) *DescribeLogConfigResponseBodyLogInfosLogInfo {
	s.SlsLogStore = &v
	return s
}

func (s *DescribeLogConfigResponseBodyLogInfosLogInfo) SetSlsProject(v string) *DescribeLogConfigResponseBodyLogInfosLogInfo {
	s.SlsProject = &v
	return s
}

func (s *DescribeLogConfigResponseBodyLogInfosLogInfo) Validate() error {
	return dara.Validate(s)
}
