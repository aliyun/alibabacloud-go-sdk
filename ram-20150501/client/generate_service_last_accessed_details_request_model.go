// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateServiceLastAccessedDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *GenerateServiceLastAccessedDetailsRequest
	GetArn() *string
	SetGranularity(v string) *GenerateServiceLastAccessedDetailsRequest
	GetGranularity() *string
}

type GenerateServiceLastAccessedDetailsRequest struct {
	// The Alibaba Cloud Resource Name (ARN) of the RAM identity for which you want to generate a permission audit report. Resource Access Management (RAM) users and regular RAM roles are supported.
	//
	// Format:
	//
	// - Resource Access Management (RAM) user: `acs:ram::<account_id>:user/<user_name>`
	//
	// - RAM role: `acs:ram::<account_id>:role/<role_name>`
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::1234567890:user/test_user
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The granularity of the permission audit report. Valid values:
	//
	// - `SERVICE_LEVEL`: analyzes the permissions and access records of the RAM identity at the cloud service level. You can view the cloud service permissions granted to the RAM identity, the cloud services that the RAM identity has accessed, and the last access time for each cloud service. For more information about the cloud services that support service-level auditing, see [Cloud services that support permission auditing](https://www.alibabacloud.com/help/en/ram/cloud-services-that-support-permission-auditing).
	//
	// - `ACTION_LEVEL`: analyzes the permissions and access records of the RAM identity at the API operation level, which provides a finer granularity than the service level. You can view the specific API operations that the RAM identity is authorized to perform and has performed. For more information about the cloud services that support operation-level auditing, see the cloud services whose **Audit granularity*	- is **Operation level*	- in [Cloud services that support permission auditing](https://www.alibabacloud.com/help/en/ram/cloud-services-that-support-permission-auditing).
	//
	// example:
	//
	// SERVICE_LEVEL
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
}

func (s GenerateServiceLastAccessedDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateServiceLastAccessedDetailsRequest) GoString() string {
	return s.String()
}

func (s *GenerateServiceLastAccessedDetailsRequest) GetArn() *string {
	return s.Arn
}

func (s *GenerateServiceLastAccessedDetailsRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *GenerateServiceLastAccessedDetailsRequest) SetArn(v string) *GenerateServiceLastAccessedDetailsRequest {
	s.Arn = &v
	return s
}

func (s *GenerateServiceLastAccessedDetailsRequest) SetGranularity(v string) *GenerateServiceLastAccessedDetailsRequest {
	s.Granularity = &v
	return s
}

func (s *GenerateServiceLastAccessedDetailsRequest) Validate() error {
	return dara.Validate(s)
}
