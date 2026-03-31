// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudClusterRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeHybridCloudClusterRulesResponseBodyData) *DescribeHybridCloudClusterRulesResponseBody
	GetData() []*DescribeHybridCloudClusterRulesResponseBodyData
	SetRequestId(v string) *DescribeHybridCloudClusterRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeHybridCloudClusterRulesResponseBody
	GetTotalCount() *int64
}

type DescribeHybridCloudClusterRulesResponseBody struct {
	Data []*DescribeHybridCloudClusterRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHybridCloudClusterRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudClusterRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudClusterRulesResponseBody) GetData() []*DescribeHybridCloudClusterRulesResponseBodyData {
	return s.Data
}

func (s *DescribeHybridCloudClusterRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridCloudClusterRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeHybridCloudClusterRulesResponseBody) SetData(v []*DescribeHybridCloudClusterRulesResponseBodyData) *DescribeHybridCloudClusterRulesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeHybridCloudClusterRulesResponseBody) SetRequestId(v string) *DescribeHybridCloudClusterRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudClusterRulesResponseBody) SetTotalCount(v int64) *DescribeHybridCloudClusterRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHybridCloudClusterRulesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHybridCloudClusterRulesResponseBodyData struct {
	// example:
	//
	// 1099
	ClusterId *int64 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// hdbc-clusterrule-*****khzre0ym0w
	ClusterRuleResourceId *string `json:"ClusterRuleResourceId,omitempty" xml:"ClusterRuleResourceId,omitempty"`
	// example:
	//
	// {\\"check_mode\\": \\"all\\", \\"type\\": \\"exact\\", \\"substance\\": \\"122\\"}
	RuleConfig *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// example:
	//
	// pullin
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeHybridCloudClusterRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudClusterRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudClusterRulesResponseBodyData) GetClusterId() *int64 {
	return s.ClusterId
}

func (s *DescribeHybridCloudClusterRulesResponseBodyData) GetClusterRuleResourceId() *string {
	return s.ClusterRuleResourceId
}

func (s *DescribeHybridCloudClusterRulesResponseBodyData) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *DescribeHybridCloudClusterRulesResponseBodyData) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeHybridCloudClusterRulesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeHybridCloudClusterRulesResponseBodyData) GetVersion() *int64 {
	return s.Version
}

func (s *DescribeHybridCloudClusterRulesResponseBodyData) SetClusterId(v int64) *DescribeHybridCloudClusterRulesResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *DescribeHybridCloudClusterRulesResponseBodyData) SetClusterRuleResourceId(v string) *DescribeHybridCloudClusterRulesResponseBodyData {
	s.ClusterRuleResourceId = &v
	return s
}

func (s *DescribeHybridCloudClusterRulesResponseBodyData) SetRuleConfig(v string) *DescribeHybridCloudClusterRulesResponseBodyData {
	s.RuleConfig = &v
	return s
}

func (s *DescribeHybridCloudClusterRulesResponseBodyData) SetRuleType(v string) *DescribeHybridCloudClusterRulesResponseBodyData {
	s.RuleType = &v
	return s
}

func (s *DescribeHybridCloudClusterRulesResponseBodyData) SetStatus(v string) *DescribeHybridCloudClusterRulesResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeHybridCloudClusterRulesResponseBodyData) SetVersion(v int64) *DescribeHybridCloudClusterRulesResponseBodyData {
	s.Version = &v
	return s
}

func (s *DescribeHybridCloudClusterRulesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
