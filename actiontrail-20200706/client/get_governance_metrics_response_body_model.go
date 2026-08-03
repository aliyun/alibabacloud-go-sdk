// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGovernanceMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetGovernanceMetricsResponseBodyData) *GetGovernanceMetricsResponseBody
	GetData() *GetGovernanceMetricsResponseBodyData
	SetRequestId(v string) *GetGovernanceMetricsResponseBody
	GetRequestId() *string
}

type GetGovernanceMetricsResponseBody struct {
	// The response parameters.
	Data *GetGovernanceMetricsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 145318BE-DEE1-4C57-AA7C-5BE7D34A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGovernanceMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetGovernanceMetricsResponseBody) GetData() *GetGovernanceMetricsResponseBodyData {
	return s.Data
}

func (s *GetGovernanceMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGovernanceMetricsResponseBody) SetData(v *GetGovernanceMetricsResponseBodyData) *GetGovernanceMetricsResponseBody {
	s.Data = v
	return s
}

func (s *GetGovernanceMetricsResponseBody) SetRequestId(v string) *GetGovernanceMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGovernanceMetricsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGovernanceMetricsResponseBodyData struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 195622768501****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// A collection of governance items that contain multiple compliance assessment dimensions.
	GovernanceMetrics []*GetGovernanceMetricsResponseBodyDataGovernanceMetrics `json:"GovernanceMetrics,omitempty" xml:"GovernanceMetrics,omitempty" type:"Repeated"`
}

func (s GetGovernanceMetricsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceMetricsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGovernanceMetricsResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *GetGovernanceMetricsResponseBodyData) GetGovernanceMetrics() []*GetGovernanceMetricsResponseBodyDataGovernanceMetrics {
	return s.GovernanceMetrics
}

func (s *GetGovernanceMetricsResponseBodyData) SetAccountId(v string) *GetGovernanceMetricsResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *GetGovernanceMetricsResponseBodyData) SetGovernanceMetrics(v []*GetGovernanceMetricsResponseBodyDataGovernanceMetrics) *GetGovernanceMetricsResponseBodyData {
	s.GovernanceMetrics = v
	return s
}

func (s *GetGovernanceMetricsResponseBodyData) Validate() error {
	if s.GovernanceMetrics != nil {
		for _, item := range s.GovernanceMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGovernanceMetricsResponseBodyDataGovernanceMetrics struct {
	// The details of the resource.
	//
	// This parameter contains the detailed configurations of all compliant resources for the governance item. This parameter is returned only if a resource instance exists.
	//
	// example:
	//
	// {
	//
	//     "trailName": "trail-test",
	//
	//     "homeRegion": "cn-hangzhou",
	//
	//     "trailRegion": "All",
	//
	//     "trailStatus": "Enable",
	//
	//     "eventRW": "All",
	//
	//     "isOrganizationTrail": false,
	//
	//     "ossDeliveryStatus": "normal",
	//
	//     "deliveryObjectLifeCycle": "999",
	//
	//     "ossBucketLifeCycle": "999",
	//
	//     "trailTotal": 100
	//
	// }
	ColumnsSchema *string `json:"ColumnsSchema,omitempty" xml:"ColumnsSchema,omitempty"`
	// The governance item. This indicates a specific category of compliance check.
	//
	// example:
	//
	// actiontrail_storage_audit_log
	GovernanceItem *string `json:"GovernanceItem,omitempty" xml:"GovernanceItem,omitempty"`
	// The compliance score for the governance item.
	//
	// Valid values: 0 to 100.
	//
	// example:
	//
	// 100
	GovernanceScore *string `json:"GovernanceScore,omitempty" xml:"GovernanceScore,omitempty"`
}

func (s GetGovernanceMetricsResponseBodyDataGovernanceMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceMetricsResponseBodyDataGovernanceMetrics) GoString() string {
	return s.String()
}

func (s *GetGovernanceMetricsResponseBodyDataGovernanceMetrics) GetColumnsSchema() *string {
	return s.ColumnsSchema
}

func (s *GetGovernanceMetricsResponseBodyDataGovernanceMetrics) GetGovernanceItem() *string {
	return s.GovernanceItem
}

func (s *GetGovernanceMetricsResponseBodyDataGovernanceMetrics) GetGovernanceScore() *string {
	return s.GovernanceScore
}

func (s *GetGovernanceMetricsResponseBodyDataGovernanceMetrics) SetColumnsSchema(v string) *GetGovernanceMetricsResponseBodyDataGovernanceMetrics {
	s.ColumnsSchema = &v
	return s
}

func (s *GetGovernanceMetricsResponseBodyDataGovernanceMetrics) SetGovernanceItem(v string) *GetGovernanceMetricsResponseBodyDataGovernanceMetrics {
	s.GovernanceItem = &v
	return s
}

func (s *GetGovernanceMetricsResponseBodyDataGovernanceMetrics) SetGovernanceScore(v string) *GetGovernanceMetricsResponseBodyDataGovernanceMetrics {
	s.GovernanceScore = &v
	return s
}

func (s *GetGovernanceMetricsResponseBodyDataGovernanceMetrics) Validate() error {
	return dara.Validate(s)
}
