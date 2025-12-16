// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueryProcessorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListQueryProcessorsResponseBody
	GetRequestId() *string
	SetResult(v []*ListQueryProcessorsResponseBodyResult) *ListQueryProcessorsResponseBody
	GetResult() []*ListQueryProcessorsResponseBodyResult
}

type ListQueryProcessorsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the query analysis rule.
	//
	// For more information, see [QueryProcessor](https://help.aliyun.com/document_detail/170014.html).
	Result []*ListQueryProcessorsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListQueryProcessorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQueryProcessorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQueryProcessorsResponseBody) GetResult() []*ListQueryProcessorsResponseBodyResult {
	return s.Result
}

func (s *ListQueryProcessorsResponseBody) SetRequestId(v string) *ListQueryProcessorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueryProcessorsResponseBody) SetResult(v []*ListQueryProcessorsResponseBodyResult) *ListQueryProcessorsResponseBody {
	s.Result = v
	return s
}

func (s *ListQueryProcessorsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQueryProcessorsResponseBodyResult struct {
	// Indicates whether the query analysis rule is a default rule.
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the query analysis rule was created.
	//
	// example:
	//
	// 1587398402
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The type of the industry to which the query analysis rule is applied. Valid values:
	//
	// 	- GENERAL
	//
	// 	- ECOMMERCE
	//
	// 	- IT_CONTENT
	//
	// example:
	//
	// GENERAL
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The indexes to which the query analysis rule is applied.
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The name of the query analysis rule.
	//
	// example:
	//
	// ner
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The features that are used in the query analysis rule.
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The time when the query analysis rule was last modified.
	//
	// example:
	//
	// 1587398402
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListQueryProcessorsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListQueryProcessorsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorsResponseBodyResult) GetActive() *bool {
	return s.Active
}

func (s *ListQueryProcessorsResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *ListQueryProcessorsResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *ListQueryProcessorsResponseBodyResult) GetIndexes() []*string {
	return s.Indexes
}

func (s *ListQueryProcessorsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListQueryProcessorsResponseBodyResult) GetProcessors() []map[string]interface{} {
	return s.Processors
}

func (s *ListQueryProcessorsResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *ListQueryProcessorsResponseBodyResult) SetActive(v bool) *ListQueryProcessorsResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) SetCreated(v int32) *ListQueryProcessorsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) SetDomain(v string) *ListQueryProcessorsResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) SetIndexes(v []*string) *ListQueryProcessorsResponseBodyResult {
	s.Indexes = v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) SetName(v string) *ListQueryProcessorsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) SetProcessors(v []map[string]interface{}) *ListQueryProcessorsResponseBodyResult {
	s.Processors = v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) SetUpdated(v int32) *ListQueryProcessorsResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
