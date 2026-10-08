// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggregations(v []*SimpleQueryResponseBodyAggregations) *SimpleQueryResponseBody
	GetAggregations() []*SimpleQueryResponseBodyAggregations
	SetFiles(v []*File) *SimpleQueryResponseBody
	GetFiles() []*File
	SetNextToken(v string) *SimpleQueryResponseBody
	GetNextToken() *string
	SetRequestId(v string) *SimpleQueryResponseBody
	GetRequestId() *string
	SetTotalHits(v int64) *SimpleQueryResponseBody
	GetTotalHits() *int64
}

type SimpleQueryResponseBody struct {
	// The list of aggregation field information. This parameter is returned only when Aggregations in the request is not empty.
	Aggregations []*SimpleQueryResponseBodyAggregations `json:"Aggregations,omitempty" xml:"Aggregations,omitempty" type:"Repeated"`
	// The list of file information. This parameter is returned only when Aggregations in the request is empty.
	Files []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The token used for pagination when the total number of files exceeds the value of MaxResults.
	//
	// When you list file information next time, set NextToken to this value to return the remaining results.
	//
	// This parameter has a value only when not all files are returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpwZw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2C5C1E0F-D8B8-4DA0-8127-EC32C771****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of matched records.
	//
	// example:
	//
	// 10
	TotalHits *int64 `json:"TotalHits,omitempty" xml:"TotalHits,omitempty"`
}

func (s SimpleQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SimpleQueryResponseBody) GoString() string {
	return s.String()
}

func (s *SimpleQueryResponseBody) GetAggregations() []*SimpleQueryResponseBodyAggregations {
	return s.Aggregations
}

func (s *SimpleQueryResponseBody) GetFiles() []*File {
	return s.Files
}

func (s *SimpleQueryResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *SimpleQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SimpleQueryResponseBody) GetTotalHits() *int64 {
	return s.TotalHits
}

func (s *SimpleQueryResponseBody) SetAggregations(v []*SimpleQueryResponseBodyAggregations) *SimpleQueryResponseBody {
	s.Aggregations = v
	return s
}

func (s *SimpleQueryResponseBody) SetFiles(v []*File) *SimpleQueryResponseBody {
	s.Files = v
	return s
}

func (s *SimpleQueryResponseBody) SetNextToken(v string) *SimpleQueryResponseBody {
	s.NextToken = &v
	return s
}

func (s *SimpleQueryResponseBody) SetRequestId(v string) *SimpleQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *SimpleQueryResponseBody) SetTotalHits(v int64) *SimpleQueryResponseBody {
	s.TotalHits = &v
	return s
}

func (s *SimpleQueryResponseBody) Validate() error {
	if s.Aggregations != nil {
		for _, item := range s.Aggregations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SimpleQueryResponseBodyAggregations struct {
	// The name of the aggregation field.
	//
	// example:
	//
	// Size
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The list of grouping and aggregation results. This parameter is returned only when an Operation of the group type exists in Aggregations of the request.
	Groups []*SimpleQueryResponseBodyAggregationsGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The aggregation operation for the aggregation field.
	//
	// example:
	//
	// sum
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The statistical result of the aggregation.
	//
	// example:
	//
	// 200
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SimpleQueryResponseBodyAggregations) String() string {
	return dara.Prettify(s)
}

func (s SimpleQueryResponseBodyAggregations) GoString() string {
	return s.String()
}

func (s *SimpleQueryResponseBodyAggregations) GetField() *string {
	return s.Field
}

func (s *SimpleQueryResponseBodyAggregations) GetGroups() []*SimpleQueryResponseBodyAggregationsGroups {
	return s.Groups
}

func (s *SimpleQueryResponseBodyAggregations) GetOperation() *string {
	return s.Operation
}

func (s *SimpleQueryResponseBodyAggregations) GetValue() *float64 {
	return s.Value
}

func (s *SimpleQueryResponseBodyAggregations) SetField(v string) *SimpleQueryResponseBodyAggregations {
	s.Field = &v
	return s
}

func (s *SimpleQueryResponseBodyAggregations) SetGroups(v []*SimpleQueryResponseBodyAggregationsGroups) *SimpleQueryResponseBodyAggregations {
	s.Groups = v
	return s
}

func (s *SimpleQueryResponseBodyAggregations) SetOperation(v string) *SimpleQueryResponseBodyAggregations {
	s.Operation = &v
	return s
}

func (s *SimpleQueryResponseBodyAggregations) SetValue(v float64) *SimpleQueryResponseBodyAggregations {
	s.Value = &v
	return s
}

func (s *SimpleQueryResponseBodyAggregations) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SimpleQueryResponseBodyAggregationsGroups struct {
	// The total count of the grouping and aggregation.
	//
	// example:
	//
	// 5
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The value of the grouping and aggregation.
	//
	// example:
	//
	// 100
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SimpleQueryResponseBodyAggregationsGroups) String() string {
	return dara.Prettify(s)
}

func (s SimpleQueryResponseBodyAggregationsGroups) GoString() string {
	return s.String()
}

func (s *SimpleQueryResponseBodyAggregationsGroups) GetCount() *int64 {
	return s.Count
}

func (s *SimpleQueryResponseBodyAggregationsGroups) GetValue() *string {
	return s.Value
}

func (s *SimpleQueryResponseBodyAggregationsGroups) SetCount(v int64) *SimpleQueryResponseBodyAggregationsGroups {
	s.Count = &v
	return s
}

func (s *SimpleQueryResponseBodyAggregationsGroups) SetValue(v string) *SimpleQueryResponseBodyAggregationsGroups {
	s.Value = &v
	return s
}

func (s *SimpleQueryResponseBodyAggregationsGroups) Validate() error {
	return dara.Validate(s)
}
