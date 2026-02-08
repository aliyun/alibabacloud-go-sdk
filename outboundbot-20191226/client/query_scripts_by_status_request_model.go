// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScriptsByStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryScriptsByStatusRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *QueryScriptsByStatusRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryScriptsByStatusRequest
	GetPageSize() *int32
	SetStatusList(v []*string) *QueryScriptsByStatusRequest
	GetStatusList() []*string
}

type QueryScriptsByStatusRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 361c8a53-0e29-42f3-8aa7-c7752d010399
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Page number
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size for paging
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// List of statuses
	//
	// This parameter is required.
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s QueryScriptsByStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryScriptsByStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryScriptsByStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryScriptsByStatusRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryScriptsByStatusRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryScriptsByStatusRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *QueryScriptsByStatusRequest) SetInstanceId(v string) *QueryScriptsByStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryScriptsByStatusRequest) SetPageNumber(v int32) *QueryScriptsByStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryScriptsByStatusRequest) SetPageSize(v int32) *QueryScriptsByStatusRequest {
	s.PageSize = &v
	return s
}

func (s *QueryScriptsByStatusRequest) SetStatusList(v []*string) *QueryScriptsByStatusRequest {
	s.StatusList = v
	return s
}

func (s *QueryScriptsByStatusRequest) Validate() error {
	return dara.Validate(s)
}
