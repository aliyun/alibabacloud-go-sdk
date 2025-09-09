// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIcpFilingInfoForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetIcpFilingInfoForPartnerResponseBodyData) *GetIcpFilingInfoForPartnerResponseBody
	GetData() *GetIcpFilingInfoForPartnerResponseBodyData
	SetRequestId(v string) *GetIcpFilingInfoForPartnerResponseBody
	GetRequestId() *string
}

type GetIcpFilingInfoForPartnerResponseBody struct {
	Data *GetIcpFilingInfoForPartnerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIcpFilingInfoForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIcpFilingInfoForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *GetIcpFilingInfoForPartnerResponseBody) GetData() *GetIcpFilingInfoForPartnerResponseBodyData {
	return s.Data
}

func (s *GetIcpFilingInfoForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIcpFilingInfoForPartnerResponseBody) SetData(v *GetIcpFilingInfoForPartnerResponseBodyData) *GetIcpFilingInfoForPartnerResponseBody {
	s.Data = v
	return s
}

func (s *GetIcpFilingInfoForPartnerResponseBody) SetRequestId(v string) *GetIcpFilingInfoForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIcpFilingInfoForPartnerResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetIcpFilingInfoForPartnerResponseBodyData struct {
	IcpNumber        *string `json:"IcpNumber,omitempty" xml:"IcpNumber,omitempty"`
	Recorded         *bool   `json:"Recorded,omitempty" xml:"Recorded,omitempty"`
	SiteRecordNumber *string `json:"SiteRecordNumber,omitempty" xml:"SiteRecordNumber,omitempty"`
}

func (s GetIcpFilingInfoForPartnerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetIcpFilingInfoForPartnerResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIcpFilingInfoForPartnerResponseBodyData) GetIcpNumber() *string {
	return s.IcpNumber
}

func (s *GetIcpFilingInfoForPartnerResponseBodyData) GetRecorded() *bool {
	return s.Recorded
}

func (s *GetIcpFilingInfoForPartnerResponseBodyData) GetSiteRecordNumber() *string {
	return s.SiteRecordNumber
}

func (s *GetIcpFilingInfoForPartnerResponseBodyData) SetIcpNumber(v string) *GetIcpFilingInfoForPartnerResponseBodyData {
	s.IcpNumber = &v
	return s
}

func (s *GetIcpFilingInfoForPartnerResponseBodyData) SetRecorded(v bool) *GetIcpFilingInfoForPartnerResponseBodyData {
	s.Recorded = &v
	return s
}

func (s *GetIcpFilingInfoForPartnerResponseBodyData) SetSiteRecordNumber(v string) *GetIcpFilingInfoForPartnerResponseBodyData {
	s.SiteRecordNumber = &v
	return s
}

func (s *GetIcpFilingInfoForPartnerResponseBodyData) Validate() error {
	return dara.Validate(s)
}
