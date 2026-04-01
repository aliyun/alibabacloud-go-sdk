// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyParameterResponseBody
	GetRequestId() *string
}

type ModifyParameterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 542BB8D6-4268-45CC-A557-B03EFD7AB30A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyParameterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyParameterResponseBody) SetRequestId(v string) *ModifyParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyParameterResponseBody) Validate() error {
	return dara.Validate(s)
}
