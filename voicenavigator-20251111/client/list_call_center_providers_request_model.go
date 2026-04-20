// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallCenterProvidersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListCallCenterProvidersRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListCallCenterProvidersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCallCenterProvidersRequest
	GetPageSize() *int32
	SetProviderId(v string) *ListCallCenterProvidersRequest
	GetProviderId() *string
}

type ListCallCenterProvidersRequest struct {
	// example:
	//
	// 36e9a4cd-a821-4f78-86fa-a9a4aefeea2e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xxxxxxxxx
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
}

func (s ListCallCenterProvidersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCallCenterProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListCallCenterProvidersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCallCenterProvidersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCallCenterProvidersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCallCenterProvidersRequest) GetProviderId() *string {
	return s.ProviderId
}

func (s *ListCallCenterProvidersRequest) SetInstanceId(v string) *ListCallCenterProvidersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCallCenterProvidersRequest) SetPageNumber(v int32) *ListCallCenterProvidersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCallCenterProvidersRequest) SetPageSize(v int32) *ListCallCenterProvidersRequest {
	s.PageSize = &v
	return s
}

func (s *ListCallCenterProvidersRequest) SetProviderId(v string) *ListCallCenterProvidersRequest {
	s.ProviderId = &v
	return s
}

func (s *ListCallCenterProvidersRequest) Validate() error {
	return dara.Validate(s)
}
