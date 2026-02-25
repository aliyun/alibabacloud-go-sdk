// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateParameterSetAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateParameterSetAttributeResponseBody
	GetRequestId() *string
}

type UpdateParameterSetAttributeResponseBody struct {
	// example:
	//
	// 81CF7E18-D318-5670-9A4D-C08476BC4899
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateParameterSetAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateParameterSetAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateParameterSetAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateParameterSetAttributeResponseBody) SetRequestId(v string) *UpdateParameterSetAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateParameterSetAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
