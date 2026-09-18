// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateJobPlanResponseBody
	GetRequestId() *string
}

type UpdateJobPlanResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 40325405-579C-4D82-9624-EC2B1779848E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateJobPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJobPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateJobPlanResponseBody) SetRequestId(v string) *UpdateJobPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateJobPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
