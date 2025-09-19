// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudVendorTrialConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthId(v int64) *ModifyCloudVendorTrialConfigRequest
	GetAuthId() *int64
	SetAuthInfo(v string) *ModifyCloudVendorTrialConfigRequest
	GetAuthInfo() *string
	SetDeleteTrail(v bool) *ModifyCloudVendorTrialConfigRequest
	GetDeleteTrail() *bool
	SetVendor(v string) *ModifyCloudVendorTrialConfigRequest
	GetVendor() *string
}

type ModifyCloudVendorTrialConfigRequest struct {
	// The ID of the audit log configuration to be modified.
	//
	// > The ID can be queried via [DescribeCloudVendorAccountAKList](~~DescribeCloudVendorAccountAKList~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// 23**
	AuthId *int64 `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	// Enter the multi-cloud configuration information:
	//
	// - AWS: parameters sqsQueueName, sqsRegion
	//
	// - Tencent: parameters kafkaUserName, kafkaBootstrapServers, kafkaTopic
	//
	// example:
	//
	// {\\"sqsRegion\\":\\"us-west-2\\",\\"sqsQueueName\\":\\"****\\"}
	AuthInfo *string `json:"AuthInfo,omitempty" xml:"AuthInfo,omitempty"`
	// Whether to delete this audit log configuration:
	//
	// - true: Delete
	//
	// - false: Do not delete
	//
	// example:
	//
	// true
	DeleteTrail *bool `json:"DeleteTrail,omitempty" xml:"DeleteTrail,omitempty"`
	// Cloud asset vendor. Values:
	//
	// - **Tencent**: Tencent Cloud
	//
	// - **AWS**: AWS
	//
	// This parameter is required.
	//
	// example:
	//
	// Tencent
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ModifyCloudVendorTrialConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudVendorTrialConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyCloudVendorTrialConfigRequest) GetAuthId() *int64 {
	return s.AuthId
}

func (s *ModifyCloudVendorTrialConfigRequest) GetAuthInfo() *string {
	return s.AuthInfo
}

func (s *ModifyCloudVendorTrialConfigRequest) GetDeleteTrail() *bool {
	return s.DeleteTrail
}

func (s *ModifyCloudVendorTrialConfigRequest) GetVendor() *string {
	return s.Vendor
}

func (s *ModifyCloudVendorTrialConfigRequest) SetAuthId(v int64) *ModifyCloudVendorTrialConfigRequest {
	s.AuthId = &v
	return s
}

func (s *ModifyCloudVendorTrialConfigRequest) SetAuthInfo(v string) *ModifyCloudVendorTrialConfigRequest {
	s.AuthInfo = &v
	return s
}

func (s *ModifyCloudVendorTrialConfigRequest) SetDeleteTrail(v bool) *ModifyCloudVendorTrialConfigRequest {
	s.DeleteTrail = &v
	return s
}

func (s *ModifyCloudVendorTrialConfigRequest) SetVendor(v string) *ModifyCloudVendorTrialConfigRequest {
	s.Vendor = &v
	return s
}

func (s *ModifyCloudVendorTrialConfigRequest) Validate() error {
	return dara.Validate(s)
}
