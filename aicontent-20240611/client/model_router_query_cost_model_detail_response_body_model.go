// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryCostModelDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CostModelDetailRespDTO) *ModelRouterQueryCostModelDetailResponseBody
	GetData() *CostModelDetailRespDTO
	SetErrCode(v string) *ModelRouterQueryCostModelDetailResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryCostModelDetailResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryCostModelDetailResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryCostModelDetailResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryCostModelDetailResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ModelRouterQueryCostModelDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryCostModelDetailResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryCostModelDetailResponseBody struct {
	// example:
	//
	// []
	Data *CostModelDetailRespDTO `json:"data,omitempty" xml:"data,omitempty"`
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

func (s ModelRouterQueryCostModelDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryCostModelDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryCostModelDetailResponseBody) GetData() *CostModelDetailRespDTO {
	return s.Data
}

func (s *ModelRouterQueryCostModelDetailResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryCostModelDetailResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryCostModelDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryCostModelDetailResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryCostModelDetailResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryCostModelDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryCostModelDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryCostModelDetailResponseBody) SetData(v *CostModelDetailRespDTO) *ModelRouterQueryCostModelDetailResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryCostModelDetailResponseBody) SetErrCode(v string) *ModelRouterQueryCostModelDetailResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailResponseBody) SetErrMessage(v string) *ModelRouterQueryCostModelDetailResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryCostModelDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailResponseBody) SetMaxResults(v int32) *ModelRouterQueryCostModelDetailResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailResponseBody) SetNextToken(v string) *ModelRouterQueryCostModelDetailResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailResponseBody) SetRequestId(v string) *ModelRouterQueryCostModelDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailResponseBody) SetSuccess(v bool) *ModelRouterQueryCostModelDetailResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
