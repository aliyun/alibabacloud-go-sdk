// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateNetworkResponseBody
	GetRequestId() *string
	SetResult(v []*UpdateNetworkResponseBodyResult) *UpdateNetworkResponseBody
	GetResult() []*UpdateNetworkResponseBodyResult
}

type UpdateNetworkResponseBody struct {
	// example:
	//
	// 2C5DAA30-****-5181-9B87-9D6181016197
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*UpdateNetworkResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s UpdateNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNetworkResponseBody) GetResult() []*UpdateNetworkResponseBodyResult {
	return s.Result
}

func (s *UpdateNetworkResponseBody) SetRequestId(v string) *UpdateNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNetworkResponseBody) SetResult(v []*UpdateNetworkResponseBodyResult) *UpdateNetworkResponseBody {
	s.Result = v
	return s
}

func (s *UpdateNetworkResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateNetworkResponseBodyResult struct {
	// example:
	//
	// test-123.serverless.cn-hangzhou.elasticsearch.aliyuncs.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// example:
	//
	// true
	Enabled *string `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// 9200
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// PUBLIC_ES
	Type         *string                                        `json:"type,omitempty" xml:"type,omitempty"`
	WhiteIpGroup []*UpdateNetworkResponseBodyResultWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s UpdateNetworkResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateNetworkResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *UpdateNetworkResponseBodyResult) GetEnabled() *string {
	return s.Enabled
}

func (s *UpdateNetworkResponseBodyResult) GetPort() *string {
	return s.Port
}

func (s *UpdateNetworkResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *UpdateNetworkResponseBodyResult) GetWhiteIpGroup() []*UpdateNetworkResponseBodyResultWhiteIpGroup {
	return s.WhiteIpGroup
}

func (s *UpdateNetworkResponseBodyResult) SetDomain(v string) *UpdateNetworkResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *UpdateNetworkResponseBodyResult) SetEnabled(v string) *UpdateNetworkResponseBodyResult {
	s.Enabled = &v
	return s
}

func (s *UpdateNetworkResponseBodyResult) SetPort(v string) *UpdateNetworkResponseBodyResult {
	s.Port = &v
	return s
}

func (s *UpdateNetworkResponseBodyResult) SetType(v string) *UpdateNetworkResponseBodyResult {
	s.Type = &v
	return s
}

func (s *UpdateNetworkResponseBodyResult) SetWhiteIpGroup(v []*UpdateNetworkResponseBodyResultWhiteIpGroup) *UpdateNetworkResponseBodyResult {
	s.WhiteIpGroup = v
	return s
}

func (s *UpdateNetworkResponseBodyResult) Validate() error {
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

type UpdateNetworkResponseBodyResultWhiteIpGroup struct {
	// example:
	//
	// default
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s UpdateNetworkResponseBodyResultWhiteIpGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkResponseBodyResultWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *UpdateNetworkResponseBodyResultWhiteIpGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateNetworkResponseBodyResultWhiteIpGroup) GetIps() []*string {
	return s.Ips
}

func (s *UpdateNetworkResponseBodyResultWhiteIpGroup) SetGroupName(v string) *UpdateNetworkResponseBodyResultWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *UpdateNetworkResponseBodyResultWhiteIpGroup) SetIps(v []*string) *UpdateNetworkResponseBodyResultWhiteIpGroup {
	s.Ips = v
	return s
}

func (s *UpdateNetworkResponseBodyResultWhiteIpGroup) Validate() error {
	return dara.Validate(s)
}
