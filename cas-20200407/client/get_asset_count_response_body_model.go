// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunAssetCount(v int32) *GetAssetCountResponseBody
	GetAliyunAssetCount() *int32
	SetAwsAssetCount(v int32) *GetAssetCountResponseBody
	GetAwsAssetCount() *int32
	SetBuyCertificateCount(v int32) *GetAssetCountResponseBody
	GetBuyCertificateCount() *int32
	SetDomainAssetCount(v int32) *GetAssetCountResponseBody
	GetDomainAssetCount() *int32
	SetFreeCertificateCount(v int32) *GetAssetCountResponseBody
	GetFreeCertificateCount() *int32
	SetHuaweiAssetCount(v int32) *GetAssetCountResponseBody
	GetHuaweiAssetCount() *int32
	SetLastPoint(v int32) *GetAssetCountResponseBody
	GetLastPoint() *int32
	SetPoint(v int32) *GetAssetCountResponseBody
	GetPoint() *int32
	SetPointRatio(v int32) *GetAssetCountResponseBody
	GetPointRatio() *int32
	SetPointTime(v int64) *GetAssetCountResponseBody
	GetPointTime() *int64
	SetRequestId(v string) *GetAssetCountResponseBody
	GetRequestId() *string
	SetTencentAssetCount(v int32) *GetAssetCountResponseBody
	GetTencentAssetCount() *int32
	SetUploadCertificateCount(v int32) *GetAssetCountResponseBody
	GetUploadCertificateCount() *int32
}

type GetAssetCountResponseBody struct {
	AliyunAssetCount     *int32 `json:"AliyunAssetCount,omitempty" xml:"AliyunAssetCount,omitempty"`
	AwsAssetCount        *int32 `json:"AwsAssetCount,omitempty" xml:"AwsAssetCount,omitempty"`
	BuyCertificateCount  *int32 `json:"BuyCertificateCount,omitempty" xml:"BuyCertificateCount,omitempty"`
	DomainAssetCount     *int32 `json:"DomainAssetCount,omitempty" xml:"DomainAssetCount,omitempty"`
	FreeCertificateCount *int32 `json:"FreeCertificateCount,omitempty" xml:"FreeCertificateCount,omitempty"`
	HuaweiAssetCount     *int32 `json:"HuaweiAssetCount,omitempty" xml:"HuaweiAssetCount,omitempty"`
	LastPoint            *int32 `json:"LastPoint,omitempty" xml:"LastPoint,omitempty"`
	Point                *int32 `json:"Point,omitempty" xml:"Point,omitempty"`
	PointRatio           *int32 `json:"PointRatio,omitempty" xml:"PointRatio,omitempty"`
	PointTime            *int64 `json:"PointTime,omitempty" xml:"PointTime,omitempty"`
	// example:
	//
	// EECA10D5-BD0F-4EF1-B3EA-B4578E5C6F8E
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TencentAssetCount      *int32  `json:"TencentAssetCount,omitempty" xml:"TencentAssetCount,omitempty"`
	UploadCertificateCount *int32  `json:"UploadCertificateCount,omitempty" xml:"UploadCertificateCount,omitempty"`
}

func (s GetAssetCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAssetCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetAssetCountResponseBody) GetAliyunAssetCount() *int32 {
	return s.AliyunAssetCount
}

func (s *GetAssetCountResponseBody) GetAwsAssetCount() *int32 {
	return s.AwsAssetCount
}

func (s *GetAssetCountResponseBody) GetBuyCertificateCount() *int32 {
	return s.BuyCertificateCount
}

func (s *GetAssetCountResponseBody) GetDomainAssetCount() *int32 {
	return s.DomainAssetCount
}

func (s *GetAssetCountResponseBody) GetFreeCertificateCount() *int32 {
	return s.FreeCertificateCount
}

func (s *GetAssetCountResponseBody) GetHuaweiAssetCount() *int32 {
	return s.HuaweiAssetCount
}

func (s *GetAssetCountResponseBody) GetLastPoint() *int32 {
	return s.LastPoint
}

func (s *GetAssetCountResponseBody) GetPoint() *int32 {
	return s.Point
}

func (s *GetAssetCountResponseBody) GetPointRatio() *int32 {
	return s.PointRatio
}

func (s *GetAssetCountResponseBody) GetPointTime() *int64 {
	return s.PointTime
}

func (s *GetAssetCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAssetCountResponseBody) GetTencentAssetCount() *int32 {
	return s.TencentAssetCount
}

func (s *GetAssetCountResponseBody) GetUploadCertificateCount() *int32 {
	return s.UploadCertificateCount
}

func (s *GetAssetCountResponseBody) SetAliyunAssetCount(v int32) *GetAssetCountResponseBody {
	s.AliyunAssetCount = &v
	return s
}

func (s *GetAssetCountResponseBody) SetAwsAssetCount(v int32) *GetAssetCountResponseBody {
	s.AwsAssetCount = &v
	return s
}

func (s *GetAssetCountResponseBody) SetBuyCertificateCount(v int32) *GetAssetCountResponseBody {
	s.BuyCertificateCount = &v
	return s
}

func (s *GetAssetCountResponseBody) SetDomainAssetCount(v int32) *GetAssetCountResponseBody {
	s.DomainAssetCount = &v
	return s
}

func (s *GetAssetCountResponseBody) SetFreeCertificateCount(v int32) *GetAssetCountResponseBody {
	s.FreeCertificateCount = &v
	return s
}

func (s *GetAssetCountResponseBody) SetHuaweiAssetCount(v int32) *GetAssetCountResponseBody {
	s.HuaweiAssetCount = &v
	return s
}

func (s *GetAssetCountResponseBody) SetLastPoint(v int32) *GetAssetCountResponseBody {
	s.LastPoint = &v
	return s
}

func (s *GetAssetCountResponseBody) SetPoint(v int32) *GetAssetCountResponseBody {
	s.Point = &v
	return s
}

func (s *GetAssetCountResponseBody) SetPointRatio(v int32) *GetAssetCountResponseBody {
	s.PointRatio = &v
	return s
}

func (s *GetAssetCountResponseBody) SetPointTime(v int64) *GetAssetCountResponseBody {
	s.PointTime = &v
	return s
}

func (s *GetAssetCountResponseBody) SetRequestId(v string) *GetAssetCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAssetCountResponseBody) SetTencentAssetCount(v int32) *GetAssetCountResponseBody {
	s.TencentAssetCount = &v
	return s
}

func (s *GetAssetCountResponseBody) SetUploadCertificateCount(v int32) *GetAssetCountResponseBody {
	s.UploadCertificateCount = &v
	return s
}

func (s *GetAssetCountResponseBody) Validate() error {
	return dara.Validate(s)
}
