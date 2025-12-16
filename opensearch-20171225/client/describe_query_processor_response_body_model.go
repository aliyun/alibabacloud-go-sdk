// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQueryProcessorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeQueryProcessorResponseBody
	GetRequestId() *string
	SetResult(v *DescribeQueryProcessorResponseBodyResult) *DescribeQueryProcessorResponseBody
	GetResult() *DescribeQueryProcessorResponseBodyResult
}

type DescribeQueryProcessorResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the query analysis rule.
	Result *DescribeQueryProcessorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeQueryProcessorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryProcessorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQueryProcessorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeQueryProcessorResponseBody) GetResult() *DescribeQueryProcessorResponseBodyResult {
	return s.Result
}

func (s *DescribeQueryProcessorResponseBody) SetRequestId(v string) *DescribeQueryProcessorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQueryProcessorResponseBody) SetResult(v *DescribeQueryProcessorResponseBodyResult) *DescribeQueryProcessorResponseBody {
	s.Result = v
	return s
}

func (s *DescribeQueryProcessorResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeQueryProcessorResponseBodyResult struct {
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
	// The type of the industry. Valid values:
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
	// The indexes to which the query analysis rule applies.
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The name of the query analysis rule.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The features that are used in the query analysis rule.
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The time when the query analysis rule was last updated.
	//
	// example:
	//
	// 1587398402
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeQueryProcessorResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryProcessorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeQueryProcessorResponseBodyResult) GetActive() *bool {
	return s.Active
}

func (s *DescribeQueryProcessorResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *DescribeQueryProcessorResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *DescribeQueryProcessorResponseBodyResult) GetIndexes() []*string {
	return s.Indexes
}

func (s *DescribeQueryProcessorResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *DescribeQueryProcessorResponseBodyResult) GetProcessors() []map[string]interface{} {
	return s.Processors
}

func (s *DescribeQueryProcessorResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *DescribeQueryProcessorResponseBodyResult) SetActive(v bool) *DescribeQueryProcessorResponseBodyResult {
	s.Active = &v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) SetCreated(v int32) *DescribeQueryProcessorResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) SetDomain(v string) *DescribeQueryProcessorResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) SetIndexes(v []*string) *DescribeQueryProcessorResponseBodyResult {
	s.Indexes = v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) SetName(v string) *DescribeQueryProcessorResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) SetProcessors(v []map[string]interface{}) *DescribeQueryProcessorResponseBodyResult {
	s.Processors = v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) SetUpdated(v int32) *DescribeQueryProcessorResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
