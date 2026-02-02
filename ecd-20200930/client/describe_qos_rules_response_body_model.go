// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQosRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQosRules(v []*DescribeQosRulesResponseBodyQosRules) *DescribeQosRulesResponseBody
	GetQosRules() []*DescribeQosRulesResponseBodyQosRules
	SetRequestId(v string) *DescribeQosRulesResponseBody
	GetRequestId() *string
}

type DescribeQosRulesResponseBody struct {
	QosRules []*DescribeQosRulesResponseBodyQosRules `json:"QosRules,omitempty" xml:"QosRules,omitempty" type:"Repeated"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeQosRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQosRulesResponseBody) GetQosRules() []*DescribeQosRulesResponseBodyQosRules {
	return s.QosRules
}

func (s *DescribeQosRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeQosRulesResponseBody) SetQosRules(v []*DescribeQosRulesResponseBodyQosRules) *DescribeQosRulesResponseBody {
	s.QosRules = v
	return s
}

func (s *DescribeQosRulesResponseBody) SetRequestId(v string) *DescribeQosRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQosRulesResponseBody) Validate() error {
	if s.QosRules != nil {
		for _, item := range s.QosRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeQosRulesResponseBodyQosRules struct {
	// example:
	//
	// 0
	DesktopCount *string `json:"DesktopCount,omitempty" xml:"DesktopCount,omitempty"`
	// example:
	//
	// 10
	Download *string `json:"Download,omitempty" xml:"Download,omitempty"`
	// example:
	//
	// np-5cjh3sqs1ty3s02xq
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	// example:
	//
	// qos-chvkz5ekzgcb6bo0f
	QosRuleId *string `json:"QosRuleId,omitempty" xml:"QosRuleId,omitempty"`
	// example:
	//
	// test
	QosRuleName *string `json:"QosRuleName,omitempty" xml:"QosRuleName,omitempty"`
	// example:
	//
	// 10
	Upload *string `json:"Upload,omitempty" xml:"Upload,omitempty"`
}

func (s DescribeQosRulesResponseBodyQosRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosRulesResponseBodyQosRules) GoString() string {
	return s.String()
}

func (s *DescribeQosRulesResponseBodyQosRules) GetDesktopCount() *string {
	return s.DesktopCount
}

func (s *DescribeQosRulesResponseBodyQosRules) GetDownload() *string {
	return s.Download
}

func (s *DescribeQosRulesResponseBodyQosRules) GetNetworkPackageId() *string {
	return s.NetworkPackageId
}

func (s *DescribeQosRulesResponseBodyQosRules) GetQosRuleId() *string {
	return s.QosRuleId
}

func (s *DescribeQosRulesResponseBodyQosRules) GetQosRuleName() *string {
	return s.QosRuleName
}

func (s *DescribeQosRulesResponseBodyQosRules) GetUpload() *string {
	return s.Upload
}

func (s *DescribeQosRulesResponseBodyQosRules) SetDesktopCount(v string) *DescribeQosRulesResponseBodyQosRules {
	s.DesktopCount = &v
	return s
}

func (s *DescribeQosRulesResponseBodyQosRules) SetDownload(v string) *DescribeQosRulesResponseBodyQosRules {
	s.Download = &v
	return s
}

func (s *DescribeQosRulesResponseBodyQosRules) SetNetworkPackageId(v string) *DescribeQosRulesResponseBodyQosRules {
	s.NetworkPackageId = &v
	return s
}

func (s *DescribeQosRulesResponseBodyQosRules) SetQosRuleId(v string) *DescribeQosRulesResponseBodyQosRules {
	s.QosRuleId = &v
	return s
}

func (s *DescribeQosRulesResponseBodyQosRules) SetQosRuleName(v string) *DescribeQosRulesResponseBodyQosRules {
	s.QosRuleName = &v
	return s
}

func (s *DescribeQosRulesResponseBodyQosRules) SetUpload(v string) *DescribeQosRulesResponseBodyQosRules {
	s.Upload = &v
	return s
}

func (s *DescribeQosRulesResponseBodyQosRules) Validate() error {
	return dara.Validate(s)
}
