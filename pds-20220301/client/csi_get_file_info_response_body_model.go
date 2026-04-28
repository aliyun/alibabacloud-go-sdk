// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCsiGetFileInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvestigationInfo(v *InvestigationInfo) *CsiGetFileInfoResponseBody
	GetInvestigationInfo() *InvestigationInfo
	SetUrl(v string) *CsiGetFileInfoResponseBody
	GetUrl() *string
}

type CsiGetFileInfoResponseBody struct {
	InvestigationInfo *InvestigationInfo `json:"investigation_info,omitempty" xml:"investigation_info,omitempty"`
	// example:
	//
	// https://data.aliyunpds.com/hz22%2F5d5b986facbec311ef844c25954f96821497b383%2F5d5b986f955410dd991646bb87c6b4e899eff525?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CsiGetFileInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CsiGetFileInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CsiGetFileInfoResponseBody) GetInvestigationInfo() *InvestigationInfo {
	return s.InvestigationInfo
}

func (s *CsiGetFileInfoResponseBody) GetUrl() *string {
	return s.Url
}

func (s *CsiGetFileInfoResponseBody) SetInvestigationInfo(v *InvestigationInfo) *CsiGetFileInfoResponseBody {
	s.InvestigationInfo = v
	return s
}

func (s *CsiGetFileInfoResponseBody) SetUrl(v string) *CsiGetFileInfoResponseBody {
	s.Url = &v
	return s
}

func (s *CsiGetFileInfoResponseBody) Validate() error {
	if s.InvestigationInfo != nil {
		if err := s.InvestigationInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
