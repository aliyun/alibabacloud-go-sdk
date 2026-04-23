// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryClientTreeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ClientTreeDTO) *ModelRouterQueryClientTreeResponseBody
	GetData() []*ClientTreeDTO
	SetErrCode(v string) *ModelRouterQueryClientTreeResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryClientTreeResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryClientTreeResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryClientTreeResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryClientTreeResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ModelRouterQueryClientTreeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryClientTreeResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryClientTreeResponseBody struct {
	Data []*ClientTreeDTO `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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
	// maxResults
	//
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

func (s ModelRouterQueryClientTreeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryClientTreeResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryClientTreeResponseBody) GetData() []*ClientTreeDTO {
	return s.Data
}

func (s *ModelRouterQueryClientTreeResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryClientTreeResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryClientTreeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryClientTreeResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryClientTreeResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryClientTreeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryClientTreeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryClientTreeResponseBody) SetData(v []*ClientTreeDTO) *ModelRouterQueryClientTreeResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryClientTreeResponseBody) SetErrCode(v string) *ModelRouterQueryClientTreeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryClientTreeResponseBody) SetErrMessage(v string) *ModelRouterQueryClientTreeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryClientTreeResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryClientTreeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryClientTreeResponseBody) SetMaxResults(v int32) *ModelRouterQueryClientTreeResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryClientTreeResponseBody) SetNextToken(v string) *ModelRouterQueryClientTreeResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryClientTreeResponseBody) SetRequestId(v string) *ModelRouterQueryClientTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryClientTreeResponseBody) SetSuccess(v bool) *ModelRouterQueryClientTreeResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryClientTreeResponseBody) Validate() error {
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
