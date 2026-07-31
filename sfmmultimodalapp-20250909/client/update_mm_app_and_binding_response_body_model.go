// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppAndBindingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMmAppAndBindingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMmAppAndBindingResponseBody
	GetSuccess() *bool
}

type UpdateMmAppAndBindingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMmAppAndBindingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMmAppAndBindingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMmAppAndBindingResponseBody) SetRequestId(v string) *UpdateMmAppAndBindingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMmAppAndBindingResponseBody) SetSuccess(v bool) *UpdateMmAppAndBindingResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMmAppAndBindingResponseBody) Validate() error {
	return dara.Validate(s)
}
