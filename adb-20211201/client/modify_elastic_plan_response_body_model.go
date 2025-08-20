// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyElasticPlanResponseBody
	GetRequestId() *string
}

type ModifyElasticPlanResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A5C433C2-001F-58E3-99F5-3274C14DF8BD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyElasticPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyElasticPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyElasticPlanResponseBody) SetRequestId(v string) *ModifyElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyElasticPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
