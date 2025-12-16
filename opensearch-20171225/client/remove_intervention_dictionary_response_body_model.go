// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInterventionDictionaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveInterventionDictionaryResponseBody
	GetRequestId() *string
	SetResult(v *RemoveInterventionDictionaryResponseBodyResult) *RemoveInterventionDictionaryResponseBody
	GetResult() *RemoveInterventionDictionaryResponseBodyResult
}

type RemoveInterventionDictionaryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 06BBD41A-5F72-34E4-2DAF-E43B0526051D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the intervention dictionary.
	Result *RemoveInterventionDictionaryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s RemoveInterventionDictionaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveInterventionDictionaryResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveInterventionDictionaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveInterventionDictionaryResponseBody) GetResult() *RemoveInterventionDictionaryResponseBodyResult {
	return s.Result
}

func (s *RemoveInterventionDictionaryResponseBody) SetRequestId(v string) *RemoveInterventionDictionaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveInterventionDictionaryResponseBody) SetResult(v *RemoveInterventionDictionaryResponseBodyResult) *RemoveInterventionDictionaryResponseBody {
	s.Result = v
	return s
}

func (s *RemoveInterventionDictionaryResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveInterventionDictionaryResponseBodyResult struct {
	// The custom analyzer.
	//
	// example:
	//
	// ""
	Analyzer *string `json:"analyzer,omitempty" xml:"analyzer,omitempty"`
	// The time when the intervention dictionary was created.
	//
	// example:
	//
	// 1539158313
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// Parameter
	//
	// example:
	//
	// tongyici
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
	// synonym
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the the intervention dictionary was modified.
	//
	// example:
	//
	// 1539158313
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s RemoveInterventionDictionaryResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s RemoveInterventionDictionaryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RemoveInterventionDictionaryResponseBodyResult) GetAnalyzer() *string {
	return s.Analyzer
}

func (s *RemoveInterventionDictionaryResponseBodyResult) GetCreated() *string {
	return s.Created
}

func (s *RemoveInterventionDictionaryResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *RemoveInterventionDictionaryResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *RemoveInterventionDictionaryResponseBodyResult) GetUpdated() *string {
	return s.Updated
}

func (s *RemoveInterventionDictionaryResponseBodyResult) SetAnalyzer(v string) *RemoveInterventionDictionaryResponseBodyResult {
	s.Analyzer = &v
	return s
}

func (s *RemoveInterventionDictionaryResponseBodyResult) SetCreated(v string) *RemoveInterventionDictionaryResponseBodyResult {
	s.Created = &v
	return s
}

func (s *RemoveInterventionDictionaryResponseBodyResult) SetName(v string) *RemoveInterventionDictionaryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *RemoveInterventionDictionaryResponseBodyResult) SetType(v string) *RemoveInterventionDictionaryResponseBodyResult {
	s.Type = &v
	return s
}

func (s *RemoveInterventionDictionaryResponseBodyResult) SetUpdated(v string) *RemoveInterventionDictionaryResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *RemoveInterventionDictionaryResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
