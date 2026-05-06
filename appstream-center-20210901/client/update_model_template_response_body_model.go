// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateModelTemplateResponseBody
	GetRequestId() *string
}

type UpdateModelTemplateResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateModelTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateModelTemplateResponseBody) SetRequestId(v string) *UpdateModelTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateModelTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
