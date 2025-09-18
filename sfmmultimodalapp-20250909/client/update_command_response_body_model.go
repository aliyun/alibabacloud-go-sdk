// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCommandResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCommandResponseBody
	GetSuccess() *bool
}

type UpdateCommandResponseBody struct {
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCommandResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCommandResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCommandResponseBody) SetRequestId(v string) *UpdateCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCommandResponseBody) SetSuccess(v bool) *UpdateCommandResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
