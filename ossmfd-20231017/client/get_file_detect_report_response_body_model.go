// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileDetectReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetFileDetectReportResponseBody
	GetCode() *string
	SetData(v *GetFileDetectReportResponseBodyData) *GetFileDetectReportResponseBody
	GetData() *GetFileDetectReportResponseBodyData
	SetMessage(v string) *GetFileDetectReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFileDetectReportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFileDetectReportResponseBody
	GetSuccess() *bool
}

type GetFileDetectReportResponseBody struct {
	// example:
	//
	// 200
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetFileDetectReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAC546A5-5EDC-5939-86A3-56DFAF******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFileDetectReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileDetectReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileDetectReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFileDetectReportResponseBody) GetData() *GetFileDetectReportResponseBodyData {
	return s.Data
}

func (s *GetFileDetectReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFileDetectReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileDetectReportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFileDetectReportResponseBody) SetCode(v string) *GetFileDetectReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetFileDetectReportResponseBody) SetData(v *GetFileDetectReportResponseBodyData) *GetFileDetectReportResponseBody {
	s.Data = v
	return s
}

func (s *GetFileDetectReportResponseBody) SetMessage(v string) *GetFileDetectReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetFileDetectReportResponseBody) SetRequestId(v string) *GetFileDetectReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileDetectReportResponseBody) SetSuccess(v bool) *GetFileDetectReportResponseBody {
	s.Success = &v
	return s
}

func (s *GetFileDetectReportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFileDetectReportResponseBodyData struct {
	// example:
	//
	// {
	//
	//         "sha256": "",
	//
	//         "sha512": "",
	//
	//         "source": "aegis",
	//
	//         "gmt_first_submit": "",
	//
	//         "sha1": "",
	//
	//         "virus_result": "",
	//
	//         "webshell_result": "",
	//
	//         "gmt_update": "",
	//
	//         "sandbox_result": "2",
	//
	//         "fileSize": "363752",
	//
	//         "virus_name": "",
	//
	//     }
	Basic *string `json:"Basic,omitempty" xml:"Basic,omitempty"`
	// example:
	//
	// 4206dbcf1c2bc80ea95ad64043******
	FileHash *string `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	// example:
	//
	// testFile******
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// example:
	//
	// true
	HasData       *bool   `json:"HasData,omitempty" xml:"HasData,omitempty"`
	Intelligences *string `json:"Intelligences,omitempty" xml:"Intelligences,omitempty"`
	// example:
	//
	// {\\"BehaviorData\\": {}, \\"ProcessData\\": {}, \\"SandboxData\\": {}, \\"AttackData\\": [], \\"NetworkData\\": {}, \\"SolutionData\\": {}, \\"FileData\\": {}}
	Sandbox *string `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	// example:
	//
	// true
	ShowTab *bool `json:"ShowTab,omitempty" xml:"ShowTab,omitempty"`
	// example:
	//
	// 2
	ThreatLevel *int64  `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	ThreatTypes *string `json:"ThreatTypes,omitempty" xml:"ThreatTypes,omitempty"`
}

func (s GetFileDetectReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFileDetectReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFileDetectReportResponseBodyData) GetBasic() *string {
	return s.Basic
}

func (s *GetFileDetectReportResponseBodyData) GetFileHash() *string {
	return s.FileHash
}

func (s *GetFileDetectReportResponseBodyData) GetFilename() *string {
	return s.Filename
}

func (s *GetFileDetectReportResponseBodyData) GetHasData() *bool {
	return s.HasData
}

func (s *GetFileDetectReportResponseBodyData) GetIntelligences() *string {
	return s.Intelligences
}

func (s *GetFileDetectReportResponseBodyData) GetSandbox() *string {
	return s.Sandbox
}

func (s *GetFileDetectReportResponseBodyData) GetShowTab() *bool {
	return s.ShowTab
}

func (s *GetFileDetectReportResponseBodyData) GetThreatLevel() *int64 {
	return s.ThreatLevel
}

func (s *GetFileDetectReportResponseBodyData) GetThreatTypes() *string {
	return s.ThreatTypes
}

func (s *GetFileDetectReportResponseBodyData) SetBasic(v string) *GetFileDetectReportResponseBodyData {
	s.Basic = &v
	return s
}

func (s *GetFileDetectReportResponseBodyData) SetFileHash(v string) *GetFileDetectReportResponseBodyData {
	s.FileHash = &v
	return s
}

func (s *GetFileDetectReportResponseBodyData) SetFilename(v string) *GetFileDetectReportResponseBodyData {
	s.Filename = &v
	return s
}

func (s *GetFileDetectReportResponseBodyData) SetHasData(v bool) *GetFileDetectReportResponseBodyData {
	s.HasData = &v
	return s
}

func (s *GetFileDetectReportResponseBodyData) SetIntelligences(v string) *GetFileDetectReportResponseBodyData {
	s.Intelligences = &v
	return s
}

func (s *GetFileDetectReportResponseBodyData) SetSandbox(v string) *GetFileDetectReportResponseBodyData {
	s.Sandbox = &v
	return s
}

func (s *GetFileDetectReportResponseBodyData) SetShowTab(v bool) *GetFileDetectReportResponseBodyData {
	s.ShowTab = &v
	return s
}

func (s *GetFileDetectReportResponseBodyData) SetThreatLevel(v int64) *GetFileDetectReportResponseBodyData {
	s.ThreatLevel = &v
	return s
}

func (s *GetFileDetectReportResponseBodyData) SetThreatTypes(v string) *GetFileDetectReportResponseBodyData {
	s.ThreatTypes = &v
	return s
}

func (s *GetFileDetectReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}
