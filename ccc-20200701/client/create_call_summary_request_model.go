// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCallSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *CreateCallSummaryRequest
	GetContactId() *string
	SetContext(v string) *CreateCallSummaryRequest
	GetContext() *string
	SetCustomerId(v string) *CreateCallSummaryRequest
	GetCustomerId() *string
	SetInstanceId(v string) *CreateCallSummaryRequest
	GetInstanceId() *string
}

type CreateCallSummaryRequest struct {
	// example:
	//
	// job-522327189435260928
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	Context   *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// example:
	//
	// 51e155ce-3747-*****-b402-13c69597b920
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateCallSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCallSummaryRequest) GoString() string {
	return s.String()
}

func (s *CreateCallSummaryRequest) GetContactId() *string {
	return s.ContactId
}

func (s *CreateCallSummaryRequest) GetContext() *string {
	return s.Context
}

func (s *CreateCallSummaryRequest) GetCustomerId() *string {
	return s.CustomerId
}

func (s *CreateCallSummaryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCallSummaryRequest) SetContactId(v string) *CreateCallSummaryRequest {
	s.ContactId = &v
	return s
}

func (s *CreateCallSummaryRequest) SetContext(v string) *CreateCallSummaryRequest {
	s.Context = &v
	return s
}

func (s *CreateCallSummaryRequest) SetCustomerId(v string) *CreateCallSummaryRequest {
	s.CustomerId = &v
	return s
}

func (s *CreateCallSummaryRequest) SetInstanceId(v string) *CreateCallSummaryRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCallSummaryRequest) Validate() error {
	return dara.Validate(s)
}
