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
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
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
