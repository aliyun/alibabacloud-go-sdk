// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterventionDictionariesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListInterventionDictionariesResponseBody
	GetRequestId() *string
	SetResult(v []*ListInterventionDictionariesResponseBodyResult) *ListInterventionDictionariesResponseBody
	GetResult() []*ListInterventionDictionariesResponseBodyResult
	SetTotalCount(v int32) *ListInterventionDictionariesResponseBody
	GetTotalCount() *int32
}

type ListInterventionDictionariesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about each intervention dictionary.
	//
	// For more information, see [InterventionDictionary](https://help.aliyun.com/document_detail/173608.html).
	Result []*ListInterventionDictionariesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInterventionDictionariesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInterventionDictionariesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionariesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInterventionDictionariesResponseBody) GetResult() []*ListInterventionDictionariesResponseBodyResult {
	return s.Result
}

func (s *ListInterventionDictionariesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInterventionDictionariesResponseBody) SetRequestId(v string) *ListInterventionDictionariesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterventionDictionariesResponseBody) SetResult(v []*ListInterventionDictionariesResponseBodyResult) *ListInterventionDictionariesResponseBody {
	s.Result = v
	return s
}

func (s *ListInterventionDictionariesResponseBody) SetTotalCount(v int32) *ListInterventionDictionariesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInterventionDictionariesResponseBody) Validate() error {
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

type ListInterventionDictionariesResponseBodyResult struct {
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
	// 1539158325
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the intervention dictionary.
	//
	// example:
	//
	// 1
	Id *int32 `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the intervention dictionary.
	//
	// example:
	//
	// tongyici
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
	// synonym
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the intervention dictionary was last updated.
	//
	// example:
	//
	// 1539158313
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListInterventionDictionariesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListInterventionDictionariesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionariesResponseBodyResult) GetAnalyzer() *string {
	return s.Analyzer
}

func (s *ListInterventionDictionariesResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *ListInterventionDictionariesResponseBodyResult) GetId() *int32 {
	return s.Id
}

func (s *ListInterventionDictionariesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListInterventionDictionariesResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListInterventionDictionariesResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *ListInterventionDictionariesResponseBodyResult) SetAnalyzer(v string) *ListInterventionDictionariesResponseBodyResult {
	s.Analyzer = &v
	return s
}

func (s *ListInterventionDictionariesResponseBodyResult) SetCreated(v int32) *ListInterventionDictionariesResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListInterventionDictionariesResponseBodyResult) SetId(v int32) *ListInterventionDictionariesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListInterventionDictionariesResponseBodyResult) SetName(v string) *ListInterventionDictionariesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListInterventionDictionariesResponseBodyResult) SetType(v string) *ListInterventionDictionariesResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListInterventionDictionariesResponseBodyResult) SetUpdated(v int32) *ListInterventionDictionariesResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListInterventionDictionariesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
