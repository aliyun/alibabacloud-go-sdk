// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRiskCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunAssetCertificateExpiredCount(v int32) *GetRiskCountResponseBody
	GetAliyunAssetCertificateExpiredCount() *int32
	SetAliyunAssetCertificateWillExpiredCount(v int32) *GetRiskCountResponseBody
	GetAliyunAssetCertificateWillExpiredCount() *int32
	SetBuyCertificateExpireCount(v int32) *GetRiskCountResponseBody
	GetBuyCertificateExpireCount() *int32
	SetBuyCertificateNotDeploymentCount(v int32) *GetRiskCountResponseBody
	GetBuyCertificateNotDeploymentCount() *int32
	SetBuyCertificateNotTrusteeCount(v int32) *GetRiskCountResponseBody
	GetBuyCertificateNotTrusteeCount() *int32
	SetBuyCertificateWillExpiredCount(v int32) *GetRiskCountResponseBody
	GetBuyCertificateWillExpiredCount() *int32
	SetBuyCheckedFailCount(v int32) *GetRiskCountResponseBody
	GetBuyCheckedFailCount() *int32
	SetDomainAssetNotMonitorCount(v int32) *GetRiskCountResponseBody
	GetDomainAssetNotMonitorCount() *int32
	SetFreeCertificateExpireCount(v int32) *GetRiskCountResponseBody
	GetFreeCertificateExpireCount() *int32
	SetFreeCertificateNotDeploymentCount(v int32) *GetRiskCountResponseBody
	GetFreeCertificateNotDeploymentCount() *int32
	SetFreeCertificateWillExpiredCount(v int32) *GetRiskCountResponseBody
	GetFreeCertificateWillExpiredCount() *int32
	SetFreeCheckedFailCount(v int32) *GetRiskCountResponseBody
	GetFreeCheckedFailCount() *int32
	SetMultiCloudAssetCertificateExpiredCount(v int32) *GetRiskCountResponseBody
	GetMultiCloudAssetCertificateExpiredCount() *int32
	SetMultiCloudAssetCertificateWillExpiredCount(v int32) *GetRiskCountResponseBody
	GetMultiCloudAssetCertificateWillExpiredCount() *int32
	SetRequestId(v string) *GetRiskCountResponseBody
	GetRequestId() *string
	SetUploadCertificateExpireCount(v int32) *GetRiskCountResponseBody
	GetUploadCertificateExpireCount() *int32
	SetUploadCertificateNotDeploymentCount(v int32) *GetRiskCountResponseBody
	GetUploadCertificateNotDeploymentCount() *int32
	SetUploadCertificateNotNoticeCount(v int32) *GetRiskCountResponseBody
	GetUploadCertificateNotNoticeCount() *int32
	SetUploadCertificateWillExpiredCount(v int32) *GetRiskCountResponseBody
	GetUploadCertificateWillExpiredCount() *int32
}

