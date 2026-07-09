// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextStoreAPIKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListContextStoreAPIKeysResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListContextStoreAPIKeysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListContextStoreAPIKeysResponseBody
	GetRequestId() *string
	SetResults(v []*ListContextStoreAPIKeysResponseBodyResults) *ListContextStoreAPIKeysResponseBody
	GetResults() []*ListContextStoreAPIKeysResponseBodyResults
	SetTotal(v int32) *ListContextStoreAPIKeysResponseBody
	GetTotal() *int32
}

type ListContextStoreAPIKeysResponseBody struct {
	// The maximum number of entries per page that was specified in the request. This value is echoed back.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next page. An empty string indicates that the current page is the last page.
	//
	// example:
	//
	// MTIzNDU2Nzg5MA==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of API keys.
	Results []*ListContextStoreAPIKeysResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	// The total number of API keys that match the query conditions.
	//
	// example:
	//
	// 3
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListContextStoreAPIKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContextStoreAPIKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListContextStoreAPIKeysResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListContextStoreAPIKeysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListContextStoreAPIKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContextStoreAPIKeysResponseBody) GetResults() []*ListContextStoreAPIKeysResponseBodyResults {
	return s.Results
}

func (s *ListContextStoreAPIKeysResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListContextStoreAPIKeysResponseBody) SetMaxResults(v int32) *ListContextStoreAPIKeysResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListContextStoreAPIKeysResponseBody) SetNextToken(v string) *ListContextStoreAPIKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListContextStoreAPIKeysResponseBody) SetRequestId(v string) *ListContextStoreAPIKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContextStoreAPIKeysResponseBody) SetResults(v []*ListContextStoreAPIKeysResponseBodyResults) *ListContextStoreAPIKeysResponseBody {
	s.Results = v
	return s
}

func (s *ListContextStoreAPIKeysResponseBody) SetTotal(v int32) *ListContextStoreAPIKeysResponseBody {
	s.Total = &v
	return s
}

func (s *ListContextStoreAPIKeysResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContextStoreAPIKeysResponseBodyResults struct {
	// The name of the AgentSpace to which the API key belongs.
	//
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The full value of the API key. The plaintext value is returned only when the API key is created. In list scenarios, the value is masked based on business rules.
	//
	// example:
	//
	// sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// The name of the context store to which the API key belongs.
	//
	// example:
	//
	// my-context-store
	ContextStoreName *string `json:"contextStoreName,omitempty" xml:"contextStoreName,omitempty"`
	// The time when the API key was created, in ISO 8601 UTC format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The display name of the API key.
	//
	// example:
	//
	// my-api-key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListContextStoreAPIKeysResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s ListContextStoreAPIKeysResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ListContextStoreAPIKeysResponseBodyResults) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *ListContextStoreAPIKeysResponseBodyResults) GetApiKey() *string {
	return s.ApiKey
}

func (s *ListContextStoreAPIKeysResponseBodyResults) GetContextStoreName() *string {
	return s.ContextStoreName
}

func (s *ListContextStoreAPIKeysResponseBodyResults) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListContextStoreAPIKeysResponseBodyResults) GetName() *string {
	return s.Name
}

func (s *ListContextStoreAPIKeysResponseBodyResults) SetAgentSpace(v string) *ListContextStoreAPIKeysResponseBodyResults {
	s.AgentSpace = &v
	return s
}

func (s *ListContextStoreAPIKeysResponseBodyResults) SetApiKey(v string) *ListContextStoreAPIKeysResponseBodyResults {
	s.ApiKey = &v
	return s
}

func (s *ListContextStoreAPIKeysResponseBodyResults) SetContextStoreName(v string) *ListContextStoreAPIKeysResponseBodyResults {
	s.ContextStoreName = &v
	return s
}

func (s *ListContextStoreAPIKeysResponseBodyResults) SetCreateTime(v string) *ListContextStoreAPIKeysResponseBodyResults {
	s.CreateTime = &v
	return s
}

func (s *ListContextStoreAPIKeysResponseBodyResults) SetName(v string) *ListContextStoreAPIKeysResponseBodyResults {
	s.Name = &v
	return s
}

func (s *ListContextStoreAPIKeysResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
