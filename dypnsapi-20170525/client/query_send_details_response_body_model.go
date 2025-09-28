// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySendDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QuerySendDetailsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QuerySendDetailsResponseBody
	GetCode() *string
	SetMessage(v string) *QuerySendDetailsResponseBody
	GetMessage() *string
	SetModel(v []*QuerySendDetailsResponseBodyModel) *QuerySendDetailsResponseBody
	GetModel() []*QuerySendDetailsResponseBodyModel
	SetSuccess(v bool) *QuerySendDetailsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *QuerySendDetailsResponseBody
	GetTotalCount() *int64
}

type QuerySendDetailsResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// none
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// If OK is returned, the request is successful. Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html?spm=a2c4g.419277.0.i11).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned data.
	Model []*QuerySendDetailsResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries in the list.
	//
	// example:
	//
	// 42
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QuerySendDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySendDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySendDetailsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QuerySendDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySendDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySendDetailsResponseBody) GetModel() []*QuerySendDetailsResponseBodyModel {
	return s.Model
}

func (s *QuerySendDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySendDetailsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QuerySendDetailsResponseBody) SetAccessDeniedDetail(v string) *QuerySendDetailsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySendDetailsResponseBody) SetCode(v string) *QuerySendDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySendDetailsResponseBody) SetMessage(v string) *QuerySendDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySendDetailsResponseBody) SetModel(v []*QuerySendDetailsResponseBodyModel) *QuerySendDetailsResponseBody {
	s.Model = v
	return s
}

func (s *QuerySendDetailsResponseBody) SetSuccess(v bool) *QuerySendDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySendDetailsResponseBody) SetTotalCount(v int64) *QuerySendDetailsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QuerySendDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}

type QuerySendDetailsResponseBodyModel struct {
	// The content of the text message.
	//
	// example:
	//
	// 203160
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The status code returned by the carrier.
	//
	// 	- If the text message was delivered, "DELIVERED" is returned.
	//
	// 	- If the text message failed to be sent, see [Error codes](https://help.aliyun.com/document_detail/101347.html?spm=a2c4g.419277.0.i8) for more information.
	//
	// example:
	//
	// DELIVERED
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The extension field.
	//
	// example:
	//
	// 12131231
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 1390000****
	PhoneNum *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	// The date and time when the text message was received.
	//
	// example:
	//
	// 2019-01-08 16:44:13
	ReceiveDate *string `json:"ReceiveDate,omitempty" xml:"ReceiveDate,omitempty"`
	// The date when the text message was sent. You can query text messages that were sent within the last 30 days.
	//
	// The date is in the yyyyMMdd format. Example: 20181225.
	//
	// example:
	//
	// 2019-01-08 16:44:13
	SendDate *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
	// The delivery status of the text message.
	//
	// 	- 1: A delivery receipt is to be sent.
	//
	// 	- 2: The text message failed to be sent.
	//
	// 	- 3: The text message was sent.
	//
	// example:
	//
	// 3
	SendStatus *int64 `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	// The code of the text message template.
	//
	// Log on to the SMS console. In the left-side navigation pane, click **Go China*	- or **Go Globe**. You can view the text message template code in the **Template Code*	- column on the **Message Templates*	- tab.
	//
	// >  The text message templates must be created on the Go Globe page and approved.
	//
	// example:
	//
	// SMS_12231****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s QuerySendDetailsResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s QuerySendDetailsResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QuerySendDetailsResponseBodyModel) GetContent() *string {
	return s.Content
}

func (s *QuerySendDetailsResponseBodyModel) GetErrCode() *string {
	return s.ErrCode
}

func (s *QuerySendDetailsResponseBodyModel) GetOutId() *string {
	return s.OutId
}

func (s *QuerySendDetailsResponseBodyModel) GetPhoneNum() *string {
	return s.PhoneNum
}

func (s *QuerySendDetailsResponseBodyModel) GetReceiveDate() *string {
	return s.ReceiveDate
}

func (s *QuerySendDetailsResponseBodyModel) GetSendDate() *string {
	return s.SendDate
}

func (s *QuerySendDetailsResponseBodyModel) GetSendStatus() *int64 {
	return s.SendStatus
}

func (s *QuerySendDetailsResponseBodyModel) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *QuerySendDetailsResponseBodyModel) SetContent(v string) *QuerySendDetailsResponseBodyModel {
	s.Content = &v
	return s
}

func (s *QuerySendDetailsResponseBodyModel) SetErrCode(v string) *QuerySendDetailsResponseBodyModel {
	s.ErrCode = &v
	return s
}

func (s *QuerySendDetailsResponseBodyModel) SetOutId(v string) *QuerySendDetailsResponseBodyModel {
	s.OutId = &v
	return s
}

func (s *QuerySendDetailsResponseBodyModel) SetPhoneNum(v string) *QuerySendDetailsResponseBodyModel {
	s.PhoneNum = &v
	return s
}

func (s *QuerySendDetailsResponseBodyModel) SetReceiveDate(v string) *QuerySendDetailsResponseBodyModel {
	s.ReceiveDate = &v
	return s
}

func (s *QuerySendDetailsResponseBodyModel) SetSendDate(v string) *QuerySendDetailsResponseBodyModel {
	s.SendDate = &v
	return s
}

func (s *QuerySendDetailsResponseBodyModel) SetSendStatus(v int64) *QuerySendDetailsResponseBodyModel {
	s.SendStatus = &v
	return s
}

func (s *QuerySendDetailsResponseBodyModel) SetTemplateCode(v string) *QuerySendDetailsResponseBodyModel {
	s.TemplateCode = &v
	return s
}

func (s *QuerySendDetailsResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
