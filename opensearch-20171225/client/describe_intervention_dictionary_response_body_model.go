// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInterventionDictionaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeInterventionDictionaryResponseBody
	GetRequestId() *string
	SetResult(v *DescribeInterventionDictionaryResponseBodyResult) *DescribeInterventionDictionaryResponseBody
	GetResult() *DescribeInterventionDictionaryResponseBodyResult
}

type DescribeInterventionDictionaryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D7CCF454-472A-030E-F254-604520B028AA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details about the intervention dictionary.
	Result *DescribeInterventionDictionaryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeInterventionDictionaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterventionDictionaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInterventionDictionaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInterventionDictionaryResponseBody) GetResult() *DescribeInterventionDictionaryResponseBodyResult {
	return s.Result
}

func (s *DescribeInterventionDictionaryResponseBody) SetRequestId(v string) *DescribeInterventionDictionaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInterventionDictionaryResponseBody) SetResult(v *DescribeInterventionDictionaryResponseBodyResult) *DescribeInterventionDictionaryResponseBody {
	s.Result = v
	return s
}

func (s *DescribeInterventionDictionaryResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInterventionDictionaryResponseBodyResult struct {
	// The custom analyzer.
	//
	// example:
	//
	// -
	Analyzer *string `json:"analyzer,omitempty" xml:"analyzer,omitempty"`
	// The time when the the intervention dictionary was created.
	//
	// example:
	//
	// 1536233287
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The name of the the intervention dictionary.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Type
	//
	// 	- stopword: an intervention dictionary for stop word filtering.
	//
	// 	- synonym: an intervention dictionary for synonym configuration.
	//
	// 	- correction: an intervention dictionary for spelling correction.
	//
	// 	- category_prediction: an intervention dictionary for category prediction.
	//
	// 	- ner: an intervention dictionary for named entity recognition.
	//
	// 	- term_weighting: an intervention dictionary for term weight analysis.
	//
	// example:
	//
	// category_prediction
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the the intervention dictionary was modified.
	//
	// example:
	//
	// 1536233287
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeInterventionDictionaryResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterventionDictionaryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeInterventionDictionaryResponseBodyResult) GetAnalyzer() *string {
	return s.Analyzer
}

func (s *DescribeInterventionDictionaryResponseBodyResult) GetCreated() *string {
	return s.Created
}

func (s *DescribeInterventionDictionaryResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *DescribeInterventionDictionaryResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *DescribeInterventionDictionaryResponseBodyResult) GetUpdated() *string {
	return s.Updated
}

func (s *DescribeInterventionDictionaryResponseBodyResult) SetAnalyzer(v string) *DescribeInterventionDictionaryResponseBodyResult {
	s.Analyzer = &v
	return s
}

func (s *DescribeInterventionDictionaryResponseBodyResult) SetCreated(v string) *DescribeInterventionDictionaryResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeInterventionDictionaryResponseBodyResult) SetName(v string) *DescribeInterventionDictionaryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeInterventionDictionaryResponseBodyResult) SetType(v string) *DescribeInterventionDictionaryResponseBodyResult {
	s.Type = &v
	return s
}

func (s *DescribeInterventionDictionaryResponseBodyResult) SetUpdated(v string) *DescribeInterventionDictionaryResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *DescribeInterventionDictionaryResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
