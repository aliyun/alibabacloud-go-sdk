// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindContacterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FindContacterResponseBody
	GetCode() *string
	SetData(v *FindContacterResponseBodyData) *FindContacterResponseBody
	GetData() *FindContacterResponseBodyData
	SetMessage(v string) *FindContacterResponseBody
	GetMessage() *string
	SetRequestId(v string) *FindContacterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FindContacterResponseBody
	GetSuccess() *bool
}

type FindContacterResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *FindContacterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FindContacterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FindContacterResponseBody) GoString() string {
	return s.String()
}

func (s *FindContacterResponseBody) GetCode() *string {
	return s.Code
}

func (s *FindContacterResponseBody) GetData() *FindContacterResponseBodyData {
	return s.Data
}

func (s *FindContacterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FindContacterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FindContacterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FindContacterResponseBody) SetCode(v string) *FindContacterResponseBody {
	s.Code = &v
	return s
}

func (s *FindContacterResponseBody) SetData(v *FindContacterResponseBodyData) *FindContacterResponseBody {
	s.Data = v
	return s
}

func (s *FindContacterResponseBody) SetMessage(v string) *FindContacterResponseBody {
	s.Message = &v
	return s
}

func (s *FindContacterResponseBody) SetRequestId(v string) *FindContacterResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindContacterResponseBody) SetSuccess(v bool) *FindContacterResponseBody {
	s.Success = &v
	return s
}

func (s *FindContacterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FindContacterResponseBodyData struct {
	ContacterAddress  *string `json:"ContacterAddress,omitempty" xml:"ContacterAddress,omitempty"`
	ContacterDingding *string `json:"ContacterDingding,omitempty" xml:"ContacterDingding,omitempty"`
	ContacterEmail    *string `json:"ContacterEmail,omitempty" xml:"ContacterEmail,omitempty"`
	ContacterId       *int64  `json:"ContacterId,omitempty" xml:"ContacterId,omitempty"`
	ContacterMobile   *string `json:"ContacterMobile,omitempty" xml:"ContacterMobile,omitempty"`
	ContacterName     *string `json:"ContacterName,omitempty" xml:"ContacterName,omitempty"`
	ContacterPosition *string `json:"ContacterPosition,omitempty" xml:"ContacterPosition,omitempty"`
	ContacterStaffNo  *string `json:"ContacterStaffNo,omitempty" xml:"ContacterStaffNo,omitempty"`
	ContacterType     *string `json:"ContacterType,omitempty" xml:"ContacterType,omitempty"`
	ContacterWangwang *string `json:"ContacterWangwang,omitempty" xml:"ContacterWangwang,omitempty"`
	EmailConfirmed    *bool   `json:"EmailConfirmed,omitempty" xml:"EmailConfirmed,omitempty"`
	MobileConfirmed   *bool   `json:"MobileConfirmed,omitempty" xml:"MobileConfirmed,omitempty"`
}

func (s FindContacterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FindContacterResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindContacterResponseBodyData) GetContacterAddress() *string {
	return s.ContacterAddress
}

func (s *FindContacterResponseBodyData) GetContacterDingding() *string {
	return s.ContacterDingding
}

func (s *FindContacterResponseBodyData) GetContacterEmail() *string {
	return s.ContacterEmail
}

func (s *FindContacterResponseBodyData) GetContacterId() *int64 {
	return s.ContacterId
}

func (s *FindContacterResponseBodyData) GetContacterMobile() *string {
	return s.ContacterMobile
}

func (s *FindContacterResponseBodyData) GetContacterName() *string {
	return s.ContacterName
}

func (s *FindContacterResponseBodyData) GetContacterPosition() *string {
	return s.ContacterPosition
}

func (s *FindContacterResponseBodyData) GetContacterStaffNo() *string {
	return s.ContacterStaffNo
}

func (s *FindContacterResponseBodyData) GetContacterType() *string {
	return s.ContacterType
}

func (s *FindContacterResponseBodyData) GetContacterWangwang() *string {
	return s.ContacterWangwang
}

func (s *FindContacterResponseBodyData) GetEmailConfirmed() *bool {
	return s.EmailConfirmed
}

func (s *FindContacterResponseBodyData) GetMobileConfirmed() *bool {
	return s.MobileConfirmed
}

func (s *FindContacterResponseBodyData) SetContacterAddress(v string) *FindContacterResponseBodyData {
	s.ContacterAddress = &v
	return s
}

func (s *FindContacterResponseBodyData) SetContacterDingding(v string) *FindContacterResponseBodyData {
	s.ContacterDingding = &v
	return s
}

func (s *FindContacterResponseBodyData) SetContacterEmail(v string) *FindContacterResponseBodyData {
	s.ContacterEmail = &v
	return s
}

func (s *FindContacterResponseBodyData) SetContacterId(v int64) *FindContacterResponseBodyData {
	s.ContacterId = &v
	return s
}

func (s *FindContacterResponseBodyData) SetContacterMobile(v string) *FindContacterResponseBodyData {
	s.ContacterMobile = &v
	return s
}

func (s *FindContacterResponseBodyData) SetContacterName(v string) *FindContacterResponseBodyData {
	s.ContacterName = &v
	return s
}

func (s *FindContacterResponseBodyData) SetContacterPosition(v string) *FindContacterResponseBodyData {
	s.ContacterPosition = &v
	return s
}

func (s *FindContacterResponseBodyData) SetContacterStaffNo(v string) *FindContacterResponseBodyData {
	s.ContacterStaffNo = &v
	return s
}

func (s *FindContacterResponseBodyData) SetContacterType(v string) *FindContacterResponseBodyData {
	s.ContacterType = &v
	return s
}

func (s *FindContacterResponseBodyData) SetContacterWangwang(v string) *FindContacterResponseBodyData {
	s.ContacterWangwang = &v
	return s
}

func (s *FindContacterResponseBodyData) SetEmailConfirmed(v bool) *FindContacterResponseBodyData {
	s.EmailConfirmed = &v
	return s
}

func (s *FindContacterResponseBodyData) SetMobileConfirmed(v bool) *FindContacterResponseBodyData {
	s.MobileConfirmed = &v
	return s
}

func (s *FindContacterResponseBodyData) Validate() error {
	return dara.Validate(s)
}
