// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridCloudClusterRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterRuleResourceId(v string) *CreateHybridCloudClusterRuleResponseBody
	GetClusterRuleResourceId() *string
	SetRequestId(v string) *CreateHybridCloudClusterRuleResponseBody
	GetRequestId() *string
}

type CreateHybridCloudClusterRuleResponseBody struct {
	// example:
	//
	// hdbc-clusterrule-*******m0w
	ClusterRuleResourceId *string `json:"ClusterRuleResourceId,omitempty" xml:"ClusterRuleResourceId,omitempty"`
	// example:
	//
	// 66A98669-CC6E-4F3E-*****-3014697B11AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHybridCloudClusterRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridCloudClusterRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHybridCloudClusterRuleResponseBody) GetClusterRuleResourceId() *string {
	return s.ClusterRuleResourceId
}

func (s *CreateHybridCloudClusterRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHybridCloudClusterRuleResponseBody) SetClusterRuleResourceId(v string) *CreateHybridCloudClusterRuleResponseBody {
	s.ClusterRuleResourceId = &v
	return s
}

func (s *CreateHybridCloudClusterRuleResponseBody) SetRequestId(v string) *CreateHybridCloudClusterRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHybridCloudClusterRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
