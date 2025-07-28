// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudClusterRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHybridCloudClusterRuleResponseBody
	GetRequestId() *string
}

type ModifyHybridCloudClusterRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 03D73D88-57D8-5BA2-96A4-6357CE***19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHybridCloudClusterRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudClusterRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudClusterRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHybridCloudClusterRuleResponseBody) SetRequestId(v string) *ModifyHybridCloudClusterRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHybridCloudClusterRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
