// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInsightTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInsightTypes(v map[string]interface{}) *GetInsightTypesResponseBody
	GetInsightTypes() map[string]interface{}
	SetRequestId(v string) *GetInsightTypesResponseBody
	GetRequestId() *string
}

type GetInsightTypesResponseBody struct {
	// The types of Insights events.
	//
	// example:
	//
	// {\\"ApiCallRateInsight\\": \\"Enable\\", \\"ApiErrorRateInsight\\": \\"Enable\\", \\"IpInsight\\": \\"Enable\\", \\"AkInsight\\": \\"Enable\\"}
	InsightTypes map[string]interface{} `json:"InsightTypes,omitempty" xml:"InsightTypes,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EC4A1F64-4927-5714-B205-5A0B16A2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInsightTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInsightTypesResponseBody) GoString() string {
	return s.String()
}

func (s *GetInsightTypesResponseBody) GetInsightTypes() map[string]interface{} {
	return s.InsightTypes
}

func (s *GetInsightTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInsightTypesResponseBody) SetInsightTypes(v map[string]interface{}) *GetInsightTypesResponseBody {
	s.InsightTypes = v
	return s
}

func (s *GetInsightTypesResponseBody) SetRequestId(v string) *GetInsightTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInsightTypesResponseBody) Validate() error {
	return dara.Validate(s)
}
