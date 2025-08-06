// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDRMCertInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDRMCertInfoList(v []*ListDRMCertInfoResponseBodyDRMCertInfoList) *ListDRMCertInfoResponseBody
	GetDRMCertInfoList() []*ListDRMCertInfoResponseBodyDRMCertInfoList
	SetRequestId(v string) *ListDRMCertInfoResponseBody
	GetRequestId() *string
}

type ListDRMCertInfoResponseBody struct {
	DRMCertInfoList []*ListDRMCertInfoResponseBodyDRMCertInfoList `json:"DRMCertInfoList,omitempty" xml:"DRMCertInfoList,omitempty" type:"Repeated"`
	RequestId       *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDRMCertInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDRMCertInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListDRMCertInfoResponseBody) GetDRMCertInfoList() []*ListDRMCertInfoResponseBodyDRMCertInfoList {
	return s.DRMCertInfoList
}

func (s *ListDRMCertInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDRMCertInfoResponseBody) SetDRMCertInfoList(v []*ListDRMCertInfoResponseBodyDRMCertInfoList) *ListDRMCertInfoResponseBody {
	s.DRMCertInfoList = v
	return s
}

func (s *ListDRMCertInfoResponseBody) SetRequestId(v string) *ListDRMCertInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDRMCertInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDRMCertInfoResponseBodyDRMCertInfoList struct {
	Ask          *string `json:"Ask,omitempty" xml:"Ask,omitempty"`
	CertId       *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertName     *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DRMType      *string `json:"DRMType,omitempty" xml:"DRMType,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PassPhrase   *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
	PrivateKey   *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	ServCert     *string `json:"ServCert,omitempty" xml:"ServCert,omitempty"`
}

func (s ListDRMCertInfoResponseBodyDRMCertInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListDRMCertInfoResponseBodyDRMCertInfoList) GoString() string {
	return s.String()
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) GetAsk() *string {
	return s.Ask
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) GetCertId() *string {
	return s.CertId
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) GetCertName() *string {
	return s.CertName
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) GetDRMType() *string {
	return s.DRMType
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) GetDescription() *string {
	return s.Description
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) GetPassPhrase() *string {
	return s.PassPhrase
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) GetServCert() *string {
	return s.ServCert
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) SetAsk(v string) *ListDRMCertInfoResponseBodyDRMCertInfoList {
	s.Ask = &v
	return s
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) SetCertId(v string) *ListDRMCertInfoResponseBodyDRMCertInfoList {
	s.CertId = &v
	return s
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) SetCertName(v string) *ListDRMCertInfoResponseBodyDRMCertInfoList {
	s.CertName = &v
	return s
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) SetCreationTime(v string) *ListDRMCertInfoResponseBodyDRMCertInfoList {
	s.CreationTime = &v
	return s
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) SetDRMType(v string) *ListDRMCertInfoResponseBodyDRMCertInfoList {
	s.DRMType = &v
	return s
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) SetDescription(v string) *ListDRMCertInfoResponseBodyDRMCertInfoList {
	s.Description = &v
	return s
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) SetPassPhrase(v string) *ListDRMCertInfoResponseBodyDRMCertInfoList {
	s.PassPhrase = &v
	return s
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) SetPrivateKey(v string) *ListDRMCertInfoResponseBodyDRMCertInfoList {
	s.PrivateKey = &v
	return s
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) SetServCert(v string) *ListDRMCertInfoResponseBodyDRMCertInfoList {
	s.ServCert = &v
	return s
}

func (s *ListDRMCertInfoResponseBodyDRMCertInfoList) Validate() error {
	return dara.Validate(s)
}
