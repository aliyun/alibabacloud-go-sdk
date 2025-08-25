// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ScanTextResponseBodyData) *ScanTextResponseBody
	GetData() *ScanTextResponseBodyData
	SetRequestId(v string) *ScanTextResponseBody
	GetRequestId() *string
}

type ScanTextResponseBody struct {
	Data *ScanTextResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C7CD87E3-57A5-4E2F-8A44-809F3554692C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScanTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScanTextResponseBody) GoString() string {
	return s.String()
}

func (s *ScanTextResponseBody) GetData() *ScanTextResponseBodyData {
	return s.Data
}

func (s *ScanTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScanTextResponseBody) SetData(v *ScanTextResponseBodyData) *ScanTextResponseBody {
	s.Data = v
	return s
}

func (s *ScanTextResponseBody) SetRequestId(v string) *ScanTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScanTextResponseBody) Validate() error {
	return dara.Validate(s)
}

type ScanTextResponseBodyData struct {
	Elements []*ScanTextResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s ScanTextResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ScanTextResponseBodyData) GoString() string {
	return s.String()
}

func (s *ScanTextResponseBodyData) GetElements() []*ScanTextResponseBodyDataElements {
	return s.Elements
}

func (s *ScanTextResponseBodyData) SetElements(v []*ScanTextResponseBodyDataElements) *ScanTextResponseBodyData {
	s.Elements = v
	return s
}

func (s *ScanTextResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ScanTextResponseBodyDataElements struct {
	Results []*ScanTextResponseBodyDataElementsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// example:
	//
	// txt6Vh5Fv0DAFy5hgdVRt3pqf-1s82jj
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ScanTextResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s ScanTextResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *ScanTextResponseBodyDataElements) GetResults() []*ScanTextResponseBodyDataElementsResults {
	return s.Results
}

func (s *ScanTextResponseBodyDataElements) GetTaskId() *string {
	return s.TaskId
}

func (s *ScanTextResponseBodyDataElements) SetResults(v []*ScanTextResponseBodyDataElementsResults) *ScanTextResponseBodyDataElements {
	s.Results = v
	return s
}

func (s *ScanTextResponseBodyDataElements) SetTaskId(v string) *ScanTextResponseBodyDataElements {
	s.TaskId = &v
	return s
}

func (s *ScanTextResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}

type ScanTextResponseBodyDataElementsResults struct {
	Details []*ScanTextResponseBodyDataElementsResultsDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// 99.91
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s ScanTextResponseBodyDataElementsResults) String() string {
	return dara.Prettify(s)
}

func (s ScanTextResponseBodyDataElementsResults) GoString() string {
	return s.String()
}

func (s *ScanTextResponseBodyDataElementsResults) GetDetails() []*ScanTextResponseBodyDataElementsResultsDetails {
	return s.Details
}

func (s *ScanTextResponseBodyDataElementsResults) GetLabel() *string {
	return s.Label
}

func (s *ScanTextResponseBodyDataElementsResults) GetRate() *float32 {
	return s.Rate
}

func (s *ScanTextResponseBodyDataElementsResults) GetSuggestion() *string {
	return s.Suggestion
}

func (s *ScanTextResponseBodyDataElementsResults) SetDetails(v []*ScanTextResponseBodyDataElementsResultsDetails) *ScanTextResponseBodyDataElementsResults {
	s.Details = v
	return s
}

func (s *ScanTextResponseBodyDataElementsResults) SetLabel(v string) *ScanTextResponseBodyDataElementsResults {
	s.Label = &v
	return s
}

func (s *ScanTextResponseBodyDataElementsResults) SetRate(v float32) *ScanTextResponseBodyDataElementsResults {
	s.Rate = &v
	return s
}

func (s *ScanTextResponseBodyDataElementsResults) SetSuggestion(v string) *ScanTextResponseBodyDataElementsResults {
	s.Suggestion = &v
	return s
}

func (s *ScanTextResponseBodyDataElementsResults) Validate() error {
	return dara.Validate(s)
}

type ScanTextResponseBodyDataElementsResultsDetails struct {
	Contexts []*ScanTextResponseBodyDataElementsResultsDetailsContexts `json:"Contexts,omitempty" xml:"Contexts,omitempty" type:"Repeated"`
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s ScanTextResponseBodyDataElementsResultsDetails) String() string {
	return dara.Prettify(s)
}

func (s ScanTextResponseBodyDataElementsResultsDetails) GoString() string {
	return s.String()
}

func (s *ScanTextResponseBodyDataElementsResultsDetails) GetContexts() []*ScanTextResponseBodyDataElementsResultsDetailsContexts {
	return s.Contexts
}

func (s *ScanTextResponseBodyDataElementsResultsDetails) GetLabel() *string {
	return s.Label
}

func (s *ScanTextResponseBodyDataElementsResultsDetails) SetContexts(v []*ScanTextResponseBodyDataElementsResultsDetailsContexts) *ScanTextResponseBodyDataElementsResultsDetails {
	s.Contexts = v
	return s
}

func (s *ScanTextResponseBodyDataElementsResultsDetails) SetLabel(v string) *ScanTextResponseBodyDataElementsResultsDetails {
	s.Label = &v
	return s
}

func (s *ScanTextResponseBodyDataElementsResultsDetails) Validate() error {
	return dara.Validate(s)
}

type ScanTextResponseBodyDataElementsResultsDetailsContexts struct {
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
}

func (s ScanTextResponseBodyDataElementsResultsDetailsContexts) String() string {
	return dara.Prettify(s)
}

func (s ScanTextResponseBodyDataElementsResultsDetailsContexts) GoString() string {
	return s.String()
}

func (s *ScanTextResponseBodyDataElementsResultsDetailsContexts) GetContext() *string {
	return s.Context
}

func (s *ScanTextResponseBodyDataElementsResultsDetailsContexts) SetContext(v string) *ScanTextResponseBodyDataElementsResultsDetailsContexts {
	s.Context = &v
	return s
}

func (s *ScanTextResponseBodyDataElementsResultsDetailsContexts) Validate() error {
	return dara.Validate(s)
}
