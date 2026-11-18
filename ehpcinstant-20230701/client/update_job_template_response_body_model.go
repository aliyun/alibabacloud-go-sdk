// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateJobTemplateResponseBody
	GetRequestId() *string
}

type UpdateJobTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateJobTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateJobTemplateResponseBody) SetRequestId(v string) *UpdateJobTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateJobTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
