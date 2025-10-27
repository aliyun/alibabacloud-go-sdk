// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateElasticPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateElasticPlanResponseBody
	GetRequestId() *string
}

type CreateElasticPlanResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateElasticPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateElasticPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateElasticPlanResponseBody) SetRequestId(v string) *CreateElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateElasticPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
