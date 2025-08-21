// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ModifyImageAttributeResponseBody
	GetCode() *int32
	SetRequestId(v string) *ModifyImageAttributeResponseBody
	GetRequestId() *string
}

type ModifyImageAttributeResponseBody struct {
	// The service code. 0 is returned for a successful request. An error code is returned for a failed request.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AC66B8F3-0B0A-5FB1-9EA2-DC03B2CD5B04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyImageAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ModifyImageAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyImageAttributeResponseBody) SetCode(v int32) *ModifyImageAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyImageAttributeResponseBody) SetRequestId(v string) *ModifyImageAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyImageAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
