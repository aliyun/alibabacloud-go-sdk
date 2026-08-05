// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsyncTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataId(v int32) *ListAsyncTasksRequest
	GetDataId() *int32
	SetDryRun(v bool) *ListAsyncTasksRequest
	GetDryRun() *bool
	SetServiceType(v string) *ListAsyncTasksRequest
	GetServiceType() *string
}

type ListAsyncTasksRequest struct {
	// The trial data ID.
	//
	// example:
	//
	// 1231
	DataId *int32 `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// Specifies whether to validate the request parameters without performing the actual operation. Default value: false.
	//
	// Valid values:
	//
	// - **true**
	//
	// - **false**.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// The service type.
	//
	// - document-analyze.
	//
	// example:
	//
	// document-analyze
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s ListAsyncTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncTasksRequest) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksRequest) GetDataId() *int32 {
	return s.DataId
}

func (s *ListAsyncTasksRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListAsyncTasksRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListAsyncTasksRequest) SetDataId(v int32) *ListAsyncTasksRequest {
	s.DataId = &v
	return s
}

func (s *ListAsyncTasksRequest) SetDryRun(v bool) *ListAsyncTasksRequest {
	s.DryRun = &v
	return s
}

func (s *ListAsyncTasksRequest) SetServiceType(v string) *ListAsyncTasksRequest {
	s.ServiceType = &v
	return s
}

func (s *ListAsyncTasksRequest) Validate() error {
	return dara.Validate(s)
}
