// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridCloudClusterRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHybridCloudClusterRuleResponseBody
	GetRequestId() *string
}

type DeleteHybridCloudClusterRuleResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHybridCloudClusterRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridCloudClusterRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHybridCloudClusterRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHybridCloudClusterRuleResponseBody) SetRequestId(v string) *DeleteHybridCloudClusterRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHybridCloudClusterRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
