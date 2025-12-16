// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueryProcessorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateQueryProcessorResponseBody
	GetRequestId() *string
	SetResult(v *CreateQueryProcessorResponseBodyResult) *CreateQueryProcessorResponseBody
	GetResult() *CreateQueryProcessorResponseBodyResult
}

type CreateQueryProcessorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the query analysis rule.
	Result *CreateQueryProcessorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateQueryProcessorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQueryProcessorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQueryProcessorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQueryProcessorResponseBody) GetResult() *CreateQueryProcessorResponseBodyResult {
	return s.Result
}

func (s *CreateQueryProcessorResponseBody) SetRequestId(v string) *CreateQueryProcessorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQueryProcessorResponseBody) SetResult(v *CreateQueryProcessorResponseBodyResult) *CreateQueryProcessorResponseBody {
	s.Result = v
	return s
}

func (s *CreateQueryProcessorResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateQueryProcessorResponseBodyResult struct {
	// Indicates whether the query analysis rule is the default one.
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
	// The type of the industry to which the query analysis rule was applied. Valid values:
	//
	// 	- GENERAL: general.
	//
	// 	- ECOMMERCE: e-commerce.
	//
	// 	- IT_CONTENT: IT content.
	//
	// example:
	//
	// GENERAL
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The indexes to which the query analysis rule was applied.
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The name of the query analysis rule.
	//
	// example:
	//
	// query_filter
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The features that are used in the query analysis rule.
	//
	// For more information, see [QueryProcessor](https://help.aliyun.com/document_detail/170014.html).
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The time when the query analysis rule was last modified.
	//
	// example:
	//
	// 1587398402
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateQueryProcessorResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateQueryProcessorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateQueryProcessorResponseBodyResult) GetActive() *bool {
	return s.Active
}

func (s *CreateQueryProcessorResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *CreateQueryProcessorResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *CreateQueryProcessorResponseBodyResult) GetIndexes() []*string {
	return s.Indexes
}

func (s *CreateQueryProcessorResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateQueryProcessorResponseBodyResult) GetProcessors() []map[string]interface{} {
	return s.Processors
}

func (s *CreateQueryProcessorResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *CreateQueryProcessorResponseBodyResult) SetActive(v bool) *CreateQueryProcessorResponseBodyResult {
	s.Active = &v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) SetCreated(v int32) *CreateQueryProcessorResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) SetDomain(v string) *CreateQueryProcessorResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) SetIndexes(v []*string) *CreateQueryProcessorResponseBodyResult {
	s.Indexes = v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) SetName(v string) *CreateQueryProcessorResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) SetProcessors(v []map[string]interface{}) *CreateQueryProcessorResponseBodyResult {
	s.Processors = v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) SetUpdated(v int32) *CreateQueryProcessorResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
