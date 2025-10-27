// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageEventOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetImageEventOperationResponseBody
	GetCode() *string
	SetData(v *GetImageEventOperationResponseBodyData) *GetImageEventOperationResponseBody
	GetData() *GetImageEventOperationResponseBodyData
	SetMessage(v string) *GetImageEventOperationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetImageEventOperationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetImageEventOperationResponseBody
	GetSuccess() *bool
}

type GetImageEventOperationResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetImageEventOperationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 5861EE3E-F0B3-48B8-A5DC-A5080BFB****
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

func (s GetImageEventOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageEventOperationResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageEventOperationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetImageEventOperationResponseBody) GetData() *GetImageEventOperationResponseBodyData {
	return s.Data
}

func (s *GetImageEventOperationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetImageEventOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageEventOperationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetImageEventOperationResponseBody) SetCode(v string) *GetImageEventOperationResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageEventOperationResponseBody) SetData(v *GetImageEventOperationResponseBodyData) *GetImageEventOperationResponseBody {
	s.Data = v
	return s
}

func (s *GetImageEventOperationResponseBody) SetMessage(v string) *GetImageEventOperationResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageEventOperationResponseBody) SetRequestId(v string) *GetImageEventOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageEventOperationResponseBody) SetSuccess(v bool) *GetImageEventOperationResponseBody {
	s.Success = &v
	return s
}

func (s *GetImageEventOperationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageEventOperationResponseBodyData struct {
	// The rule conditions. The value is in the JSON format. Valid values of keys:
	//
	// 	- **condition**: the matching condition.
	//
	// 	- **type**: the matching type.
	//
	// 	- **value**: the matching value.
	//
	// example:
	//
	// [{\\"condition\\": \\"MD5\\", \\"type\\": \\"equals\\", \\"value\\": \\"0083a31cc0083a31ccf7c10367a6e783e\\"}]
	Conditions *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	// The keyword of the alert item.
	//
	// example:
	//
	// PEM
	EventKey *string `json:"EventKey,omitempty" xml:"EventKey,omitempty"`
	// The name of the alert item.
	//
	// example:
	//
	// PEM
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The alert type. Valid values:
	//
	// 	- **sensitiveFile**
	//
	// example:
	//
	// sensitiveFile
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The primary key of the alert handling rule.
	//
	// example:
	//
	// 443496
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The remarks.
	//
	// example:
	//
	// No warning.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The operation code. Valid values:
	//
	// 	- **whitelist**: added to the whitelist.
	//
	// example:
	//
	// whitelist
	OperationCode *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	// The application scope of the rule. The value is in the JSON format. Valid values of keys:
	//
	// 	- **type**
	//
	// 	- **value**
	//
	// example:
	//
	// {\\"type\\": \\"repo\\", \\"value\\": \\"test-aaa/shenzhen-repo-01\\"}
	Scenarios *string `json:"Scenarios,omitempty" xml:"Scenarios,omitempty"`
	// The source of the whitelist. Valid values:
	//
	// 	- **default**: image
	//
	// 	- **agentless**: agentless detection
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetImageEventOperationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetImageEventOperationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImageEventOperationResponseBodyData) GetConditions() *string {
	return s.Conditions
}

func (s *GetImageEventOperationResponseBodyData) GetEventKey() *string {
	return s.EventKey
}

func (s *GetImageEventOperationResponseBodyData) GetEventName() *string {
	return s.EventName
}

func (s *GetImageEventOperationResponseBodyData) GetEventType() *string {
	return s.EventType
}

func (s *GetImageEventOperationResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetImageEventOperationResponseBodyData) GetNote() *string {
	return s.Note
}

func (s *GetImageEventOperationResponseBodyData) GetOperationCode() *string {
	return s.OperationCode
}

func (s *GetImageEventOperationResponseBodyData) GetScenarios() *string {
	return s.Scenarios
}

func (s *GetImageEventOperationResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *GetImageEventOperationResponseBodyData) SetConditions(v string) *GetImageEventOperationResponseBodyData {
	s.Conditions = &v
	return s
}

func (s *GetImageEventOperationResponseBodyData) SetEventKey(v string) *GetImageEventOperationResponseBodyData {
	s.EventKey = &v
	return s
}

func (s *GetImageEventOperationResponseBodyData) SetEventName(v string) *GetImageEventOperationResponseBodyData {
	s.EventName = &v
	return s
}

func (s *GetImageEventOperationResponseBodyData) SetEventType(v string) *GetImageEventOperationResponseBodyData {
	s.EventType = &v
	return s
}

func (s *GetImageEventOperationResponseBodyData) SetId(v int64) *GetImageEventOperationResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetImageEventOperationResponseBodyData) SetNote(v string) *GetImageEventOperationResponseBodyData {
	s.Note = &v
	return s
}

func (s *GetImageEventOperationResponseBodyData) SetOperationCode(v string) *GetImageEventOperationResponseBodyData {
	s.OperationCode = &v
	return s
}

func (s *GetImageEventOperationResponseBodyData) SetScenarios(v string) *GetImageEventOperationResponseBodyData {
	s.Scenarios = &v
	return s
}

func (s *GetImageEventOperationResponseBodyData) SetSource(v string) *GetImageEventOperationResponseBodyData {
	s.Source = &v
	return s
}

func (s *GetImageEventOperationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
