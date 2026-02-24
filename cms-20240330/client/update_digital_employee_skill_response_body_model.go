// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDigitalEmployeeSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDigitalEmployeeSkillResponseBody
	GetRequestId() *string
}

type UpdateDigitalEmployeeSkillResponseBody struct {
	// example:
	//
	// 350779DC-980D-58FD-BECB-D2275D2487CA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateDigitalEmployeeSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDigitalEmployeeSkillResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDigitalEmployeeSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDigitalEmployeeSkillResponseBody) SetRequestId(v string) *UpdateDigitalEmployeeSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDigitalEmployeeSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
