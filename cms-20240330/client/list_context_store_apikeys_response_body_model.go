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
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xCs4wJD41qEejNkappMSJ1OL2Ky2GeKLqmBLJrC61WrgUOj9F-31jHbo5Kgqzifv
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Results   []*ListContextStoreAPIKeysResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	// example:
	//
	// 454
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
	// example:
	//
	// sk-3ac8d45d741e4f31b81aa6ee984ce9fd
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// example:
	//
	// test-context-Store
	ContextStoreName *string `json:"contextStoreName,omitempty" xml:"contextStoreName,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 1778205145
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// Production Service Key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// test-workspace
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListContextStoreAPIKeysResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s ListContextStoreAPIKeysResponseBodyResults) GoString() string {
	return s.String()
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

func (s *ListContextStoreAPIKeysResponseBodyResults) GetWorkspace() *string {
	return s.Workspace
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

func (s *ListContextStoreAPIKeysResponseBodyResults) SetWorkspace(v string) *ListContextStoreAPIKeysResponseBodyResults {
	s.Workspace = &v
	return s
}

func (s *ListContextStoreAPIKeysResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
