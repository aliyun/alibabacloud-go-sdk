// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDigitalEmployeeSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDigitalEmployeeSkillResponseBody
	GetRequestId() *string
}

type DeleteDigitalEmployeeSkillResponseBody struct {
	// example:
	//
	// 0CEC5375-XXXX-XXXX-XXXX-9A629907C1F0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteDigitalEmployeeSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDigitalEmployeeSkillResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDigitalEmployeeSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDigitalEmployeeSkillResponseBody) SetRequestId(v string) *DeleteDigitalEmployeeSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDigitalEmployeeSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