type GetRiskCountResponseBody struct {
	AliyunAssetCertificateExpiredCount         *int32 `json:"AliyunAssetCertificateExpiredCount,omitempty" xml:"AliyunAssetCertificateExpiredCount,omitempty"`
	AliyunAssetCertificateWillExpiredCount     *int32 `json:"AliyunAssetCertificateWillExpiredCount,omitempty" xml:"AliyunAssetCertificateWillExpiredCount,omitempty"`
	BuyCertificateExpireCount                  *int32 `json:"BuyCertificateExpireCount,omitempty" xml:"BuyCertificateExpireCount,omitempty"`
	BuyCertificateNotDeploymentCount           *int32 `json:"BuyCertificateNotDeploymentCount,omitempty" xml:"BuyCertificateNotDeploymentCount,omitempty"`
	BuyCertificateNotTrusteeCount              *int32 `json:"BuyCertificateNotTrusteeCount,omitempty" xml:"BuyCertificateNotTrusteeCount,omitempty"`
	BuyCertificateWillExpiredCount             *int32 `json:"BuyCertificateWillExpiredCount,omitempty" xml:"BuyCertificateWillExpiredCount,omitempty"`
	BuyCheckedFailCount                        *int32 `json:"BuyCheckedFailCount,omitempty" xml:"BuyCheckedFailCount,omitempty"`
	DomainAssetNotMonitorCount                 *int32 `json:"DomainAssetNotMonitorCount,omitempty" xml:"DomainAssetNotMonitorCount,omitempty"`
	FreeCertificateExpireCount                 *int32 `json:"FreeCertificateExpireCount,omitempty" xml:"FreeCertificateExpireCount,omitempty"`
	FreeCertificateNotDeploymentCount          *int32 `json:"FreeCertificateNotDeploymentCount,omitempty" xml:"FreeCertificateNotDeploymentCount,omitempty"`
	FreeCertificateWillExpiredCount            *int32 `json:"FreeCertificateWillExpiredCount,omitempty" xml:"FreeCertificateWillExpiredCount,omitempty"`
	FreeCheckedFailCount                       *int32 `json:"FreeCheckedFailCount,omitempty" xml:"FreeCheckedFailCount,omitempty"`
	MultiCloudAssetCertificateExpiredCount     *int32 `json:"MultiCloudAssetCertificateExpiredCount,omitempty" xml:"MultiCloudAssetCertificateExpiredCount,omitempty"`
	MultiCloudAssetCertificateWillExpiredCount *int32 `json:"MultiCloudAssetCertificateWillExpiredCount,omitempty" xml:"MultiCloudAssetCertificateWillExpiredCount,omitempty"`
	// example:
	//
	// 5BCD2F6C-7A9D-47C1-8588-2CC6A4E0BE5E
	RequestId                           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UploadCertificateExpireCount        *int32  `json:"UploadCertificateExpireCount,omitempty" xml:"UploadCertificateExpireCount,omitempty"`
	UploadCertificateNotDeploymentCount *int32  `json:"UploadCertificateNotDeploymentCount,omitempty" xml:"UploadCertificateNotDeploymentCount,omitempty"`
	UploadCertificateNotNoticeCount     *int32  `json:"UploadCertificateNotNoticeCount,omitempty" xml:"UploadCertificateNotNoticeCount,omitempty"`
	UploadCertificateWillExpiredCount   *int32  `json:"UploadCertificateWillExpiredCount,omitempty" xml:"UploadCertificateWillExpiredCount,omitempty"`
}

func (s GetRiskCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRiskCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetRiskCountResponseBody) GetAliyunAssetCertificateExpiredCount() *int32 {
	return s.AliyunAssetCertificateExpiredCount
}

func (s *GetRiskCountResponseBody) GetAliyunAssetCertificateWillExpiredCount() *int32 {
	return s.AliyunAssetCertificateWillExpiredCount
}

func (s *GetRiskCountResponseBody) GetBuyCertificateExpireCount() *int32 {
	return s.BuyCertificateExpireCount
}

func (s *GetRiskCountResponseBody) GetBuyCertificateNotDeploymentCount() *int32 {
	return s.BuyCertificateNotDeploymentCount
}

func (s *GetRiskCountResponseBody) GetBuyCertificateNotTrusteeCount() *int32 {
	return s.BuyCertificateNotTrusteeCount
}

func (s *GetRiskCountResponseBody) GetBuyCertificateWillExpiredCount() *int32 {
	return s.BuyCertificateWillExpiredCount
}

func (s *GetRiskCountResponseBody) GetBuyCheckedFailCount() *int32 {
	return s.BuyCheckedFailCount
}

func (s *GetRiskCountResponseBody) GetDomainAssetNotMonitorCount() *int32 {
	return s.DomainAssetNotMonitorCount
}

func (s *GetRiskCountResponseBody) GetFreeCertificateExpireCount() *int32 {
	return s.FreeCertificateExpireCount
}

func (s *GetRiskCountResponseBody) GetFreeCertificateNotDeploymentCount() *int32 {
	return s.FreeCertificateNotDeploymentCount
}

func (s *GetRiskCountResponseBody) GetFreeCertificateWillExpiredCount() *int32 {
	return s.FreeCertificateWillExpiredCount
}

func (s *GetRiskCountResponseBody) GetFreeCheckedFailCount() *int32 {
	return s.FreeCheckedFailCount
}

func (s *GetRiskCountResponseBody) GetMultiCloudAssetCertificateExpiredCount() *int32 {
	return s.MultiCloudAssetCertificateExpiredCount
}

func (s *GetRiskCountResponseBody) GetMultiCloudAssetCertificateWillExpiredCount() *int32 {
	return s.MultiCloudAssetCertificateWillExpiredCount
}

func (s *GetRiskCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRiskCountResponseBody) GetUploadCertificateExpireCount() *int32 {
	return s.UploadCertificateExpireCount
}

