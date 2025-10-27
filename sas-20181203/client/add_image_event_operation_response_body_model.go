// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageEventOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddImageEventOperationResponseBody
	GetCode() *string
	SetData(v *AddImageEventOperationResponseBodyData) *AddImageEventOperationResponseBody
	GetData() *AddImageEventOperationResponseBodyData
	SetMessage(v string) *AddImageEventOperationResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddImageEventOperationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddImageEventOperationResponseBody
	GetSuccess() *bool
}

type AddImageEventOperationResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *AddImageEventOperationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 0C8487EF-50C2-54BB-8634-10F8C35D****
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

func (s AddImageEventOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddImageEventOperationResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageEventOperationResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddImageEventOperationResponseBody) GetData() *AddImageEventOperationResponseBodyData {
	return s.Data
}

func (s *AddImageEventOperationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddImageEventOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddImageEventOperationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddImageEventOperationResponseBody) SetCode(v string) *AddImageEventOperationResponseBody {
	s.Code = &v
	return s
}

func (s *AddImageEventOperationResponseBody) SetData(v *AddImageEventOperationResponseBodyData) *AddImageEventOperationResponseBody {
	s.Data = v
	return s
}

func (s *AddImageEventOperationResponseBody) SetMessage(v string) *AddImageEventOperationResponseBody {
	s.Message = &v
	return s
}

func (s *AddImageEventOperationResponseBody) SetRequestId(v string) *AddImageEventOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddImageEventOperationResponseBody) SetSuccess(v bool) *AddImageEventOperationResponseBody {
	s.Success = &v
	return s
}

func (s *AddImageEventOperationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddImageEventOperationResponseBodyData struct {
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
	// The alert type.
	//
	// 	- Only **sensitiveFile*	- may be returned.
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
	// The operation code.
	//
	// 	- Only **whitelist*	- may be returned, which indicates that the alert item is added to the whitelist.
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
}

func (s AddImageEventOperationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddImageEventOperationResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddImageEventOperationResponseBodyData) GetConditions() *string {
	return s.Conditions
}

func (s *AddImageEventOperationResponseBodyData) GetEventKey() *string {
	return s.EventKey
}

func (s *AddImageEventOperationResponseBodyData) GetEventName() *string {
	return s.EventName
}

func (s *AddImageEventOperationResponseBodyData) GetEventType() *string {
	return s.EventType
}

func (s *AddImageEventOperationResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *AddImageEventOperationResponseBodyData) GetOperationCode() *string {
	return s.OperationCode
}

func (s *AddImageEventOperationResponseBodyData) GetScenarios() *string {
	return s.Scenarios
}

func (s *AddImageEventOperationResponseBodyData) SetConditions(v string) *AddImageEventOperationResponseBodyData {
	s.Conditions = &v
	return s
}

func (s *AddImageEventOperationResponseBodyData) SetEventKey(v string) *AddImageEventOperationResponseBodyData {
	s.EventKey = &v
	return s
}

func (s *AddImageEventOperationResponseBodyData) SetEventName(v string) *AddImageEventOperationResponseBodyData {
	s.EventName = &v
	return s
}

func (s *AddImageEventOperationResponseBodyData) SetEventType(v string) *AddImageEventOperationResponseBodyData {
	s.EventType = &v
	return s
}

func (s *AddImageEventOperationResponseBodyData) SetId(v int64) *AddImageEventOperationResponseBodyData {
	s.Id = &v
	return s
}

func (s *AddImageEventOperationResponseBodyData) SetOperationCode(v string) *AddImageEventOperationResponseBodyData {
	s.OperationCode = &v
	return s
}

func (s *AddImageEventOperationResponseBodyData) SetScenarios(v string) *AddImageEventOperationResponseBodyData {
	s.Scenarios = &v
	return s
}

func (s *AddImageEventOperationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
