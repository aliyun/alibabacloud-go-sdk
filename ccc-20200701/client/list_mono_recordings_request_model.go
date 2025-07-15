// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMonoRecordingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *ListMonoRecordingsRequest
	GetContactId() *string
	SetInstanceId(v string) *ListMonoRecordingsRequest
	GetInstanceId() *string
}

type ListMonoRecordingsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// job-25697383427137****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListMonoRecordingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMonoRecordingsRequest) GoString() string {
	return s.String()
}

func (s *ListMonoRecordingsRequest) GetContactId() *string {
	return s.ContactId
}

func (s *ListMonoRecordingsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListMonoRecordingsRequest) SetContactId(v string) *ListMonoRecordingsRequest {
	s.ContactId = &v
	return s
}

func (s *ListMonoRecordingsRequest) SetInstanceId(v string) *ListMonoRecordingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListMonoRecordingsRequest) Validate() error {
	return dara.Validate(s)
}
