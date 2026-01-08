// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateNetwrokRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*UpdatePrivateNetwrokRequestBody) *UpdatePrivateNetwrokRequest
	GetBody() []*UpdatePrivateNetwrokRequestBody
}

type UpdatePrivateNetwrokRequest struct {
	Body []*UpdatePrivateNetwrokRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s UpdatePrivateNetwrokRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateNetwrokRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivateNetwrokRequest) GetBody() []*UpdatePrivateNetwrokRequestBody {
	return s.Body
}

func (s *UpdatePrivateNetwrokRequest) SetBody(v []*UpdatePrivateNetwrokRequestBody) *UpdatePrivateNetwrokRequest {
	s.Body = v
	return s
}

func (s *UpdatePrivateNetwrokRequest) Validate() error {
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdatePrivateNetwrokRequestBody struct {
	// example:
	//
	// autotest-8k8a8.serverless.cn-hangzhou.elasticsearch.aliyuncs.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// 9200
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// ep-bp1id6dc2568****
	PvlEndpointId *string `json:"pvlEndpointId,omitempty" xml:"pvlEndpointId,omitempty"`
	// example:
	//
	// PRIVATE_ES
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// vpc-uf664nyle5khp5***
	VpcId        *string                                        `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	WhiteIpGroup []*UpdatePrivateNetwrokRequestBodyWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s UpdatePrivateNetwrokRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateNetwrokRequestBody) GoString() string {
	return s.String()
}

func (s *UpdatePrivateNetwrokRequestBody) GetDomain() *string {
	return s.Domain
}

func (s *UpdatePrivateNetwrokRequestBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdatePrivateNetwrokRequestBody) GetPort() *int32 {
	return s.Port
}

func (s *UpdatePrivateNetwrokRequestBody) GetPvlEndpointId() *string {
	return s.PvlEndpointId
}

func (s *UpdatePrivateNetwrokRequestBody) GetType() *string {
	return s.Type
}

func (s *UpdatePrivateNetwrokRequestBody) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdatePrivateNetwrokRequestBody) GetWhiteIpGroup() []*UpdatePrivateNetwrokRequestBodyWhiteIpGroup {
	return s.WhiteIpGroup
}

func (s *UpdatePrivateNetwrokRequestBody) SetDomain(v string) *UpdatePrivateNetwrokRequestBody {
	s.Domain = &v
	return s
}

func (s *UpdatePrivateNetwrokRequestBody) SetEnabled(v bool) *UpdatePrivateNetwrokRequestBody {
	s.Enabled = &v
	return s
}

func (s *UpdatePrivateNetwrokRequestBody) SetPort(v int32) *UpdatePrivateNetwrokRequestBody {
	s.Port = &v
	return s
}

func (s *UpdatePrivateNetwrokRequestBody) SetPvlEndpointId(v string) *UpdatePrivateNetwrokRequestBody {
	s.PvlEndpointId = &v
	return s
}

func (s *UpdatePrivateNetwrokRequestBody) SetType(v string) *UpdatePrivateNetwrokRequestBody {
	s.Type = &v
	return s
}

func (s *UpdatePrivateNetwrokRequestBody) SetVpcId(v string) *UpdatePrivateNetwrokRequestBody {
	s.VpcId = &v
	return s
}

func (s *UpdatePrivateNetwrokRequestBody) SetWhiteIpGroup(v []*UpdatePrivateNetwrokRequestBodyWhiteIpGroup) *UpdatePrivateNetwrokRequestBody {
	s.WhiteIpGroup = v
	return s
}

func (s *UpdatePrivateNetwrokRequestBody) Validate() error {
	if s.WhiteIpGroup != nil {
		for _, item := range s.WhiteIpGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdatePrivateNetwrokRequestBodyWhiteIpGroup struct {
	// example:
	//
	// default
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
	// example:
	//
	// Cover
	ModifyMode *string `json:"modifyMode,omitempty" xml:"modifyMode,omitempty"`
}

func (s UpdatePrivateNetwrokRequestBodyWhiteIpGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateNetwrokRequestBodyWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *UpdatePrivateNetwrokRequestBodyWhiteIpGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdatePrivateNetwrokRequestBodyWhiteIpGroup) GetIps() []*string {
	return s.Ips
}

func (s *UpdatePrivateNetwrokRequestBodyWhiteIpGroup) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *UpdatePrivateNetwrokRequestBodyWhiteIpGroup) SetGroupName(v string) *UpdatePrivateNetwrokRequestBodyWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *UpdatePrivateNetwrokRequestBodyWhiteIpGroup) SetIps(v []*string) *UpdatePrivateNetwrokRequestBodyWhiteIpGroup {
	s.Ips = v
	return s
}

func (s *UpdatePrivateNetwrokRequestBodyWhiteIpGroup) SetModifyMode(v string) *UpdatePrivateNetwrokRequestBodyWhiteIpGroup {
	s.ModifyMode = &v
	return s
}

func (s *UpdatePrivateNetwrokRequestBodyWhiteIpGroup) Validate() error {
	return dara.Validate(s)
}
