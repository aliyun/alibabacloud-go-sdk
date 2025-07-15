// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallDetailRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *GetCallDetailRecordRequest
	GetContactId() *string
	SetInstanceId(v string) *GetCallDetailRecordRequest
	GetInstanceId() *string
}

type GetCallDetailRecordRequest struct {
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
}

func (s GetCallDetailRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordRequest) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordRequest) GetContactId() *string {
	return s.ContactId
}

func (s *GetCallDetailRecordRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCallDetailRecordRequest) SetContactId(v string) *GetCallDetailRecordRequest {
	s.ContactId = &v
	return s
}

func (s *GetCallDetailRecordRequest) SetInstanceId(v string) *GetCallDetailRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCallDetailRecordRequest) Validate() error {
	return dara.Validate(s)
}
