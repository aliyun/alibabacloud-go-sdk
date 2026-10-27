// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentCodeRuntimesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentCodeRuntimesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListAgentCodeRuntimesResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListAgentCodeRuntimesResponseBodyItems) *ListAgentCodeRuntimesResponseBody
	GetItems() []*ListAgentCodeRuntimesResponseBodyItems
	SetMaxResults(v int32) *ListAgentCodeRuntimesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListAgentCodeRuntimesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListAgentCodeRuntimesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAgentCodeRuntimesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgentCodeRuntimesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListAgentCodeRuntimesResponseBody
	GetTotalCount() *int64
}

type ListAgentCodeRuntimesResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list of supported code runtime environments.
	Items []*ListAgentCodeRuntimesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The actual page size.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The token for the next page.
	//
	// example:
	//
	// d3MtYTFiMmMzZDRlNWY2CmNvZGUtcnVudGltZS1leGFtcGxlCjE
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The number of available runtime environments.
	//
	// example:
	//
	// 3
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAgentCodeRuntimesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentCodeRuntimesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentCodeRuntimesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAgentCodeRuntimesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAgentCodeRuntimesResponseBody) GetItems() []*ListAgentCodeRuntimesResponseBodyItems {
	return s.Items
}

func (s *ListAgentCodeRuntimesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentCodeRuntimesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAgentCodeRuntimesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentCodeRuntimesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentCodeRuntimesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgentCodeRuntimesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAgentCodeRuntimesResponseBody) SetCode(v string) *ListAgentCodeRuntimesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentCodeRuntimesResponseBody) SetHttpStatusCode(v int32) *ListAgentCodeRuntimesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAgentCodeRuntimesResponseBody) SetItems(v []*ListAgentCodeRuntimesResponseBodyItems) *ListAgentCodeRuntimesResponseBody {
	s.Items = v
	return s
}

func (s *ListAgentCodeRuntimesResponseBody) SetMaxResults(v int32) *ListAgentCodeRuntimesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAgentCodeRuntimesResponseBody) SetMessage(v string) *ListAgentCodeRuntimesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentCodeRuntimesResponseBody) SetNextToken(v string) *ListAgentCodeRuntimesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAgentCodeRuntimesResponseBody) SetRequestId(v string) *ListAgentCodeRuntimesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentCodeRuntimesResponseBody) SetSuccess(v bool) *ListAgentCodeRuntimesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgentCodeRuntimesResponseBody) SetTotalCount(v int64) *ListAgentCodeRuntimesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAgentCodeRuntimesResponseBody) Validate() error {
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

type ListAgentCodeRuntimesResponseBodyItems struct {
	// The runtime architecture.
	//
	// example:
	//
	// linux/amd64
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	// The display name of the runtime environment.
	//
	// example:
	//
	// Python 3.12
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The language runtime environment.
	//
	// example:
	//
	// python3.12
	LanguageRuntime *string `json:"languageRuntime,omitempty" xml:"languageRuntime,omitempty"`
}

func (s ListAgentCodeRuntimesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListAgentCodeRuntimesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListAgentCodeRuntimesResponseBodyItems) GetArchitecture() *string {
	return s.Architecture
}

func (s *ListAgentCodeRuntimesResponseBodyItems) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListAgentCodeRuntimesResponseBodyItems) GetLanguageRuntime() *string {
	return s.LanguageRuntime
}

func (s *ListAgentCodeRuntimesResponseBodyItems) SetArchitecture(v string) *ListAgentCodeRuntimesResponseBodyItems {
	s.Architecture = &v
	return s
}

func (s *ListAgentCodeRuntimesResponseBodyItems) SetDisplayName(v string) *ListAgentCodeRuntimesResponseBodyItems {
	s.DisplayName = &v
	return s
}

func (s *ListAgentCodeRuntimesResponseBodyItems) SetLanguageRuntime(v string) *ListAgentCodeRuntimesResponseBodyItems {
	s.LanguageRuntime = &v
	return s
}

func (s *ListAgentCodeRuntimesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
