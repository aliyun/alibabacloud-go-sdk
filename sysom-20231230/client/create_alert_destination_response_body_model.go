// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertDestinationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAlertDestinationResponseBody
	GetCode() *string
	SetData(v *CreateAlertDestinationResponseBodyData) *CreateAlertDestinationResponseBody
	GetData() *CreateAlertDestinationResponseBodyData
	SetMessage(v string) *CreateAlertDestinationResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAlertDestinationResponseBody
	GetRequestId() *string
}

type CreateAlertDestinationResponseBody struct {
	// error code
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Returned data.
	Data *CreateAlertDestinationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Error message
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error message.
	//
	// example:
	//
	// SysomOpenAPIException: SysomOpenAPI.InvalidParameter Invalid params, should be json string or dict
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateAlertDestinationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertDestinationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlertDestinationResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAlertDestinationResponseBody) GetData() *CreateAlertDestinationResponseBodyData {
	return s.Data
}

func (s *CreateAlertDestinationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAlertDestinationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAlertDestinationResponseBody) SetCode(v string) *CreateAlertDestinationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAlertDestinationResponseBody) SetData(v *CreateAlertDestinationResponseBodyData) *CreateAlertDestinationResponseBody {
	s.Data = v
	return s
}

func (s *CreateAlertDestinationResponseBody) SetMessage(v string) *CreateAlertDestinationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAlertDestinationResponseBody) SetRequestId(v string) *CreateAlertDestinationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlertDestinationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAlertDestinationResponseBodyData struct {
	// Creation Time.
	//
	// example:
	//
	// 1753669116286
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Policy ID
	//
	// example:
	//
	// 1
	Id *int32 `json:"id,omitempty" xml:"id,omitempty"`
	// Policy Name
	//
	// example:
	//
	// SysOM
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Configuration Parameter of alert contact
	Params *CreateAlertDestinationResponseBodyDataParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// Configuration Source
	//
	// example:
	//
	// console
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// Push Target. Currently, only DingTalk Robot is supported.
	//
	// example:
	//
	// dingtalk
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
	// User ID
	//
	// example:
	//
	// 1222933234714935
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// Update Time
	//
	// example:
	//
	// 1751254826285
	UpdatedAt *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s CreateAlertDestinationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertDestinationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAlertDestinationResponseBodyData) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *CreateAlertDestinationResponseBodyData) GetId() *int32 {
	return s.Id
}

func (s *CreateAlertDestinationResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateAlertDestinationResponseBodyData) GetParams() *CreateAlertDestinationResponseBodyDataParams {
	return s.Params
}

func (s *CreateAlertDestinationResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *CreateAlertDestinationResponseBodyData) GetTarget() *string {
	return s.Target
}

func (s *CreateAlertDestinationResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *CreateAlertDestinationResponseBodyData) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *CreateAlertDestinationResponseBodyData) SetCreatedAt(v int64) *CreateAlertDestinationResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateAlertDestinationResponseBodyData) SetId(v int32) *CreateAlertDestinationResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateAlertDestinationResponseBodyData) SetName(v string) *CreateAlertDestinationResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateAlertDestinationResponseBodyData) SetParams(v *CreateAlertDestinationResponseBodyDataParams) *CreateAlertDestinationResponseBodyData {
	s.Params = v
	return s
}

func (s *CreateAlertDestinationResponseBodyData) SetSource(v string) *CreateAlertDestinationResponseBodyData {
	s.Source = &v
	return s
}

func (s *CreateAlertDestinationResponseBodyData) SetTarget(v string) *CreateAlertDestinationResponseBodyData {
	s.Target = &v
	return s
}

func (s *CreateAlertDestinationResponseBodyData) SetUid(v string) *CreateAlertDestinationResponseBodyData {
	s.Uid = &v
	return s
}

func (s *CreateAlertDestinationResponseBodyData) SetUpdatedAt(v int64) *CreateAlertDestinationResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *CreateAlertDestinationResponseBodyData) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAlertDestinationResponseBodyDataParams struct {
	// mailbox
	//
	// example:
	//
	// xxx@email.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// Phone
	//
	// example:
	//
	// 1xxx
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// Robot key
	//
	// example:
	//
	// SECxxx
	Sec *string `json:"sec,omitempty" xml:"sec,omitempty"`
	// Robot address
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=xxx
	Webhook *string `json:"webhook,omitempty" xml:"webhook,omitempty"`
}

func (s CreateAlertDestinationResponseBodyDataParams) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertDestinationResponseBodyDataParams) GoString() string {
	return s.String()
}

func (s *CreateAlertDestinationResponseBodyDataParams) GetEmail() *string {
	return s.Email
}

func (s *CreateAlertDestinationResponseBodyDataParams) GetPhone() *string {
	return s.Phone
}

func (s *CreateAlertDestinationResponseBodyDataParams) GetSec() *string {
	return s.Sec
}

func (s *CreateAlertDestinationResponseBodyDataParams) GetWebhook() *string {
	return s.Webhook
}

func (s *CreateAlertDestinationResponseBodyDataParams) SetEmail(v string) *CreateAlertDestinationResponseBodyDataParams {
	s.Email = &v
	return s
}

func (s *CreateAlertDestinationResponseBodyDataParams) SetPhone(v string) *CreateAlertDestinationResponseBodyDataParams {
	s.Phone = &v
	return s
}

func (s *CreateAlertDestinationResponseBodyDataParams) SetSec(v string) *CreateAlertDestinationResponseBodyDataParams {
	s.Sec = &v
	return s
}

func (s *CreateAlertDestinationResponseBodyDataParams) SetWebhook(v string) *CreateAlertDestinationResponseBodyDataParams {
	s.Webhook = &v
	return s
}

func (s *CreateAlertDestinationResponseBodyDataParams) Validate() error {
	return dara.Validate(s)
}
