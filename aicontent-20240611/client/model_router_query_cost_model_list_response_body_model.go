// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryCostModelListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CostQueryModelsDTO) *ModelRouterQueryCostModelListResponseBody
	GetData() *CostQueryModelsDTO
	SetErrCode(v string) *ModelRouterQueryCostModelListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryCostModelListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryCostModelListResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryCostModelListResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryCostModelListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ModelRouterQueryCostModelListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryCostModelListResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryCostModelListResponseBody struct {
	// The data object.
	//
	// example:
	//
	// []
	Data *CostQueryModelsDTO `json:"data,omitempty" xml:"data,omitempty"`
	// The error code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The number of results returned on the current page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for retrieving the next page of results. If this parameter is not returned, all results have been returned.
	//
	// example:
	//
	// xxxx-xxx-xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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

func (s ModelRouterQueryCostModelListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryCostModelListResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryCostModelListResponseBody) GetData() *CostQueryModelsDTO {
	return s.Data
}

func (s *ModelRouterQueryCostModelListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryCostModelListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryCostModelListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryCostModelListResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryCostModelListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryCostModelListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryCostModelListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryCostModelListResponseBody) SetData(v *CostQueryModelsDTO) *ModelRouterQueryCostModelListResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryCostModelListResponseBody) SetErrCode(v string) *ModelRouterQueryCostModelListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryCostModelListResponseBody) SetErrMessage(v string) *ModelRouterQueryCostModelListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryCostModelListResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryCostModelListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryCostModelListResponseBody) SetMaxResults(v int32) *ModelRouterQueryCostModelListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryCostModelListResponseBody) SetNextToken(v string) *ModelRouterQueryCostModelListResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryCostModelListResponseBody) SetRequestId(v string) *ModelRouterQueryCostModelListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryCostModelListResponseBody) SetSuccess(v bool) *ModelRouterQueryCostModelListResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryCostModelListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
