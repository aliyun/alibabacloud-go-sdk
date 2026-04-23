// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBillingCostTabsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*CostTabDTO) *ModelRouterBillingCostTabsResponseBody
	GetData() []*CostTabDTO
	SetErrCode(v string) *ModelRouterBillingCostTabsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterBillingCostTabsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterBillingCostTabsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterBillingCostTabsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterBillingCostTabsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ModelRouterBillingCostTabsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterBillingCostTabsResponseBody
	GetSuccess() *bool
}

type ModelRouterBillingCostTabsResponseBody struct {
	// example:
	//
	// []
	Data []*CostTabDTO `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// nextToken
	//
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

func (s ModelRouterBillingCostTabsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBillingCostTabsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterBillingCostTabsResponseBody) GetData() []*CostTabDTO {
	return s.Data
}

func (s *ModelRouterBillingCostTabsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterBillingCostTabsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterBillingCostTabsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterBillingCostTabsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterBillingCostTabsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterBillingCostTabsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterBillingCostTabsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterBillingCostTabsResponseBody) SetData(v []*CostTabDTO) *ModelRouterBillingCostTabsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterBillingCostTabsResponseBody) SetErrCode(v string) *ModelRouterBillingCostTabsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterBillingCostTabsResponseBody) SetErrMessage(v string) *ModelRouterBillingCostTabsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterBillingCostTabsResponseBody) SetHttpStatusCode(v int32) *ModelRouterBillingCostTabsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterBillingCostTabsResponseBody) SetMaxResults(v int32) *ModelRouterBillingCostTabsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterBillingCostTabsResponseBody) SetNextToken(v string) *ModelRouterBillingCostTabsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterBillingCostTabsResponseBody) SetRequestId(v string) *ModelRouterBillingCostTabsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterBillingCostTabsResponseBody) SetSuccess(v bool) *ModelRouterBillingCostTabsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterBillingCostTabsResponseBody) Validate() error {
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
