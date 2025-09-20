// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTriggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateTriggerResponseBody
	GetId() *string
	SetRequestId(v string) *CreateTriggerResponseBody
	GetRequestId() *string
}

type CreateTriggerResponseBody struct {
	// The ID of the trigger.
	//
	// example:
	//
	// trigger-9f72636a-0f0c-4baf-ae78-38b27b******
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EC564A9A-BA5C-4499-A087-D9B9E76E*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTriggerResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateTriggerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTriggerResponseBody) SetId(v string) *CreateTriggerResponseBody {
	s.Id = &v
	return s
}

func (s *CreateTriggerResponseBody) SetRequestId(v string) *CreateTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTriggerResponseBody) Validate() error {
	return dara.Validate(s)
}
