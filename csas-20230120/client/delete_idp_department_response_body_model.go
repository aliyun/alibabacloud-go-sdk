// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIdpDepartmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIdpDepartmentResponseBody
	GetRequestId() *string
}

type DeleteIdpDepartmentResponseBody struct {
	// example:
	//
	// FEF1144C-95D1-5F7C-81EF-9DB70EA49FCE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIdpDepartmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdpDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIdpDepartmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIdpDepartmentResponseBody) SetRequestId(v string) *DeleteIdpDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIdpDepartmentResponseBody) Validate() error {
	return dara.Validate(s)
}
