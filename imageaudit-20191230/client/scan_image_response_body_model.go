// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ScanImageResponseBodyData) *ScanImageResponseBody
	GetData() *ScanImageResponseBodyData
	SetRequestId(v string) *ScanImageResponseBody
	GetRequestId() *string
}

type ScanImageResponseBody struct {
	Data *ScanImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 69B41AE8-1234-1234-1234-12D395695D2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScanImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScanImageResponseBody) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBody) GetData() *ScanImageResponseBodyData {
	return s.Data
}

func (s *ScanImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScanImageResponseBody) SetData(v *ScanImageResponseBodyData) *ScanImageResponseBody {
	s.Data = v
	return s
}

func (s *ScanImageResponseBody) SetRequestId(v string) *ScanImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScanImageResponseBody) Validate() error {
	return dara.Validate(s)
}

type ScanImageResponseBodyData struct {
	Results []*ScanImageResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s ScanImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ScanImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyData) GetResults() []*ScanImageResponseBodyDataResults {
	return s.Results
}

func (s *ScanImageResponseBodyData) SetResults(v []*ScanImageResponseBodyDataResults) *ScanImageResponseBodyData {
	s.Results = v
	return s
}

func (s *ScanImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ScanImageResponseBodyDataResults struct {
	// example:
	//
	// uuid-xxxx-xxx-1234
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// example:
	//
	// http://xxx.xxx.xxx/xxx.jpg
	ImageURL   *string                                       `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	SubResults []*ScanImageResponseBodyDataResultsSubResults `json:"SubResults,omitempty" xml:"SubResults,omitempty" type:"Repeated"`
	// example:
	//
	// img4wlJcb7p4wH4lAP3111111-123456
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ScanImageResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s ScanImageResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResults) GetDataId() *string {
	return s.DataId
}

func (s *ScanImageResponseBodyDataResults) GetImageURL() *string {
	return s.ImageURL
}

func (s *ScanImageResponseBodyDataResults) GetSubResults() []*ScanImageResponseBodyDataResultsSubResults {
	return s.SubResults
}

func (s *ScanImageResponseBodyDataResults) GetTaskId() *string {
	return s.TaskId
}

func (s *ScanImageResponseBodyDataResults) SetDataId(v string) *ScanImageResponseBodyDataResults {
	s.DataId = &v
	return s
}

func (s *ScanImageResponseBodyDataResults) SetImageURL(v string) *ScanImageResponseBodyDataResults {
	s.ImageURL = &v
	return s
}

func (s *ScanImageResponseBodyDataResults) SetSubResults(v []*ScanImageResponseBodyDataResultsSubResults) *ScanImageResponseBodyDataResults {
	s.SubResults = v
	return s
}

func (s *ScanImageResponseBodyDataResults) SetTaskId(v string) *ScanImageResponseBodyDataResults {
	s.TaskId = &v
	return s
}

func (s *ScanImageResponseBodyDataResults) Validate() error {
	return dara.Validate(s)
}

type ScanImageResponseBodyDataResultsSubResults struct {
	Frames            []*ScanImageResponseBodyDataResultsSubResultsFrames            `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Repeated"`
	HintWordsInfoList []*ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList `json:"HintWordsInfoList,omitempty" xml:"HintWordsInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// normal
	Label        *string                                                   `json:"Label,omitempty" xml:"Label,omitempty"`
	LogoDataList []*ScanImageResponseBodyDataResultsSubResultsLogoDataList `json:"LogoDataList,omitempty" xml:"LogoDataList,omitempty" type:"Repeated"`
	// 1
	OCRDataList         []*string                                                        `json:"OCRDataList,omitempty" xml:"OCRDataList,omitempty" type:"Repeated"`
	ProgramCodeDataList []*ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList `json:"ProgramCodeDataList,omitempty" xml:"ProgramCodeDataList,omitempty" type:"Repeated"`
	// example:
	//
	// 99.91
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// example:
	//
	// ad
	Scene         *string                                                    `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SfaceDataList []*ScanImageResponseBodyDataResultsSubResultsSfaceDataList `json:"SfaceDataList,omitempty" xml:"SfaceDataList,omitempty" type:"Repeated"`
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResults) String() string {
	return dara.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResults) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResults) GetFrames() []*ScanImageResponseBodyDataResultsSubResultsFrames {
	return s.Frames
}

func (s *ScanImageResponseBodyDataResultsSubResults) GetHintWordsInfoList() []*ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList {
	return s.HintWordsInfoList
}

func (s *ScanImageResponseBodyDataResultsSubResults) GetLabel() *string {
	return s.Label
}

func (s *ScanImageResponseBodyDataResultsSubResults) GetLogoDataList() []*ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	return s.LogoDataList
}

func (s *ScanImageResponseBodyDataResultsSubResults) GetOCRDataList() []*string {
	return s.OCRDataList
}

func (s *ScanImageResponseBodyDataResultsSubResults) GetProgramCodeDataList() []*ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList {
	return s.ProgramCodeDataList
}

func (s *ScanImageResponseBodyDataResultsSubResults) GetRate() *float32 {
	return s.Rate
}

func (s *ScanImageResponseBodyDataResultsSubResults) GetScene() *string {
	return s.Scene
}

func (s *ScanImageResponseBodyDataResultsSubResults) GetSfaceDataList() []*ScanImageResponseBodyDataResultsSubResultsSfaceDataList {
	return s.SfaceDataList
}

