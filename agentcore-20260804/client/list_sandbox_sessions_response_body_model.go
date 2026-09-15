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
	Code           *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	HttpStatusCode *int32                                  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Items          []*ListSandboxSessionsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	MaxResults     *int32                                  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	Message        *string                                 `json:"message,omitempty" xml:"message,omitempty"`
	NextToken      *string                                 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId      *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount     *int64                                  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	SessionId   *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	SourceType  *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
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