func (s *GetRiskCountResponseBody) GetUploadCertificateNotDeploymentCount() *int32 {
	return s.UploadCertificateNotDeploymentCount
}

func (s *GetRiskCountResponseBody) GetUploadCertificateNotNoticeCount() *int32 {
	return s.UploadCertificateNotNoticeCount
}

func (s *GetRiskCountResponseBody) GetUploadCertificateWillExpiredCount() *int32 {
	return s.UploadCertificateWillExpiredCount
}

func (s *GetRiskCountResponseBody) SetAliyunAssetCertificateExpiredCount(v int32) *GetRiskCountResponseBody {
	s.AliyunAssetCertificateExpiredCount = &v
	return s
}

func (s *GetRiskCountResponseBody) SetAliyunAssetCertificateWillExpiredCount(v int32) *GetRiskCountResponseBody {
	s.AliyunAssetCertificateWillExpiredCount = &v
	return s
}

func (s *GetRiskCountResponseBody) SetBuyCertificateExpireCount(v int32) *GetRiskCountResponseBody {
	s.BuyCertificateExpireCount = &v
	return s
}

func (s *GetRiskCountResponseBody) SetBuyCertificateNotDeploymentCount(v int32) *GetRiskCountResponseBody {
	s.BuyCertificateNotDeploymentCount = &v
	return s
}

func (s *GetRiskCountResponseBody) SetBuyCertificateNotTrusteeCount(v int32) *GetRiskCountResponseBody {
	s.BuyCertificateNotTrusteeCount = &v
	return s
}

func (s *GetRiskCountResponseBody) SetBuyCertificateWillExpiredCount(v int32) *GetRiskCountResponseBody {
	s.BuyCertificateWillExpiredCount = &v
	return s
}

func (s *GetRiskCountResponseBody) SetBuyCheckedFailCount(v int32) *GetRiskCountResponseBody {
	s.BuyCheckedFailCount = &v
	return s
}

func (s *GetRiskCountResponseBody) SetDomainAssetNotMonitorCount(v int32) *GetRiskCountResponseBody {
	s.DomainAssetNotMonitorCount = &v
	return s
}

func (s *GetRiskCountResponseBody) SetFreeCertificateExpireCount(v int32) *GetRiskCountResponseBody {
	s.FreeCertificateExpireCount = &v
	return s
}

func (s *GetRiskCountResponseBody) SetFreeCertificateNotDeploymentCount(v int32) *GetRiskCountResponseBody {
	s.FreeCertificateNotDeploymentCount = &v
	return s
}

func (s *GetRiskCountResponseBody) SetFreeCertificateWillExpiredCount(v int32) *GetRiskCountResponseBody {
	s.FreeCertificateWillExpiredCount = &v
	return s
}

func (s *GetRiskCountResponseBody) SetFreeCheckedFailCount(v int32) *GetRiskCountResponseBody {
	s.FreeCheckedFailCount = &v
	return s
}

func (s *GetRiskCountResponseBody) SetMultiCloudAssetCertificateExpiredCount(v int32) *GetRiskCountResponseBody {
	s.MultiCloudAssetCertificateExpiredCount = &v
	return s
}

func (s *GetRiskCountResponseBody) SetMultiCloudAssetCertificateWillExpiredCount(v int32) *GetRiskCountResponseBody {
	s.MultiCloudAssetCertificateWillExpiredCount = &v
	return s
}

func (s *GetRiskCountResponseBody) SetRequestId(v string) *GetRiskCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRiskCountResponseBody) SetUploadCertificateExpireCount(v int32) *GetRiskCountResponseBody {
	s.UploadCertificateExpireCount = &v
	return s
}

func (s *GetRiskCountResponseBody) SetUploadCertificateNotDeploymentCount(v int32) *GetRiskCountResponseBody {
	s.UploadCertificateNotDeploymentCount = &v
	return s
}

func (s *GetRiskCountResponseBody) SetUploadCertificateNotNoticeCount(v int32) *GetRiskCountResponseBody {
	s.UploadCertificateNotNoticeCount = &v
	return s
}

func (s *GetRiskCountResponseBody) SetUploadCertificateWillExpiredCount(v int32) *GetRiskCountResponseBody {
	s.UploadCertificateWillExpiredCount = &v
	return s
}

func (s *GetRiskCountResponseBody) Validate() error {
	return dara.Validate(s)
}
