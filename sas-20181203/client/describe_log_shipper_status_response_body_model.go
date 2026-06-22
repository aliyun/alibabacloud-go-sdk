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
	// The log delivery status collection.
	LogShipperStatus *DescribeLogShipperStatusResponseBodyLogShipperStatus `json:"LogShipperStatus,omitempty" xml:"LogShipperStatus,omitempty" type:"Struct"`
	// The ID of the request. Alibaba Cloud generates a unique identifier for each request. You can use the ID to troubleshoot issues.
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
	// The service authorization status of the log analysis feature. Valid values:
	//
	// - **yes**: authorized
	//
	// - **no**: not authorized.
	//
	// example:
	//
	// yes
	AuthStatus *string `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	// The purchase status of the log analysis feature. Valid values:
	//
	// - **yes**: purchased
	//
	// - **no**: not purchased.
	//
	// example:
	//
	// yes
	BuyStatus *string `json:"BuyStatus,omitempty" xml:"BuyStatus,omitempty"`
	// The version of the log delivery fields for log analysis. Valid values:
	//
	// - **SAS_V1**
	//
	// - **SAS_V2**.
	//
	// example:
	//
	// SAS_V1
	EtlMetaVersion *string `json:"EtlMetaVersion,omitempty" xml:"EtlMetaVersion,omitempty"`
	// The enabling status of log analysis. Valid values:
	//
	// - **yes**: enabled
	//
	// - **no**: not enabled.
	//
	// example:
	//
	// yes
	OpenStatus *string `json:"OpenStatus,omitempty" xml:"OpenStatus,omitempty"`
	// The pay-as-you-go activation status of the log analysis feature. Valid values:
	//
	// - **yes**: activated
	//
	// - **no**: not activated.
	//
	// example:
	//
	// no
	PostPaidOpenStatus *string `json:"PostPaidOpenStatus,omitempty" xml:"PostPaidOpenStatus,omitempty"`
	// The pay-as-you-go support status of the log analysis feature. Valid values:
	//
	// - **yes**: supported
	//
	// - **no**: not supported.
	//
	// example:
	//
	// yes
	PostPaidSupportStatus *string `json:"PostPaidSupportStatus,omitempty" xml:"PostPaidSupportStatus,omitempty"`
	// The status of the log project used by the log analysis feature. Valid values:
	//
	// - **Normal**: Normal.
	//
	// - **Disable**: Disabled.
	//
	// example:
	//
	// Normal
	SlsProjectStatus *string `json:"SlsProjectStatus,omitempty" xml:"SlsProjectStatus,omitempty"`
	// The activation status of Simple Log Service (SLS). Valid values:
	//
	// - **yes**: activated
	//
	// - **no**: not activated.
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
