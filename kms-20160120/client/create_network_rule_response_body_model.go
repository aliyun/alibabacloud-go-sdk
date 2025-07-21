// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *CreateNetworkRuleResponseBody
	GetArn() *string
	SetDescription(v string) *CreateNetworkRuleResponseBody
	GetDescription() *string
	SetName(v string) *CreateNetworkRuleResponseBody
	GetName() *string
	SetRequestId(v string) *CreateNetworkRuleResponseBody
	GetRequestId() *string
	SetSourcePrivateIp(v string) *CreateNetworkRuleResponseBody
	GetSourcePrivateIp() *string
	SetType(v string) *CreateNetworkRuleResponseBody
	GetType() *string
}

type CreateNetworkRuleResponseBody struct {
	// The ARN of the access control rule.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:119285303511****:network/networkrule_test
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The description.
	//
	// example:
	//
	// networkrule description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the access control rule.
	//
	// example:
	//
	// networkrule_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3bf02f7a-015b-4f93-be0f-cc043fda2dd3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The private IP address or private CIDR block.
	//
	// example:
	//
	// ["192.10.XX.XX","192.168.XX.XX/24"]
	SourcePrivateIp *string `json:"SourcePrivateIp,omitempty" xml:"SourcePrivateIp,omitempty"`
	// The network type.
	//
	// example:
	//
	// Private
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateNetworkRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkRuleResponseBody) GetArn() *string {
	return s.Arn
}

func (s *CreateNetworkRuleResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateNetworkRuleResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateNetworkRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkRuleResponseBody) GetSourcePrivateIp() *string {
	return s.SourcePrivateIp
}

func (s *CreateNetworkRuleResponseBody) GetType() *string {
	return s.Type
}

func (s *CreateNetworkRuleResponseBody) SetArn(v string) *CreateNetworkRuleResponseBody {
	s.Arn = &v
	return s
}

func (s *CreateNetworkRuleResponseBody) SetDescription(v string) *CreateNetworkRuleResponseBody {
	s.Description = &v
	return s
}

func (s *CreateNetworkRuleResponseBody) SetName(v string) *CreateNetworkRuleResponseBody {
	s.Name = &v
	return s
}

func (s *CreateNetworkRuleResponseBody) SetRequestId(v string) *CreateNetworkRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkRuleResponseBody) SetSourcePrivateIp(v string) *CreateNetworkRuleResponseBody {
	s.SourcePrivateIp = &v
	return s
}

func (s *CreateNetworkRuleResponseBody) SetType(v string) *CreateNetworkRuleResponseBody {
	s.Type = &v
	return s
}

func (s *CreateNetworkRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
