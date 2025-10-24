// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomGroupId(v string) *CreateCustomGroupResponseBody
	GetCustomGroupId() *string
	SetRequestId(v string) *CreateCustomGroupResponseBody
	GetRequestId() *string
}

type CreateCustomGroupResponseBody struct {
	// example:
	//
	// 129****
	CustomGroupId *string `json:"CustomGroupId,omitempty" xml:"CustomGroupId,omitempty"`
	// example:
	//
	// 580e8ce3-3b80-44c5-9f3f-36ac3cc5bdd5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomGroupResponseBody) GetCustomGroupId() *string {
	return s.CustomGroupId
}

func (s *CreateCustomGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomGroupResponseBody) SetCustomGroupId(v string) *CreateCustomGroupResponseBody {
	s.CustomGroupId = &v
	return s
}

func (s *CreateCustomGroupResponseBody) SetRequestId(v string) *CreateCustomGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
