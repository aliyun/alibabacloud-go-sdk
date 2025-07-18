// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdpDepartmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateIdpDepartmentResponseBody
	GetData() *string
	SetRequestId(v string) *CreateIdpDepartmentResponseBody
	GetRequestId() *string
}

type CreateIdpDepartmentResponseBody struct {
	// example:
	//
	// 726
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// D6707286-A50E-57B1-B2CF-EFAC59E850D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIdpDepartmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIdpDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIdpDepartmentResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateIdpDepartmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIdpDepartmentResponseBody) SetData(v string) *CreateIdpDepartmentResponseBody {
	s.Data = &v
	return s
}

func (s *CreateIdpDepartmentResponseBody) SetRequestId(v string) *CreateIdpDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIdpDepartmentResponseBody) Validate() error {
	return dara.Validate(s)
}
