// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableElasticPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableElasticPlanResponseBody
	GetRequestId() *string
}

type DisableElasticPlanResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A5C433C2-001F-58E3-99F5-3274C14DF8BD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableElasticPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DisableElasticPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableElasticPlanResponseBody) SetRequestId(v string) *DisableElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableElasticPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
