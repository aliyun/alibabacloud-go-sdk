// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomEntityId(v string) *CreateCustomEntityResponseBody
	GetCustomEntityId() *string
	SetRequestId(v string) *CreateCustomEntityResponseBody
	GetRequestId() *string
}

type CreateCustomEntityResponseBody struct {
	// example:
	//
	// 1
	CustomEntityId *string `json:"CustomEntityId,omitempty" xml:"CustomEntityId,omitempty"`
	// example:
	//
	// 580e8ce3-3b80-44c5-9f3f-36ac3cc5bdd5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomEntityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomEntityResponseBody) GetCustomEntityId() *string {
	return s.CustomEntityId
}

func (s *CreateCustomEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomEntityResponseBody) SetCustomEntityId(v string) *CreateCustomEntityResponseBody {
	s.CustomEntityId = &v
	return s
}

func (s *CreateCustomEntityResponseBody) SetRequestId(v string) *CreateCustomEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
