// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*UpdateNetworkRequestBody) *UpdateNetworkRequest
	GetBody() []*UpdateNetworkRequestBody
}

type UpdateNetworkRequest struct {
	Body []*UpdateNetworkRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s UpdateNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkRequest) GoString() string {
	return s.String()
}

func (s *UpdateNetworkRequest) GetBody() []*UpdateNetworkRequestBody {
	return s.Body
}

func (s *UpdateNetworkRequest) SetBody(v []*UpdateNetworkRequestBody) *UpdateNetworkRequest {
	s.Body = v
	return s
}

func (s *UpdateNetworkRequest) Validate() error {
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

type UpdateNetworkRequestBody struct {
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
	// PUBLIC_ES
	Type         *string                                 `json:"type,omitempty" xml:"type,omitempty"`
	WhiteIpGroup []*UpdateNetworkRequestBodyWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s UpdateNetworkRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateNetworkRequestBody) GetDomain() *string {
	return s.Domain
}

func (s *UpdateNetworkRequestBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateNetworkRequestBody) GetPort() *int32 {
	return s.Port
}

func (s *UpdateNetworkRequestBody) GetType() *string {
	return s.Type
}

func (s *UpdateNetworkRequestBody) GetWhiteIpGroup() []*UpdateNetworkRequestBodyWhiteIpGroup {
	return s.WhiteIpGroup
}

func (s *UpdateNetworkRequestBody) SetDomain(v string) *UpdateNetworkRequestBody {
	s.Domain = &v
	return s
}

func (s *UpdateNetworkRequestBody) SetEnabled(v bool) *UpdateNetworkRequestBody {
	s.Enabled = &v
	return s
}

func (s *UpdateNetworkRequestBody) SetPort(v int32) *UpdateNetworkRequestBody {
	s.Port = &v
	return s
}

func (s *UpdateNetworkRequestBody) SetType(v string) *UpdateNetworkRequestBody {
	s.Type = &v
	return s
}

func (s *UpdateNetworkRequestBody) SetWhiteIpGroup(v []*UpdateNetworkRequestBodyWhiteIpGroup) *UpdateNetworkRequestBody {
	s.WhiteIpGroup = v
	return s
}

func (s *UpdateNetworkRequestBody) Validate() error {
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

type UpdateNetworkRequestBodyWhiteIpGroup struct {
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

func (s UpdateNetworkRequestBodyWhiteIpGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkRequestBodyWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *UpdateNetworkRequestBodyWhiteIpGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateNetworkRequestBodyWhiteIpGroup) GetIps() []*string {
	return s.Ips
}

func (s *UpdateNetworkRequestBodyWhiteIpGroup) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *UpdateNetworkRequestBodyWhiteIpGroup) SetGroupName(v string) *UpdateNetworkRequestBodyWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *UpdateNetworkRequestBodyWhiteIpGroup) SetIps(v []*string) *UpdateNetworkRequestBodyWhiteIpGroup {
	s.Ips = v
	return s
}

func (s *UpdateNetworkRequestBodyWhiteIpGroup) SetModifyMode(v string) *UpdateNetworkRequestBodyWhiteIpGroup {
	s.ModifyMode = &v
	return s
}

func (s *UpdateNetworkRequestBodyWhiteIpGroup) Validate() error {
	return dara.Validate(s)
}
