// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppMetadataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimeMs(v int64) *QueryAppMetadataRequest
	GetEndTimeMs() *int64
	SetMetaIds(v string) *QueryAppMetadataRequest
	GetMetaIds() *string
	SetMetaType(v string) *QueryAppMetadataRequest
	GetMetaType() *string
	SetPid(v string) *QueryAppMetadataRequest
	GetPid() *string
	SetRegionId(v string) *QueryAppMetadataRequest
	GetRegionId() *string
	SetStartTimeMs(v int64) *QueryAppMetadataRequest
	GetStartTimeMs() *int64
}

type QueryAppMetadataRequest struct {
	EndTimeMs *int64 `json:"EndTimeMs,omitempty" xml:"EndTimeMs,omitempty"`
	// The metadata IDs. Use a comma (,) to separate multiple IDs.
	//
	// You can obtain the exception ID on the **exception analysis*	- page of the target application in the ARMS console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4c9dd447,3c76c565
	MetaIds *string `json:"MetaIds,omitempty" xml:"MetaIds,omitempty"`
	// The type of the metadata. Valid values:
	//
	// - sql: Retrieves the SQL statement by sqlId.
	//
	// - exception: Retrieves the exception stack by exceptionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// sql
	MetaType *string `json:"MetaType,omitempty" xml:"MetaType,omitempty"`
	// The application ID. To obtain the ID, call the **ListTraceApps*	- operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ggxw4lnjuz@54364d85b******
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTimeMs *int64  `json:"StartTimeMs,omitempty" xml:"StartTimeMs,omitempty"`
}

func (s QueryAppMetadataRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAppMetadataRequest) GoString() string {
	return s.String()
}

func (s *QueryAppMetadataRequest) GetEndTimeMs() *int64 {
	return s.EndTimeMs
}

func (s *QueryAppMetadataRequest) GetMetaIds() *string {
	return s.MetaIds
}

func (s *QueryAppMetadataRequest) GetMetaType() *string {
	return s.MetaType
}

func (s *QueryAppMetadataRequest) GetPid() *string {
	return s.Pid
}

func (s *QueryAppMetadataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryAppMetadataRequest) GetStartTimeMs() *int64 {
	return s.StartTimeMs
}

func (s *QueryAppMetadataRequest) SetEndTimeMs(v int64) *QueryAppMetadataRequest {
	s.EndTimeMs = &v
	return s
}

func (s *QueryAppMetadataRequest) SetMetaIds(v string) *QueryAppMetadataRequest {
	s.MetaIds = &v
	return s
}

func (s *QueryAppMetadataRequest) SetMetaType(v string) *QueryAppMetadataRequest {
	s.MetaType = &v
	return s
}

func (s *QueryAppMetadataRequest) SetPid(v string) *QueryAppMetadataRequest {
	s.Pid = &v
	return s
}

func (s *QueryAppMetadataRequest) SetRegionId(v string) *QueryAppMetadataRequest {
	s.RegionId = &v
	return s
}

func (s *QueryAppMetadataRequest) SetStartTimeMs(v int64) *QueryAppMetadataRequest {
	s.StartTimeMs = &v
	return s
}

func (s *QueryAppMetadataRequest) Validate() error {
	return dara.Validate(s)
}
