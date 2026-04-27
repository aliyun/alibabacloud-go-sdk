// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryUsageBreakdownResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UsageBreakdownRespDTO) *ModelRouterQueryUsageBreakdownResponseBody
	GetData() *UsageBreakdownRespDTO
	SetErrCode(v string) *ModelRouterQueryUsageBreakdownResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryUsageBreakdownResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryUsageBreakdownResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryUsageBreakdownResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryUsageBreakdownResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ModelRouterQueryUsageBreakdownResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryUsageBreakdownResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryUsageBreakdownResponseBody struct {
	// example:
	//
	// {}
	Data *UsageBreakdownRespDTO `json:"data,omitempty" xml:"data,omitempty"`
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

func (s ModelRouterQueryUsageBreakdownResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryUsageBreakdownResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryUsageBreakdownResponseBody) GetData() *UsageBreakdownRespDTO {
	return s.Data
}

func (s *ModelRouterQueryUsageBreakdownResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryUsageBreakdownResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryUsageBreakdownResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryUsageBreakdownResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryUsageBreakdownResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryUsageBreakdownResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryUsageBreakdownResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryUsageBreakdownResponseBody) SetData(v *UsageBreakdownRespDTO) *ModelRouterQueryUsageBreakdownResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryUsageBreakdownResponseBody) SetErrCode(v string) *ModelRouterQueryUsageBreakdownResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryUsageBreakdownResponseBody) SetErrMessage(v string) *ModelRouterQueryUsageBreakdownResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryUsageBreakdownResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryUsageBreakdownResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryUsageBreakdownResponseBody) SetMaxResults(v int32) *ModelRouterQueryUsageBreakdownResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryUsageBreakdownResponseBody) SetNextToken(v string) *ModelRouterQueryUsageBreakdownResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryUsageBreakdownResponseBody) SetRequestId(v string) *ModelRouterQueryUsageBreakdownResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryUsageBreakdownResponseBody) SetSuccess(v bool) *ModelRouterQueryUsageBreakdownResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryUsageBreakdownResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
