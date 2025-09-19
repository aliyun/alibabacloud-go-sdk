// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaliciousFileWhitelistConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMaliciousFileWhitelistConfigResponseBody
	GetCode() *string
	SetData(v *CreateMaliciousFileWhitelistConfigResponseBodyData) *CreateMaliciousFileWhitelistConfigResponseBody
	GetData() *CreateMaliciousFileWhitelistConfigResponseBodyData
	SetHttpStatusCode(v int32) *CreateMaliciousFileWhitelistConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateMaliciousFileWhitelistConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMaliciousFileWhitelistConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMaliciousFileWhitelistConfigResponseBody
	GetSuccess() *bool
}

type CreateMaliciousFileWhitelistConfigResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *CreateMaliciousFileWhitelistConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	// A4EB8B1C-1DEC-5E18-BCD0-XXXXXXXXX
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

func (s CreateMaliciousFileWhitelistConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMaliciousFileWhitelistConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMaliciousFileWhitelistConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMaliciousFileWhitelistConfigResponseBody) GetData() *CreateMaliciousFileWhitelistConfigResponseBodyData {
	return s.Data
}

func (s *CreateMaliciousFileWhitelistConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateMaliciousFileWhitelistConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMaliciousFileWhitelistConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMaliciousFileWhitelistConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMaliciousFileWhitelistConfigResponseBody) SetCode(v string) *CreateMaliciousFileWhitelistConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponseBody) SetData(v *CreateMaliciousFileWhitelistConfigResponseBodyData) *CreateMaliciousFileWhitelistConfigResponseBody {
	s.Data = v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponseBody) SetHttpStatusCode(v int32) *CreateMaliciousFileWhitelistConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponseBody) SetMessage(v string) *CreateMaliciousFileWhitelistConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponseBody) SetRequestId(v string) *CreateMaliciousFileWhitelistConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponseBody) SetSuccess(v bool) *CreateMaliciousFileWhitelistConfigResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateMaliciousFileWhitelistConfigResponseBodyData struct {
	// The number of the assets on which the whitelist rule takes effect.
	//
	// >  The value of this parameter is returned only if the value of TargetType is SELECTION_KEY.
	//
	// example:
	//
	// 1
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the alert.
	//
	// 	- The value is fixed as ALL, which indicates all alert types.
	//
	// example:
	//
	// ALL
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The field that is used in the whitelist rule.
	//
	// example:
	//
	// fileMd5
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The value of the field that is used in the whitelist rule.
	//
	// example:
	//
	// b2cf9747ee49d8d9b105cf16e078cc16
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1671607025000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 1671607025000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the whitelist rule.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The logical operator that is used in the whitelist rule.
	//
	// 	- The value is fixed as strEqual, which indicates the equality operator (=).
	//
	// example:
	//
	// strEqual
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The feature to which this operation belongs.
	//
	// 	- The value is fixed as agentless, which indicates the agentless detection feature.
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the assets on which the whitelist rule takes effect. Valid values:
	//
	// 	- ALL: all assets
	//
	// 	- SELECTION_KEY: selected assets
	//
	// example:
	//
	// ALL
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The assets on which the whitelist rule takes effect. Valid values:
	//
	// 	- ALL: all assets
	//
	// 	- Others: selected assets
	//
	// example:
	//
	// ALL
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s CreateMaliciousFileWhitelistConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMaliciousFileWhitelistConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) GetCount() *string {
	return s.Count
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) GetEventName() *string {
	return s.EventName
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) GetField() *string {
	return s.Field
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) GetFieldValue() *string {
	return s.FieldValue
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) GetId() *string {
	return s.Id
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) GetOperator() *string {
	return s.Operator
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) GetTargetValue() *string {
	return s.TargetValue
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) SetCount(v string) *CreateMaliciousFileWhitelistConfigResponseBodyData {
	s.Count = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) SetEventName(v string) *CreateMaliciousFileWhitelistConfigResponseBodyData {
	s.EventName = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) SetField(v string) *CreateMaliciousFileWhitelistConfigResponseBodyData {
	s.Field = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) SetFieldValue(v string) *CreateMaliciousFileWhitelistConfigResponseBodyData {
	s.FieldValue = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) SetGmtCreate(v string) *CreateMaliciousFileWhitelistConfigResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) SetGmtModified(v string) *CreateMaliciousFileWhitelistConfigResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) SetId(v string) *CreateMaliciousFileWhitelistConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) SetOperator(v string) *CreateMaliciousFileWhitelistConfigResponseBodyData {
	s.Operator = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) SetSource(v string) *CreateMaliciousFileWhitelistConfigResponseBodyData {
	s.Source = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) SetTargetType(v string) *CreateMaliciousFileWhitelistConfigResponseBodyData {
	s.TargetType = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) SetTargetValue(v string) *CreateMaliciousFileWhitelistConfigResponseBodyData {
	s.TargetValue = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
