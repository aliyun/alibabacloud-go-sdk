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
	// The subscription information.
	Data *ModelRouterListSubscriptionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The fault error message encoding.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error.
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
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
	// The list of subscription information.
	List []*SubscriptionDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The maximum number of results per page. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Do not specify this parameter for the first query. For subsequent queries, specify the value returned from the previous query. Set to "" when no more data is available. Set to "5" when there is a next page.
	//
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
