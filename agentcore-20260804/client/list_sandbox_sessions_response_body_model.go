// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSandboxSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSandboxSessionsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListSandboxSessionsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListSandboxSessionsResponseBodyItems) *ListSandboxSessionsResponseBody
	GetItems() []*ListSandboxSessionsResponseBodyItems
	SetMaxResults(v int32) *ListSandboxSessionsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListSandboxSessionsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListSandboxSessionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSandboxSessionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSandboxSessionsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListSandboxSessionsResponseBody
	GetTotalCount() *int64
}

type ListSandboxSessionsResponseBody struct {
	// The business status code. The value SUCCESS is returned if the request was successful.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code. The value 200 is returned if the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list of active sessions in the sandbox.
	Items []*ListSandboxSessionsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of records per page for this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message. The value success is returned if the request was successful.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page. An empty value indicates that no more results are available.
	//
	// example:
	//
	// next-token-1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 1a2b3c4d-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records that match the query conditions.
	//
	// example:
	//
	// 27
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListSandboxSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSandboxSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSandboxSessionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSandboxSessionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSandboxSessionsResponseBody) GetItems() []*ListSandboxSessionsResponseBodyItems {
	return s.Items
}

func (s *ListSandboxSessionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSandboxSessionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSandboxSessionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSandboxSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSandboxSessionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSandboxSessionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSandboxSessionsResponseBody) SetCode(v string) *ListSandboxSessionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSandboxSessionsResponseBody) SetHttpStatusCode(v int32) *ListSandboxSessionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSandboxSessionsResponseBody) SetItems(v []*ListSandboxSessionsResponseBodyItems) *ListSandboxSessionsResponseBody {
	s.Items = v
	return s
}

func (s *ListSandboxSessionsResponseBody) SetMaxResults(v int32) *ListSandboxSessionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSandboxSessionsResponseBody) SetMessage(v string) *ListSandboxSessionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSandboxSessionsResponseBody) SetNextToken(v string) *ListSandboxSessionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSandboxSessionsResponseBody) SetRequestId(v string) *ListSandboxSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSandboxSessionsResponseBody) SetSuccess(v bool) *ListSandboxSessionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSandboxSessionsResponseBody) SetTotalCount(v int64) *ListSandboxSessionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSandboxSessionsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSandboxSessionsResponseBodyItems struct {
	// The external channel type, such as DINGTALK, FEISHU, or WECOM. This parameter is empty for non-external channels.
	//
	// example:
	//
	// DINGTALK
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	// The unique identifier of the active session.
	//
	// example:
	//
	// sess-2
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The session source type. Valid values:
	//
	// - API: API call.
	//
	// - CONSOLE_DEBUG: Console debugging.
	//
	// - EXTERNAL_CHANNEL: External channel.
	//
	// - UNKNOWN: Unknown source.
	//
	// example:
	//
	// EXTERNAL_CHANNEL
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListSandboxSessionsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListSandboxSessionsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListSandboxSessionsResponseBodyItems) GetChannelType() *string {
	return s.ChannelType
}

func (s *ListSandboxSessionsResponseBodyItems) GetSessionId() *string {
	return s.SessionId
}

func (s *ListSandboxSessionsResponseBodyItems) GetSourceType() *string {
	return s.SourceType
}

func (s *ListSandboxSessionsResponseBodyItems) SetChannelType(v string) *ListSandboxSessionsResponseBodyItems {
	s.ChannelType = &v
	return s
}

func (s *ListSandboxSessionsResponseBodyItems) SetSessionId(v string) *ListSandboxSessionsResponseBodyItems {
	s.SessionId = &v
	return s
}

func (s *ListSandboxSessionsResponseBodyItems) SetSourceType(v string) *ListSandboxSessionsResponseBodyItems {
	s.SourceType = &v
	return s
}

func (s *ListSandboxSessionsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
