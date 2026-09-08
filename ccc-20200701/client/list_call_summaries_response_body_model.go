// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallSummariesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCallSummariesResponseBody
	GetCode() *string
	SetData(v []*ListCallSummariesResponseBodyData) *ListCallSummariesResponseBody
	GetData() []*ListCallSummariesResponseBodyData
	SetHttpStatusCode(v int32) *ListCallSummariesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCallSummariesResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListCallSummariesResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListCallSummariesResponseBody
	GetRequestId() *string
}

type ListCallSummariesResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data []*ListCallSummariesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// List of error parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// 932579BC-811A-503D-B322-4C2E57087CAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCallSummariesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCallSummariesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCallSummariesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCallSummariesResponseBody) GetData() []*ListCallSummariesResponseBodyData {
	return s.Data
}

func (s *ListCallSummariesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCallSummariesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCallSummariesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListCallSummariesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCallSummariesResponseBody) SetCode(v string) *ListCallSummariesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCallSummariesResponseBody) SetData(v []*ListCallSummariesResponseBodyData) *ListCallSummariesResponseBody {
	s.Data = v
	return s
}

func (s *ListCallSummariesResponseBody) SetHttpStatusCode(v int32) *ListCallSummariesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCallSummariesResponseBody) SetMessage(v string) *ListCallSummariesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCallSummariesResponseBody) SetParams(v []*string) *ListCallSummariesResponseBody {
	s.Params = v
	return s
}

func (s *ListCallSummariesResponseBody) SetRequestId(v string) *ListCallSummariesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCallSummariesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCallSummariesResponseBodyData struct {
	// Call ID.
	//
	// example:
	//
	// job-544789******759424
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// Summary information.
	//
	// example:
	//
	// {
	//
	//     "_sys_field_config": [
	//
	//         {
	//
	//             "Array": false,
	//
	//             "Required": true,
	//
	//             "DisplayName": "主题",
	//
	//             "Attributes": "{}",
	//
	//             "DataType": "TEXT",
	//
	//             "EditorType": "TEXT",
	//
	//             "MaxLength": 30,
	//
	//             "System": true,
	//
	//             "Name": "summaryTitle"
	//
	//         },
	//
	//         {
	//
	//             "Array": false,
	//
	//             "DisplayName": "小结",
	//
	//             "Attributes": "{}",
	//
	//             "DataType": "TEXT",
	//
	//             "EditorType": "TEXTAREA",
	//
	//             "MaxLength": 1000,
	//
	//             "System": true,
	//
	//             "Name": "summaryContent"
	//
	//         },
	//
	//         {
	//
	//             "Array": true,
	//
	//             "DisplayName": "标签",
	//
	//             "Attributes": "{}",
	//
	//             "DataType": "TEXT",
	//
	//             "EditorType": "SELECT",
	//
	//             "System": true,
	//
	//             "Name": "keywords"
	//
	//         }
	//
	//     ],
	//
	//     "summaryContent": "小结内容",
	//
	//     "creator": "creator@ccc-test",
	//
	//     "bizType": "Summary",
	//
	//     "keywords": [
	//
	//         "标签1"
	//
	//     ],
	//
	//     "source": "Console",
	//
	//     "summaryTitle": "测试会话小结",
	//
	//     "instanceId": "ccc-test",
	//
	//     "externalInput": "{\\"summaryTitle\\":\\"测试会话小结\\",\\"summaryContent\\":\\"小结内容\\",\\"keywords\\":[\\"标签1\\"],\\"_sys_field_config\\":[{\\"DisplayName\\":\\"主题\\",\\"Name\\":\\"summaryTitle\\",\\"Required\\":true,\\"EditorType\\":\\"TEXT\\",\\"Array\\":false,\\"DataType\\":\\"TEXT\\",\\"System\\":true,\\"MaxLength\\":30,\\"Attributes\\":\\"{}\\"},{\\"DisplayName\\":\\"小结\\",\\"Name\\":\\"summaryContent\\",\\"EditorType\\":\\"TEXTAREA\\",\\"Array\\":false,\\"DataType\\":\\"TEXT\\",\\"System\\":true,\\"MaxLength\\":1000,\\"Attributes\\":\\"{}\\"},{\\"DisplayName\\":\\"标签\\",\\"Name\\":\\"keywords\\",\\"EditorType\\":\\"SELECT\\",\\"Array\\":true,\\"DataType\\":\\"TEXT\\",\\"System\\":true,\\"Attributes\\":\\"{}\\"}]}",
	//
	//     "templateName": "Summary-Ticket",
	//
	//     "bizId": "job-544789******759424",
	//
	//     "ticketStartTime": 1723449513735,
	//
	//     "ticketId": "ac0dd304-****-****-****-4a90010f0d38",
	//
	//     "definitionId": "Summary-Ticket:2:6881f43a-****-11ee-*****-6eb35a90a7a6"
	//
	// }
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1723449513735
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// Creator.
	//
	// example:
	//
	// creator@ccc-test
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// Editor.
	//
	// example:
	//
	// editor@ccc-test
	Editor *string `json:"Editor,omitempty" xml:"Editor,omitempty"`
	// Summary ID.
	//
	// example:
	//
	// ac0dd304-****-****-****-4a90010f0d38
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s ListCallSummariesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCallSummariesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCallSummariesResponseBodyData) GetContactId() *string {
	return s.ContactId
}

func (s *ListCallSummariesResponseBodyData) GetContext() *string {
	return s.Context
}

func (s *ListCallSummariesResponseBodyData) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListCallSummariesResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *ListCallSummariesResponseBodyData) GetEditor() *string {
	return s.Editor
}

func (s *ListCallSummariesResponseBodyData) GetTicketId() *string {
	return s.TicketId
}

func (s *ListCallSummariesResponseBodyData) SetContactId(v string) *ListCallSummariesResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *ListCallSummariesResponseBodyData) SetContext(v string) *ListCallSummariesResponseBodyData {
	s.Context = &v
	return s
}

func (s *ListCallSummariesResponseBodyData) SetCreatedTime(v int64) *ListCallSummariesResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *ListCallSummariesResponseBodyData) SetCreator(v string) *ListCallSummariesResponseBodyData {
	s.Creator = &v
	return s
}

func (s *ListCallSummariesResponseBodyData) SetEditor(v string) *ListCallSummariesResponseBodyData {
	s.Editor = &v
	return s
}

func (s *ListCallSummariesResponseBodyData) SetTicketId(v string) *ListCallSummariesResponseBodyData {
	s.TicketId = &v
	return s
}

func (s *ListCallSummariesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
