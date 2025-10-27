// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteElasticPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteElasticPlanResponseBody
	GetRequestId() *string
}

type DeleteElasticPlanResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteElasticPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteElasticPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteElasticPlanResponseBody) SetRequestId(v string) *DeleteElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteElasticPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
