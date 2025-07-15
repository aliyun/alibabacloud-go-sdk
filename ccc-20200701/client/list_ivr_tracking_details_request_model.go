// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIvrTrackingDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *ListIvrTrackingDetailsRequest
	GetContactId() *string
	SetInstanceId(v string) *ListIvrTrackingDetailsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListIvrTrackingDetailsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListIvrTrackingDetailsRequest
	GetPageSize() *int32
}

type ListIvrTrackingDetailsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// job-10963442671187****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListIvrTrackingDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIvrTrackingDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListIvrTrackingDetailsRequest) GetContactId() *string {
	return s.ContactId
}

func (s *ListIvrTrackingDetailsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListIvrTrackingDetailsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIvrTrackingDetailsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIvrTrackingDetailsRequest) SetContactId(v string) *ListIvrTrackingDetailsRequest {
	s.ContactId = &v
	return s
}

func (s *ListIvrTrackingDetailsRequest) SetInstanceId(v string) *ListIvrTrackingDetailsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIvrTrackingDetailsRequest) SetPageNumber(v int32) *ListIvrTrackingDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIvrTrackingDetailsRequest) SetPageSize(v int32) *ListIvrTrackingDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *ListIvrTrackingDetailsRequest) Validate() error {
	return dara.Validate(s)
}
