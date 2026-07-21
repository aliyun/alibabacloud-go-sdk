// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterListSubscriptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterListSubscriptionsResponseBodyData) *ModelRouterListSubscriptionsResponseBody
	GetData() *ModelRouterListSubscriptionsResponseBodyData
	SetErrCode(v string) *ModelRouterListSubscriptionsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterListSubscriptionsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterListSubscriptionsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterListSubscriptionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterListSubscriptionsResponseBody
	GetSuccess() *bool
}

type ModelRouterListSubscriptionsResponseBody struct {
	Data *ModelRouterListSubscriptionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterListSubscriptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListSubscriptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterListSubscriptionsResponseBody) GetData() *ModelRouterListSubscriptionsResponseBodyData {
	return s.Data
}

func (s *ModelRouterListSubscriptionsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterListSubscriptionsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterListSubscriptionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterListSubscriptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterListSubscriptionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterListSubscriptionsResponseBody) SetData(v *ModelRouterListSubscriptionsResponseBodyData) *ModelRouterListSubscriptionsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterListSubscriptionsResponseBody) SetErrCode(v string) *ModelRouterListSubscriptionsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterListSubscriptionsResponseBody) SetErrMessage(v string) *ModelRouterListSubscriptionsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterListSubscriptionsResponseBody) SetHttpStatusCode(v int32) *ModelRouterListSubscriptionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterListSubscriptionsResponseBody) SetRequestId(v string) *ModelRouterListSubscriptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterListSubscriptionsResponseBody) SetSuccess(v bool) *ModelRouterListSubscriptionsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterListSubscriptionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterListSubscriptionsResponseBodyData struct {
	List []*SubscriptionDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 5" or ""
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ModelRouterListSubscriptionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListSubscriptionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterListSubscriptionsResponseBodyData) GetList() []*SubscriptionDTO {
	return s.List
}

func (s *ModelRouterListSubscriptionsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterListSubscriptionsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterListSubscriptionsResponseBodyData) SetList(v []*SubscriptionDTO) *ModelRouterListSubscriptionsResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterListSubscriptionsResponseBodyData) SetMaxResults(v int32) *ModelRouterListSubscriptionsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterListSubscriptionsResponseBodyData) SetNextToken(v string) *ModelRouterListSubscriptionsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ModelRouterListSubscriptionsResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
