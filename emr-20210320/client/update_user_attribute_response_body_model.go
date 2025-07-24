// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateUserAttributeResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateUserAttributeResponseBody
	GetRequestId() *string
}

type UpdateUserAttributeResponseBody struct {
	// Deprecated
	//
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserAttributeResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateUserAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserAttributeResponseBody) SetData(v bool) *UpdateUserAttributeResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateUserAttributeResponseBody) SetRequestId(v string) *UpdateUserAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
