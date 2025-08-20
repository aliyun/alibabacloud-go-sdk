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
	// A5C433C2-001F-58E3-99F5-3274C14DF8BD
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
