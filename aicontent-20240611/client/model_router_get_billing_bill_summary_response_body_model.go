// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetBillingBillSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *BillingBillSummaryRespDTO) *ModelRouterGetBillingBillSummaryResponseBody
	GetData() *BillingBillSummaryRespDTO
	SetErrCode(v string) *ModelRouterGetBillingBillSummaryResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterGetBillingBillSummaryResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterGetBillingBillSummaryResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterGetBillingBillSummaryResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterGetBillingBillSummaryResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ModelRouterGetBillingBillSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterGetBillingBillSummaryResponseBody
	GetSuccess() *bool
}

type ModelRouterGetBillingBillSummaryResponseBody struct {
	// The data object.
	Data *BillingBillSummaryRespDTO `json:"data,omitempty" xml:"data,omitempty"`
	// The fault code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	MaxResults     *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken      *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterGetBillingBillSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetBillingBillSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterGetBillingBillSummaryResponseBody) GetData() *BillingBillSummaryRespDTO {
	return s.Data
}

func (s *ModelRouterGetBillingBillSummaryResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterGetBillingBillSummaryResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterGetBillingBillSummaryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterGetBillingBillSummaryResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterGetBillingBillSummaryResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterGetBillingBillSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterGetBillingBillSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterGetBillingBillSummaryResponseBody) SetData(v *BillingBillSummaryRespDTO) *ModelRouterGetBillingBillSummaryResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterGetBillingBillSummaryResponseBody) SetErrCode(v string) *ModelRouterGetBillingBillSummaryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterGetBillingBillSummaryResponseBody) SetErrMessage(v string) *ModelRouterGetBillingBillSummaryResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterGetBillingBillSummaryResponseBody) SetHttpStatusCode(v int32) *ModelRouterGetBillingBillSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterGetBillingBillSummaryResponseBody) SetMaxResults(v int32) *ModelRouterGetBillingBillSummaryResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterGetBillingBillSummaryResponseBody) SetNextToken(v string) *ModelRouterGetBillingBillSummaryResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterGetBillingBillSummaryResponseBody) SetRequestId(v string) *ModelRouterGetBillingBillSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterGetBillingBillSummaryResponseBody) SetSuccess(v bool) *ModelRouterGetBillingBillSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterGetBillingBillSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
