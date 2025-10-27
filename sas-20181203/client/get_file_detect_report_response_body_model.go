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
	// The status code that is returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *GetFileDetectReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB393XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
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
	// The basic information about the detected file.
	//
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
	// The hash value of the file.
	//
	// example:
	//
	// c42b5f6bde0b730ece2923266333****
	FileHash *string `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// test.zip
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// Indicates whether the file data exists in the cloud sandbox. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HasData *bool `json:"HasData,omitempty" xml:"HasData,omitempty"`
	// The threat intelligence event, which is a JSON array.
	//
	// Valid values:
	//
	// 	- The threat type. The value is an array. The elements in the array can be DDoS trojans, mining programs, network layer intrusions, network service scans, network sharing and discovery, mining pools, exploits, dark webs, malicious logons, malicious download sources, C\\&C servers, webshells, and web attacks.
	//
	// example:
	//
	// ["The threat type"]
	Intelligences *string `json:"Intelligences,omitempty" xml:"Intelligences,omitempty"`
	// The details of the cloud sandbox check results.
	//
	// example:
	//
	// {\\"BehaviorData\\": {}, \\"ProcessData\\": {}, \\"SandboxData\\": {}, \\"AttackData\\": [], \\"NetworkData\\": {}, \\"SolutionData\\": {}, \\"FileData\\": {}}
	Sandbox *string `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	// Indicates whether the check report is displayed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ShowTab *bool `json:"ShowTab,omitempty" xml:"ShowTab,omitempty"`
	// The threat level. Valid values:
	//
	// 	- **0**: normal
	//
	// 	- **1**: suspicious
	//
	// 	- **2**: high
	//
	// example:
	//
	// 2
	ThreatLevel *int64 `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	// The risk tags and server tags that are generated by analyzing threat intelligence and security events. The value is a string array. The array includes the following elements:
	//
	// 	- **threat_type_desc**: the threat type.
	//
	// 	- **last_find_time**: the last time the threat was detected.
	//
	// 	- **risk_type**: indicates whether the tag is malicious. The value 0 indicates that the tag is not malicious. The value 1 indicates that the tag is malicious. The value -1 indicates that whether the tag type is malicious is unknown.
	//
	// 	- **threat_type**: the threat type. The value is an array. The elements in the array can be network layer intrusion, network service scanning, network sharing and discovery, mining pool, exploits, darknet, malicious logon, malicious download source, central control, web shell, and web attack.
	//
	// example:
	//
	// [{"threat_type_desc": "test","risk_type": 1,"threat_type": ""}]
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
