// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCustomerInfoResponseBody
	GetCode() *string
	SetData(v *GetCustomerInfoResponseBodyData) *GetCustomerInfoResponseBody
	GetData() *GetCustomerInfoResponseBodyData
	SetMessage(v string) *GetCustomerInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCustomerInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomerInfoResponseBody
	GetSuccess() *bool
}

type GetCustomerInfoResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Membership information.
	Data *GetCustomerInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// DF6A3FB7-A5AA-43BE-A65B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCustomerInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCustomerInfoResponseBody) GetData() *GetCustomerInfoResponseBodyData {
	return s.Data
}

func (s *GetCustomerInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCustomerInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomerInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomerInfoResponseBody) SetCode(v string) *GetCustomerInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomerInfoResponseBody) SetData(v *GetCustomerInfoResponseBodyData) *GetCustomerInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomerInfoResponseBody) SetMessage(v string) *GetCustomerInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomerInfoResponseBody) SetRequestId(v string) *GetCustomerInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomerInfoResponseBody) SetSuccess(v bool) *GetCustomerInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomerInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomerInfoResponseBodyData struct {
	// Custom fields.
	CustomizeFields map[string]interface{} `json:"CustomizeFields,omitempty" xml:"CustomizeFields,omitempty"`
	// Nickname.
	//
	// example:
	//
	// 测试会员
	Nick *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	// External ID.
	//
	// example:
	//
	// 6666666
	OuterId *string `json:"OuterId,omitempty" xml:"OuterId,omitempty"`
	// Profile picture.
	//
	// example:
	//
	// https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTLSW7XPFlJDwVunXP8pr84TvltwtLlNqTlOVSFeM3bCgn57mAB4JuZZmvMW0qicqW0PyzyUdZpxiaFQ
	Photo *string `json:"Photo,omitempty" xml:"Photo,omitempty"`
	// Real name.
	//
	// example:
	//
	// 刘测试
	RealName *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	// Membership ID.
	//
	// example:
	//
	// 823456789023
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetCustomerInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomerInfoResponseBodyData) GetCustomizeFields() map[string]interface{} {
	return s.CustomizeFields
}

func (s *GetCustomerInfoResponseBodyData) GetNick() *string {
	return s.Nick
}

func (s *GetCustomerInfoResponseBodyData) GetOuterId() *string {
	return s.OuterId
}

func (s *GetCustomerInfoResponseBodyData) GetPhoto() *string {
	return s.Photo
}

func (s *GetCustomerInfoResponseBodyData) GetRealName() *string {
	return s.RealName
}

func (s *GetCustomerInfoResponseBodyData) GetUserId() *int64 {
	return s.UserId
}

func (s *GetCustomerInfoResponseBodyData) SetCustomizeFields(v map[string]interface{}) *GetCustomerInfoResponseBodyData {
	s.CustomizeFields = v
	return s
}

func (s *GetCustomerInfoResponseBodyData) SetNick(v string) *GetCustomerInfoResponseBodyData {
	s.Nick = &v
	return s
}

func (s *GetCustomerInfoResponseBodyData) SetOuterId(v string) *GetCustomerInfoResponseBodyData {
	s.OuterId = &v
	return s
}

func (s *GetCustomerInfoResponseBodyData) SetPhoto(v string) *GetCustomerInfoResponseBodyData {
	s.Photo = &v
	return s
}

func (s *GetCustomerInfoResponseBodyData) SetRealName(v string) *GetCustomerInfoResponseBodyData {
	s.RealName = &v
	return s
}

func (s *GetCustomerInfoResponseBodyData) SetUserId(v int64) *GetCustomerInfoResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetCustomerInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
