// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteJobPlanResponseBody
	GetRequestId() *string
}

type DeleteJobPlanResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82-9624-EC2B1779848E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteJobPlanResponseBody) SetRequestId(v string) *DeleteJobPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteJobPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
