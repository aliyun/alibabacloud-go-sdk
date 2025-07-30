// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteElasticRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteElasticRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteElasticRuleResponseBody
	GetSuccess() *bool
}

type DeleteElasticRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4773E4EC-025D-509F-AEA9-D53123FDFB0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the deletion is successful.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteElasticRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteElasticRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteElasticRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteElasticRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteElasticRuleResponseBody) SetRequestId(v string) *DeleteElasticRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteElasticRuleResponseBody) SetSuccess(v bool) *DeleteElasticRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteElasticRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
