// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEncryptedAccountProfileInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptedProfileInfo(v *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo) *QueryEncryptedAccountProfileInfoResponseBody
	GetEncryptedProfileInfo() *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo
	SetRequestId(v string) *QueryEncryptedAccountProfileInfoResponseBody
	GetRequestId() *string
}

type QueryEncryptedAccountProfileInfoResponseBody struct {
	EncryptedProfileInfo *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo `json:"EncryptedProfileInfo,omitempty" xml:"EncryptedProfileInfo,omitempty" type:"Struct"`
	RequestId            *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryEncryptedAccountProfileInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryEncryptedAccountProfileInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEncryptedAccountProfileInfoResponseBody) GetEncryptedProfileInfo() *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo {
	return s.EncryptedProfileInfo
}

func (s *QueryEncryptedAccountProfileInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryEncryptedAccountProfileInfoResponseBody) SetEncryptedProfileInfo(v *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo) *QueryEncryptedAccountProfileInfoResponseBody {
	s.EncryptedProfileInfo = v
	return s
}

func (s *QueryEncryptedAccountProfileInfoResponseBody) SetRequestId(v string) *QueryEncryptedAccountProfileInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEncryptedAccountProfileInfoResponseBody) Validate() error {
	if s.EncryptedProfileInfo != nil {
		if err := s.EncryptedProfileInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo struct {
	EncryptedAliyunID       *string `json:"EncryptedAliyunID,omitempty" xml:"EncryptedAliyunID,omitempty"`
	EncryptedEmail          *string `json:"EncryptedEmail,omitempty" xml:"EncryptedEmail,omitempty"`
	EncryptedMobile         *string `json:"EncryptedMobile,omitempty" xml:"EncryptedMobile,omitempty"`
	EncryptedSecurityMobile *string `json:"EncryptedSecurityMobile,omitempty" xml:"EncryptedSecurityMobile,omitempty"`
	IsAliyunIdAnEmail       *bool   `json:"IsAliyunIdAnEmail,omitempty" xml:"IsAliyunIdAnEmail,omitempty"`
	Pk                      *string `json:"pk,omitempty" xml:"pk,omitempty"`
}

func (s QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo) GoString() string {
	return s.String()
}

func (s *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo) GetEncryptedAliyunID() *string {
	return s.EncryptedAliyunID
}

func (s *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo) GetEncryptedEmail() *string {
	return s.EncryptedEmail
}

func (s *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo) GetEncryptedMobile() *string {
	return s.EncryptedMobile
}

func (s *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo) GetEncryptedSecurityMobile() *string {
	return s.EncryptedSecurityMobile
}

func (s *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo) GetIsAliyunIdAnEmail() *bool {
	return s.IsAliyunIdAnEmail
}

func (s *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo) GetPk() *string {
	return s.Pk
}

func (s *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo) SetEncryptedAliyunID(v string) *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo {
	s.EncryptedAliyunID = &v
	return s
}

func (s *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo) SetEncryptedEmail(v string) *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo {
	s.EncryptedEmail = &v
	return s
}

func (s *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo) SetEncryptedMobile(v string) *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo {
	s.EncryptedMobile = &v
	return s
}

func (s *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo) SetEncryptedSecurityMobile(v string) *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo {
	s.EncryptedSecurityMobile = &v
	return s
}

func (s *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo) SetIsAliyunIdAnEmail(v bool) *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo {
	s.IsAliyunIdAnEmail = &v
	return s
}

func (s *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo) SetPk(v string) *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo {
	s.Pk = &v
	return s
}

func (s *QueryEncryptedAccountProfileInfoResponseBodyEncryptedProfileInfo) Validate() error {
	return dara.Validate(s)
}
