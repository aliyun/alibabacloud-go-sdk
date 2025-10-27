// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMaliciousFileWhitelistConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMaliciousFileWhitelistConfigResponseBody
	GetCode() *string
	SetData(v *GetMaliciousFileWhitelistConfigResponseBodyData) *GetMaliciousFileWhitelistConfigResponseBody
	GetData() *GetMaliciousFileWhitelistConfigResponseBodyData
	SetHttpStatusCode(v int32) *GetMaliciousFileWhitelistConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMaliciousFileWhitelistConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMaliciousFileWhitelistConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMaliciousFileWhitelistConfigResponseBody
	GetSuccess() *bool
}

type GetMaliciousFileWhitelistConfigResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *GetMaliciousFileWhitelistConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code. The value 200 indicates that the request was successful.
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
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
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

func (s GetMaliciousFileWhitelistConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMaliciousFileWhitelistConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetMaliciousFileWhitelistConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMaliciousFileWhitelistConfigResponseBody) GetData() *GetMaliciousFileWhitelistConfigResponseBodyData {
	return s.Data
}

func (s *GetMaliciousFileWhitelistConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMaliciousFileWhitelistConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMaliciousFileWhitelistConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMaliciousFileWhitelistConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMaliciousFileWhitelistConfigResponseBody) SetCode(v string) *GetMaliciousFileWhitelistConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponseBody) SetData(v *GetMaliciousFileWhitelistConfigResponseBodyData) *GetMaliciousFileWhitelistConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponseBody) SetHttpStatusCode(v int32) *GetMaliciousFileWhitelistConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponseBody) SetMessage(v string) *GetMaliciousFileWhitelistConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponseBody) SetRequestId(v string) *GetMaliciousFileWhitelistConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponseBody) SetSuccess(v bool) *GetMaliciousFileWhitelistConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMaliciousFileWhitelistConfigResponseBodyData struct {
	// The number of assets on which the whitelist rule takes effect.
	//
	// >  The value of this parameter is returned only if the value of TargetType is SELECTION_KEY.
	//
	// example:
	//
	// 0
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
	// 1674095396000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the whitelist rule.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s GetMaliciousFileWhitelistConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMaliciousFileWhitelistConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) GetCount() *string {
	return s.Count
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) GetEventName() *string {
	return s.EventName
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) GetField() *string {
	return s.Field
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) GetFieldValue() *string {
	return s.FieldValue
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) GetOperator() *string {
	return s.Operator
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) GetTargetType() *string {
	return s.TargetType
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) GetTargetValue() *string {
	return s.TargetValue
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) SetCount(v string) *GetMaliciousFileWhitelistConfigResponseBodyData {
	s.Count = &v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) SetEventName(v string) *GetMaliciousFileWhitelistConfigResponseBodyData {
	s.EventName = &v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) SetField(v string) *GetMaliciousFileWhitelistConfigResponseBodyData {
	s.Field = &v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) SetFieldValue(v string) *GetMaliciousFileWhitelistConfigResponseBodyData {
	s.FieldValue = &v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) SetGmtCreate(v string) *GetMaliciousFileWhitelistConfigResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) SetGmtModified(v string) *GetMaliciousFileWhitelistConfigResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) SetId(v int64) *GetMaliciousFileWhitelistConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) SetOperator(v string) *GetMaliciousFileWhitelistConfigResponseBodyData {
	s.Operator = &v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) SetSource(v string) *GetMaliciousFileWhitelistConfigResponseBodyData {
	s.Source = &v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) SetTargetType(v string) *GetMaliciousFileWhitelistConfigResponseBodyData {
	s.TargetType = &v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) SetTargetValue(v string) *GetMaliciousFileWhitelistConfigResponseBodyData {
	s.TargetValue = &v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
