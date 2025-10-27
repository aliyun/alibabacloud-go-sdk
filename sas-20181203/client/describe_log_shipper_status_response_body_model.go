// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogShipperStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogShipperStatus(v *DescribeLogShipperStatusResponseBodyLogShipperStatus) *DescribeLogShipperStatusResponseBody
	GetLogShipperStatus() *DescribeLogShipperStatusResponseBodyLogShipperStatus
	SetRequestId(v string) *DescribeLogShipperStatusResponseBody
	GetRequestId() *string
}

type DescribeLogShipperStatusResponseBody struct {
	// The status information.
	LogShipperStatus *DescribeLogShipperStatusResponseBodyLogShipperStatus `json:"LogShipperStatus,omitempty" xml:"LogShipperStatus,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLogShipperStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogShipperStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogShipperStatusResponseBody) GetLogShipperStatus() *DescribeLogShipperStatusResponseBodyLogShipperStatus {
	return s.LogShipperStatus
}

func (s *DescribeLogShipperStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogShipperStatusResponseBody) SetLogShipperStatus(v *DescribeLogShipperStatusResponseBodyLogShipperStatus) *DescribeLogShipperStatusResponseBody {
	s.LogShipperStatus = v
	return s
}

func (s *DescribeLogShipperStatusResponseBody) SetRequestId(v string) *DescribeLogShipperStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogShipperStatusResponseBody) Validate() error {
	if s.LogShipperStatus != nil {
		if err := s.LogShipperStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLogShipperStatusResponseBodyLogShipperStatus struct {
	// Indicates whether Security Center is authorized to access Log Service. Valid values:
	//
	// 	- **yes**
	//
	// 	- **no**
	//
	// example:
	//
	// yes
	AuthStatus *string `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	// Indicates whether the log analysis feature is purchased. Valid values:
	//
	// 	- **yes**
	//
	// 	- **no**
	//
	// example:
	//
	// yes
	BuyStatus *string `json:"BuyStatus,omitempty" xml:"BuyStatus,omitempty"`
	// The version of the log analysis field. Valid values:
	//
	// - SAS_V1
	//
	// - SAS_V2
	//
	// example:
	//
	// SAS_V1
	EtlMetaVersion *string `json:"EtlMetaVersion,omitempty" xml:"EtlMetaVersion,omitempty"`
	// The status of the log analysis feature. Valid values:
	//
	// 	- **yes**: enabled
	//
	// 	- **no**: disabled
	//
	// example:
	//
	// yes
	OpenStatus *string `json:"OpenStatus,omitempty" xml:"OpenStatus,omitempty"`
	// Indicates whether the pay-as-you-go billing method is used. Valid values:
	//
	// 	- **yes**
	//
	// 	- **no**
	//
	// example:
	//
	// no
	PostPaidOpenStatus *string `json:"PostPaidOpenStatus,omitempty" xml:"PostPaidOpenStatus,omitempty"`
	// Indicates whether the log analysis feature supports the pay-as-you-go billing method. Valid values:
	//
	// 	- **yes**
	//
	// 	- **no**
	//
	// example:
	//
	// yes
	PostPaidSupportStatus *string `json:"PostPaidSupportStatus,omitempty" xml:"PostPaidSupportStatus,omitempty"`
	// The status of the dedicated Log Service project. Valid values:
	//
	// 	- **Normal**: normal
	//
	// 	- **Disable**: disabled
	//
	// example:
	//
	// Normal
	SlsProjectStatus *string `json:"SlsProjectStatus,omitempty" xml:"SlsProjectStatus,omitempty"`
	// Indicates whether Log Service is activated. Valid values:
	//
	// 	- **yes**
	//
	// 	- **no**
	//
	// example:
	//
	// yes
	SlsServiceStatus *string `json:"SlsServiceStatus,omitempty" xml:"SlsServiceStatus,omitempty"`
}

func (s DescribeLogShipperStatusResponseBodyLogShipperStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogShipperStatusResponseBodyLogShipperStatus) GoString() string {
	return s.String()
}

func (s *DescribeLogShipperStatusResponseBodyLogShipperStatus) GetAuthStatus() *string {
	return s.AuthStatus
}

func (s *DescribeLogShipperStatusResponseBodyLogShipperStatus) GetBuyStatus() *string {
	return s.BuyStatus
}

func (s *DescribeLogShipperStatusResponseBodyLogShipperStatus) GetEtlMetaVersion() *string {
	return s.EtlMetaVersion
}

func (s *DescribeLogShipperStatusResponseBodyLogShipperStatus) GetOpenStatus() *string {
	return s.OpenStatus
}

func (s *DescribeLogShipperStatusResponseBodyLogShipperStatus) GetPostPaidOpenStatus() *string {
	return s.PostPaidOpenStatus
}

func (s *DescribeLogShipperStatusResponseBodyLogShipperStatus) GetPostPaidSupportStatus() *string {
	return s.PostPaidSupportStatus
}

func (s *DescribeLogShipperStatusResponseBodyLogShipperStatus) GetSlsProjectStatus() *string {
	return s.SlsProjectStatus
}

func (s *DescribeLogShipperStatusResponseBodyLogShipperStatus) GetSlsServiceStatus() *string {
	return s.SlsServiceStatus
}

func (s *DescribeLogShipperStatusResponseBodyLogShipperStatus) SetAuthStatus(v string) *DescribeLogShipperStatusResponseBodyLogShipperStatus {
	s.AuthStatus = &v
	return s
}

func (s *DescribeLogShipperStatusResponseBodyLogShipperStatus) SetBuyStatus(v string) *DescribeLogShipperStatusResponseBodyLogShipperStatus {
	s.BuyStatus = &v
	return s
}

func (s *DescribeLogShipperStatusResponseBodyLogShipperStatus) SetEtlMetaVersion(v string) *DescribeLogShipperStatusResponseBodyLogShipperStatus {
	s.EtlMetaVersion = &v
	return s
}

func (s *DescribeLogShipperStatusResponseBodyLogShipperStatus) SetOpenStatus(v string) *DescribeLogShipperStatusResponseBodyLogShipperStatus {
	s.OpenStatus = &v
	return s
}

func (s *DescribeLogShipperStatusResponseBodyLogShipperStatus) SetPostPaidOpenStatus(v string) *DescribeLogShipperStatusResponseBodyLogShipperStatus {
	s.PostPaidOpenStatus = &v
	return s
}

func (s *DescribeLogShipperStatusResponseBodyLogShipperStatus) SetPostPaidSupportStatus(v string) *DescribeLogShipperStatusResponseBodyLogShipperStatus {
	s.PostPaidSupportStatus = &v
	return s
}

func (s *DescribeLogShipperStatusResponseBodyLogShipperStatus) SetSlsProjectStatus(v string) *DescribeLogShipperStatusResponseBodyLogShipperStatus {
	s.SlsProjectStatus = &v
	return s
}

func (s *DescribeLogShipperStatusResponseBodyLogShipperStatus) SetSlsServiceStatus(v string) *DescribeLogShipperStatusResponseBodyLogShipperStatus {
	s.SlsServiceStatus = &v
	return s
}

func (s *DescribeLogShipperStatusResponseBodyLogShipperStatus) Validate() error {
	return dara.Validate(s)
}
