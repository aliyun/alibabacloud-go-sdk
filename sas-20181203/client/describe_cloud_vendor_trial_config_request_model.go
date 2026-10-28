// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudVendorTrialConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthId(v int64) *DescribeCloudVendorTrialConfigRequest
	GetAuthId() *int64
	SetVendor(v string) *DescribeCloudVendorTrialConfigRequest
	GetVendor() *string
}

type DescribeCloudVendorTrialConfigRequest struct {
	// The unique ID of the AccessKey.
	//
	// > You can call [DescribeCloudVendorAccountAKList](~~DescribeCloudVendorAccountAKList~~) to obtain the AuthId.
	//
	// > -
	//
	// This parameter is required.
	//
	// example:
	//
	// 23**
	AuthId *int64 `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	// The cloud asset vendor. Valid values:
	//
	// - **Tencent**: Tencent Cloud
	//
	// - **AWS**: Amazon Web Services
	//
	// This parameter is required.
	//
	// example:
	//
	// AWS
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeCloudVendorTrialConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudVendorTrialConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudVendorTrialConfigRequest) GetAuthId() *int64 {
	return s.AuthId
}

func (s *DescribeCloudVendorTrialConfigRequest) GetVendor() *string {
	return s.Vendor
}

func (s *DescribeCloudVendorTrialConfigRequest) SetAuthId(v int64) *DescribeCloudVendorTrialConfigRequest {
	s.AuthId = &v
	return s
}

func (s *DescribeCloudVendorTrialConfigRequest) SetVendor(v string) *DescribeCloudVendorTrialConfigRequest {
	s.Vendor = &v
	return s
}

func (s *DescribeCloudVendorTrialConfigRequest) Validate() error {
	return dara.Validate(s)
}
