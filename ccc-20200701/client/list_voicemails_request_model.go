// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoicemailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaller(v string) *ListVoicemailsRequest
	GetCaller() *string
	SetContactId(v string) *ListVoicemailsRequest
	GetContactId() *string
	SetEndTime(v int64) *ListVoicemailsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListVoicemailsRequest
	GetInstanceId() *string
	SetName(v string) *ListVoicemailsRequest
	GetName() *string
	SetPageNumber(v int32) *ListVoicemailsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVoicemailsRequest
	GetPageSize() *int32
	SetStartTime(v int64) *ListVoicemailsRequest
	GetStartTime() *int64
}

type ListVoicemailsRequest struct {
	// example:
	//
	// 073xxxx7539
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// example:
	//
	// job-125152394144124921
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// example:
	//
	// 1532707199000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Deprecated
	//
	// example:
	//
	// voicemail-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// example:
	//
	// 1532448000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListVoicemailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVoicemailsRequest) GoString() string {
	return s.String()
}

func (s *ListVoicemailsRequest) GetCaller() *string {
	return s.Caller
}

func (s *ListVoicemailsRequest) GetContactId() *string {
	return s.ContactId
}

func (s *ListVoicemailsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListVoicemailsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVoicemailsRequest) GetName() *string {
	return s.Name
}

func (s *ListVoicemailsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVoicemailsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVoicemailsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListVoicemailsRequest) SetCaller(v string) *ListVoicemailsRequest {
	s.Caller = &v
	return s
}

func (s *ListVoicemailsRequest) SetContactId(v string) *ListVoicemailsRequest {
	s.ContactId = &v
	return s
}

func (s *ListVoicemailsRequest) SetEndTime(v int64) *ListVoicemailsRequest {
	s.EndTime = &v
	return s
}

func (s *ListVoicemailsRequest) SetInstanceId(v string) *ListVoicemailsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVoicemailsRequest) SetName(v string) *ListVoicemailsRequest {
	s.Name = &v
	return s
}

func (s *ListVoicemailsRequest) SetPageNumber(v int32) *ListVoicemailsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVoicemailsRequest) SetPageSize(v int32) *ListVoicemailsRequest {
	s.PageSize = &v
	return s
}

func (s *ListVoicemailsRequest) SetStartTime(v int64) *ListVoicemailsRequest {
	s.StartTime = &v
	return s
}

func (s *ListVoicemailsRequest) Validate() error {
	return dara.Validate(s)
}
