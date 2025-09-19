// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCloudVendorTrialConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthId(v int64) *AddCloudVendorTrialConfigRequest
	GetAuthId() *int64
	SetAuthInfo(v string) *AddCloudVendorTrialConfigRequest
	GetAuthInfo() *string
	SetVendor(v string) *AddCloudVendorTrialConfigRequest
	GetVendor() *string
}

type AddCloudVendorTrialConfigRequest struct {
	// The AccessKey ID.
	//
	// >  [](#-describecloudvendoraccountaklist--authid)You can call the [DescribeCloudVendorAccountAKList](~~DescribeCloudVendorAccountAKList~~) operation to query the AccessKey ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2363
	AuthId *int64 `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	// The configurations of the third-party cloud asset. Valid values:
	//
	// 	- *AWS*: Configure the sqsQueueName and sqsRegion parameters.
	//
	// 	- *Tencent*: Configure the kafkaUserName, kafkaBootstrapServers, and kafkaTopic parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"sqsRegion\\":\\"us-west-2\\",\\"sqsQueueName\\":\\"****\\"}
	AuthInfo *string `json:"AuthInfo,omitempty" xml:"AuthInfo,omitempty"`
	// The service provider of the cloud asset. Valid values:
	//
	// 	- **Tencent**: Tencent Cloud.
	//
	// 	- **AWS**: Amazon Web Services (AWS).
	//
	// This parameter is required.
	//
	// example:
	//
	// Tencent
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s AddCloudVendorTrialConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCloudVendorTrialConfigRequest) GoString() string {
	return s.String()
}

func (s *AddCloudVendorTrialConfigRequest) GetAuthId() *int64 {
	return s.AuthId
}

func (s *AddCloudVendorTrialConfigRequest) GetAuthInfo() *string {
	return s.AuthInfo
}

func (s *AddCloudVendorTrialConfigRequest) GetVendor() *string {
	return s.Vendor
}

func (s *AddCloudVendorTrialConfigRequest) SetAuthId(v int64) *AddCloudVendorTrialConfigRequest {
	s.AuthId = &v
	return s
}

func (s *AddCloudVendorTrialConfigRequest) SetAuthInfo(v string) *AddCloudVendorTrialConfigRequest {
	s.AuthInfo = &v
	return s
}

func (s *AddCloudVendorTrialConfigRequest) SetVendor(v string) *AddCloudVendorTrialConfigRequest {
	s.Vendor = &v
	return s
}

func (s *AddCloudVendorTrialConfigRequest) Validate() error {
	return dara.Validate(s)
}
