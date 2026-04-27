// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryBillingCostBreakdownResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *BillingCostBreakdownRespDTO) *ModelRouterQueryBillingCostBreakdownResponseBody
	GetData() *BillingCostBreakdownRespDTO
	SetErrCode(v string) *ModelRouterQueryBillingCostBreakdownResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryBillingCostBreakdownResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryBillingCostBreakdownResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryBillingCostBreakdownResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryBillingCostBreakdownResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ModelRouterQueryBillingCostBreakdownResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryBillingCostBreakdownResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryBillingCostBreakdownResponseBody struct {
	// example:
	//
	// {}
	Data *BillingCostBreakdownRespDTO `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xxxx-xxx-xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterQueryBillingCostBreakdownResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryBillingCostBreakdownResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryBillingCostBreakdownResponseBody) GetData() *BillingCostBreakdownRespDTO {
	return s.Data
}

func (s *ModelRouterQueryBillingCostBreakdownResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryBillingCostBreakdownResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryBillingCostBreakdownResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryBillingCostBreakdownResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryBillingCostBreakdownResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryBillingCostBreakdownResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryBillingCostBreakdownResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryBillingCostBreakdownResponseBody) SetData(v *BillingCostBreakdownRespDTO) *ModelRouterQueryBillingCostBreakdownResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownResponseBody) SetErrCode(v string) *ModelRouterQueryBillingCostBreakdownResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownResponseBody) SetErrMessage(v string) *ModelRouterQueryBillingCostBreakdownResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryBillingCostBreakdownResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownResponseBody) SetMaxResults(v int32) *ModelRouterQueryBillingCostBreakdownResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownResponseBody) SetNextToken(v string) *ModelRouterQueryBillingCostBreakdownResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownResponseBody) SetRequestId(v string) *ModelRouterQueryBillingCostBreakdownResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownResponseBody) SetSuccess(v bool) *ModelRouterQueryBillingCostBreakdownResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
