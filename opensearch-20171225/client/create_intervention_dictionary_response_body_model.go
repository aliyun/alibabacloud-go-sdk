// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInterventionDictionaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateInterventionDictionaryResponseBody
	GetRequestId() *string
	SetResult(v *CreateInterventionDictionaryResponseBodyResult) *CreateInterventionDictionaryResponseBody
	GetResult() *CreateInterventionDictionaryResponseBodyResult
}

type CreateInterventionDictionaryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 80326EFE-485F-46E7-B291-5A1DD08D2198
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	Result *CreateInterventionDictionaryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateInterventionDictionaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInterventionDictionaryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInterventionDictionaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInterventionDictionaryResponseBody) GetResult() *CreateInterventionDictionaryResponseBodyResult {
	return s.Result
}

func (s *CreateInterventionDictionaryResponseBody) SetRequestId(v string) *CreateInterventionDictionaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInterventionDictionaryResponseBody) SetResult(v *CreateInterventionDictionaryResponseBodyResult) *CreateInterventionDictionaryResponseBody {
	s.Result = v
	return s
}

func (s *CreateInterventionDictionaryResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateInterventionDictionaryResponseBodyResult struct {
	// The custom analyzer.
	//
	// example:
	//
	// dianshang
	Analyzer *string `json:"analyzer,omitempty" xml:"analyzer,omitempty"`
	// The time when the test scenario was created.
	//
	// example:
	//
	// 1591086323
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The name of the test group.
	//
	// example:
	//
	// testb
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the intervention dictionary. Valid values:
	//
	// 	- stopword: an intervention dictionary for stop word filtering
	//
	// 	- synonym: an intervention dictionary for synonym configuration
	//
	// 	- correction: an intervention dictionary for spelling correction
	//
	// 	- category_prediction: an intervention dictionary for category prediction
	//
	// 	- ner: an intervention dictionary for named entity recognition (NER)
	//
	// 	- term_weighting: an intervention dictionary for term weight analysis
	//
	// example:
	//
	// ner
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the intervention dictionary was last updated.
	//
	// example:
	//
	// 1591086323
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateInterventionDictionaryResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateInterventionDictionaryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateInterventionDictionaryResponseBodyResult) GetAnalyzer() *string {
	return s.Analyzer
}

func (s *CreateInterventionDictionaryResponseBodyResult) GetCreated() *string {
	return s.Created
}

func (s *CreateInterventionDictionaryResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateInterventionDictionaryResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *CreateInterventionDictionaryResponseBodyResult) GetUpdated() *string {
	return s.Updated
}

func (s *CreateInterventionDictionaryResponseBodyResult) SetAnalyzer(v string) *CreateInterventionDictionaryResponseBodyResult {
	s.Analyzer = &v
	return s
}

func (s *CreateInterventionDictionaryResponseBodyResult) SetCreated(v string) *CreateInterventionDictionaryResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateInterventionDictionaryResponseBodyResult) SetName(v string) *CreateInterventionDictionaryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateInterventionDictionaryResponseBodyResult) SetType(v string) *CreateInterventionDictionaryResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateInterventionDictionaryResponseBodyResult) SetUpdated(v string) *CreateInterventionDictionaryResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *CreateInterventionDictionaryResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