func (s *ScanImageResponseBodyDataResultsSubResults) GetSuggestion() *string {
	return s.Suggestion
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetFrames(v []*ScanImageResponseBodyDataResultsSubResultsFrames) *ScanImageResponseBodyDataResultsSubResults {
	s.Frames = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetHintWordsInfoList(v []*ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList) *ScanImageResponseBodyDataResultsSubResults {
	s.HintWordsInfoList = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetLabel(v string) *ScanImageResponseBodyDataResultsSubResults {
	s.Label = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetLogoDataList(v []*ScanImageResponseBodyDataResultsSubResultsLogoDataList) *ScanImageResponseBodyDataResultsSubResults {
	s.LogoDataList = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetOCRDataList(v []*string) *ScanImageResponseBodyDataResultsSubResults {
	s.OCRDataList = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetProgramCodeDataList(v []*ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) *ScanImageResponseBodyDataResultsSubResults {
	s.ProgramCodeDataList = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetRate(v float32) *ScanImageResponseBodyDataResultsSubResults {
	s.Rate = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetScene(v string) *ScanImageResponseBodyDataResultsSubResults {
	s.Scene = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetSfaceDataList(v []*ScanImageResponseBodyDataResultsSubResultsSfaceDataList) *ScanImageResponseBodyDataResultsSubResults {
	s.SfaceDataList = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetSuggestion(v string) *ScanImageResponseBodyDataResultsSubResults {
	s.Suggestion = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) Validate() error {
	return dara.Validate(s)
}

type ScanImageResponseBodyDataResultsSubResultsFrames struct {
	// example:
	//
	// 89.85
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// example:
	//
	// http://xxx.xxx.com/xxx-0.jpg
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResultsFrames) String() string {
	return dara.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResultsFrames) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResultsFrames) GetRate() *float32 {
	return s.Rate
}

func (s *ScanImageResponseBodyDataResultsSubResultsFrames) GetURL() *string {
	return s.URL
}

func (s *ScanImageResponseBodyDataResultsSubResultsFrames) SetRate(v float32) *ScanImageResponseBodyDataResultsSubResultsFrames {
	s.Rate = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsFrames) SetURL(v string) *ScanImageResponseBodyDataResultsSubResultsFrames {
	s.URL = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsFrames) Validate() error {
	return dara.Validate(s)
}

type ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList struct {
	// example:
	//
	// abc
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList) String() string {
	return dara.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList) GetContext() *string {
	return s.Context
}

func (s *ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList) SetContext(v string) *ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList {
	s.Context = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList) Validate() error {
	return dara.Validate(s)
}

type ScanImageResponseBodyDataResultsSubResultsLogoDataList struct {
	// example:
	//
	// 106
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// TV
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 106
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 140
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 68
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResultsLogoDataList) String() string {
	return dara.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResultsLogoDataList) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) GetHeight() *float32 {
	return s.Height
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) GetName() *string {
	return s.Name
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) GetType() *string {
	return s.Type
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) GetWidth() *float32 {
	return s.Width
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) GetX() *float32 {
	return s.X
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) GetY() *float32 {
	return s.Y
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) SetHeight(v float32) *ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	s.Height = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) SetName(v string) *ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	s.Name = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) SetType(v string) *ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	s.Type = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) SetWidth(v float32) *ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	s.Width = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) SetX(v float32) *ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	s.X = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) SetY(v float32) *ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	s.Y = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) Validate() error {
	return dara.Validate(s)
}

type ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList struct {
	// example:
	//
	// 413.0
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 402.0
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 11.0
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.0
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) String() string {
	return dara.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) GetHeight() *float32 {
	return s.Height
}

func (s *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) GetWidth() *float32 {
	return s.Width
}

func (s *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) GetX() *float32 {
	return s.X
}

func (s *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) GetY() *float32 {
	return s.Y
}

func (s *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) SetHeight(v float32) *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList {
	s.Height = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) SetWidth(v float32) *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList {
	s.Width = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) SetX(v float32) *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList {
	s.X = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) SetY(v float32) *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList {
	s.Y = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) Validate() error {
	return dara.Validate(s)
}

type ScanImageResponseBodyDataResultsSubResultsSfaceDataList struct {
	Faces []*ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	// example:
	//
	// 131
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 97
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 49
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 39
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResultsSfaceDataList) String() string {
	return dara.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResultsSfaceDataList) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) GetFaces() []*ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces {
	return s.Faces
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) GetHeight() *float32 {
	return s.Height
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) GetWidth() *float32 {
	return s.Width
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) GetX() *float32 {
	return s.X
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) GetY() *float32 {
	return s.Y
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) SetFaces(v []*ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) *ScanImageResponseBodyDataResultsSubResultsSfaceDataList {
	s.Faces = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) SetHeight(v float32) *ScanImageResponseBodyDataResultsSubResultsSfaceDataList {
	s.Height = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) SetWidth(v float32) *ScanImageResponseBodyDataResultsSubResultsSfaceDataList {
	s.Width = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) SetX(v float32) *ScanImageResponseBodyDataResultsSubResultsSfaceDataList {
	s.X = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) SetY(v float32) *ScanImageResponseBodyDataResultsSubResultsSfaceDataList {
	s.Y = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) Validate() error {
	return dara.Validate(s)
}

type ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces struct {
	// example:
	//
	// AliFace_0001234
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 91.54
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) String() string {
	return dara.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) GetId() *string {
	return s.Id
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) GetName() *string {
	return s.Name
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) GetRate() *float32 {
	return s.Rate
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) SetId(v string) *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces {
	s.Id = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) SetName(v string) *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces {
	s.Name = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) SetRate(v float32) *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces {
	s.Rate = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) Validate() error {
	return dara.Validate(s)
}
