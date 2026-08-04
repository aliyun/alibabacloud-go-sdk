// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindAllContacterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FindAllContacterResponseBody
	GetCode() *string
	SetData(v *FindAllContacterResponseBodyData) *FindAllContacterResponseBody
	GetData() *FindAllContacterResponseBodyData
	SetMessage(v string) *FindAllContacterResponseBody
	GetMessage() *string
	SetRequestId(v string) *FindAllContacterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FindAllContacterResponseBody
	GetSuccess() *bool
}

type FindAllContacterResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *FindAllContacterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FindAllContacterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FindAllContacterResponseBody) GoString() string {
	return s.String()
}

func (s *FindAllContacterResponseBody) GetCode() *string {
	return s.Code
}

func (s *FindAllContacterResponseBody) GetData() *FindAllContacterResponseBodyData {
	return s.Data
}

func (s *FindAllContacterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FindAllContacterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FindAllContacterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FindAllContacterResponseBody) SetCode(v string) *FindAllContacterResponseBody {
	s.Code = &v
	return s
}

func (s *FindAllContacterResponseBody) SetData(v *FindAllContacterResponseBodyData) *FindAllContacterResponseBody {
	s.Data = v
	return s
}

func (s *FindAllContacterResponseBody) SetMessage(v string) *FindAllContacterResponseBody {
	s.Message = &v
	return s
}

func (s *FindAllContacterResponseBody) SetRequestId(v string) *FindAllContacterResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindAllContacterResponseBody) SetSuccess(v bool) *FindAllContacterResponseBody {
	s.Success = &v
	return s
}

func (s *FindAllContacterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FindAllContacterResponseBodyData struct {
	ContacterInfo []*FindAllContacterResponseBodyDataContacterInfo `json:"ContacterInfo,omitempty" xml:"ContacterInfo,omitempty" type:"Repeated"`
}

func (s FindAllContacterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FindAllContacterResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindAllContacterResponseBodyData) GetContacterInfo() []*FindAllContacterResponseBodyDataContacterInfo {
	return s.ContacterInfo
}

func (s *FindAllContacterResponseBodyData) SetContacterInfo(v []*FindAllContacterResponseBodyDataContacterInfo) *FindAllContacterResponseBodyData {
	s.ContacterInfo = v
	return s
}

func (s *FindAllContacterResponseBodyData) Validate() error {
	if s.ContacterInfo != nil {
		for _, item := range s.ContacterInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type FindAllContacterResponseBodyDataContacterInfo struct {
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

func (s FindAllContacterResponseBodyDataContacterInfo) String() string {
	return dara.Prettify(s)
}

func (s FindAllContacterResponseBodyDataContacterInfo) GoString() string {
	return s.String()
}

func (s *FindAllContacterResponseBodyDataContacterInfo) GetContacterAddress() *string {
	return s.ContacterAddress
}

func (s *FindAllContacterResponseBodyDataContacterInfo) GetContacterDingding() *string {
	return s.ContacterDingding
}

func (s *FindAllContacterResponseBodyDataContacterInfo) GetContacterEmail() *string {
	return s.ContacterEmail
}

func (s *FindAllContacterResponseBodyDataContacterInfo) GetContacterId() *int64 {
	return s.ContacterId
}

func (s *FindAllContacterResponseBodyDataContacterInfo) GetContacterMobile() *string {
	return s.ContacterMobile
}

func (s *FindAllContacterResponseBodyDataContacterInfo) GetContacterName() *string {
	return s.ContacterName
}

func (s *FindAllContacterResponseBodyDataContacterInfo) GetContacterPosition() *string {
	return s.ContacterPosition
}

func (s *FindAllContacterResponseBodyDataContacterInfo) GetContacterStaffNo() *string {
	return s.ContacterStaffNo
}

func (s *FindAllContacterResponseBodyDataContacterInfo) GetContacterType() *string {
	return s.ContacterType
}

func (s *FindAllContacterResponseBodyDataContacterInfo) GetContacterWangwang() *string {
	return s.ContacterWangwang
}

func (s *FindAllContacterResponseBodyDataContacterInfo) GetEmailConfirmed() *bool {
	return s.EmailConfirmed
}

func (s *FindAllContacterResponseBodyDataContacterInfo) GetMobileConfirmed() *bool {
	return s.MobileConfirmed
}

func (s *FindAllContacterResponseBodyDataContacterInfo) SetContacterAddress(v string) *FindAllContacterResponseBodyDataContacterInfo {
	s.ContacterAddress = &v
	return s
}

func (s *FindAllContacterResponseBodyDataContacterInfo) SetContacterDingding(v string) *FindAllContacterResponseBodyDataContacterInfo {
	s.ContacterDingding = &v
	return s
}

func (s *FindAllContacterResponseBodyDataContacterInfo) SetContacterEmail(v string) *FindAllContacterResponseBodyDataContacterInfo {
	s.ContacterEmail = &v
	return s
}

func (s *FindAllContacterResponseBodyDataContacterInfo) SetContacterId(v int64) *FindAllContacterResponseBodyDataContacterInfo {
	s.ContacterId = &v
	return s
}

func (s *FindAllContacterResponseBodyDataContacterInfo) SetContacterMobile(v string) *FindAllContacterResponseBodyDataContacterInfo {
	s.ContacterMobile = &v
	return s
}

func (s *FindAllContacterResponseBodyDataContacterInfo) SetContacterName(v string) *FindAllContacterResponseBodyDataContacterInfo {
	s.ContacterName = &v
	return s
}

func (s *FindAllContacterResponseBodyDataContacterInfo) SetContacterPosition(v string) *FindAllContacterResponseBodyDataContacterInfo {
	s.ContacterPosition = &v
	return s
}

func (s *FindAllContacterResponseBodyDataContacterInfo) SetContacterStaffNo(v string) *FindAllContacterResponseBodyDataContacterInfo {
	s.ContacterStaffNo = &v
	return s
}

func (s *FindAllContacterResponseBodyDataContacterInfo) SetContacterType(v string) *FindAllContacterResponseBodyDataContacterInfo {
	s.ContacterType = &v
	return s
}

func (s *FindAllContacterResponseBodyDataContacterInfo) SetContacterWangwang(v string) *FindAllContacterResponseBodyDataContacterInfo {
	s.ContacterWangwang = &v
	return s
}

func (s *FindAllContacterResponseBodyDataContacterInfo) SetEmailConfirmed(v bool) *FindAllContacterResponseBodyDataContacterInfo {
	s.EmailConfirmed = &v
	return s
}

func (s *FindAllContacterResponseBodyDataContacterInfo) SetMobileConfirmed(v bool) *FindAllContacterResponseBodyDataContacterInfo {
	s.MobileConfirmed = &v
	return s
}

func (s *FindAllContacterResponseBodyDataContacterInfo) Validate() error {
	return dara.Validate(s)
}
