// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateNetwrokResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePrivateNetwrokResponseBody
	GetRequestId() *string
	SetResult(v []*UpdatePrivateNetwrokResponseBodyResult) *UpdatePrivateNetwrokResponseBody
	GetResult() []*UpdatePrivateNetwrokResponseBodyResult
}

type UpdatePrivateNetwrokResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2C5DAA30-****-5181-9B87-9D6181016197
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*UpdatePrivateNetwrokResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s UpdatePrivateNetwrokResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateNetwrokResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrivateNetwrokResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePrivateNetwrokResponseBody) GetResult() []*UpdatePrivateNetwrokResponseBodyResult {
	return s.Result
}

func (s *UpdatePrivateNetwrokResponseBody) SetRequestId(v string) *UpdatePrivateNetwrokResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrivateNetwrokResponseBody) SetResult(v []*UpdatePrivateNetwrokResponseBodyResult) *UpdatePrivateNetwrokResponseBody {
	s.Result = v
	return s
}

func (s *UpdatePrivateNetwrokResponseBody) Validate() error {
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

type UpdatePrivateNetwrokResponseBodyResult struct {
	// example:
	//
	// test-**.private.cn-hangzhou.es-serverless.aliyuncs.com
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
	// ep-bp1id6dc2568****
	PvlEndpointId *string `json:"pvlEndpointId,omitempty" xml:"pvlEndpointId,omitempty"`
	// example:
	//
	// PRIVATE_ES
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// vpc-uf6gykvwcirp886ef****
	VpcId        *string                                               `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	WhiteIpGroup []*UpdatePrivateNetwrokResponseBodyResultWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s UpdatePrivateNetwrokResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateNetwrokResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdatePrivateNetwrokResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *UpdatePrivateNetwrokResponseBodyResult) GetEnabled() *string {
	return s.Enabled
}

func (s *UpdatePrivateNetwrokResponseBodyResult) GetPort() *string {
	return s.Port
}

func (s *UpdatePrivateNetwrokResponseBodyResult) GetPvlEndpointId() *string {
	return s.PvlEndpointId
}

func (s *UpdatePrivateNetwrokResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *UpdatePrivateNetwrokResponseBodyResult) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdatePrivateNetwrokResponseBodyResult) GetWhiteIpGroup() []*UpdatePrivateNetwrokResponseBodyResultWhiteIpGroup {
	return s.WhiteIpGroup
}

func (s *UpdatePrivateNetwrokResponseBodyResult) SetDomain(v string) *UpdatePrivateNetwrokResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *UpdatePrivateNetwrokResponseBodyResult) SetEnabled(v string) *UpdatePrivateNetwrokResponseBodyResult {
	s.Enabled = &v
	return s
}

func (s *UpdatePrivateNetwrokResponseBodyResult) SetPort(v string) *UpdatePrivateNetwrokResponseBodyResult {
	s.Port = &v
	return s
}

func (s *UpdatePrivateNetwrokResponseBodyResult) SetPvlEndpointId(v string) *UpdatePrivateNetwrokResponseBodyResult {
	s.PvlEndpointId = &v
	return s
}

func (s *UpdatePrivateNetwrokResponseBodyResult) SetType(v string) *UpdatePrivateNetwrokResponseBodyResult {
	s.Type = &v
	return s
}

func (s *UpdatePrivateNetwrokResponseBodyResult) SetVpcId(v string) *UpdatePrivateNetwrokResponseBodyResult {
	s.VpcId = &v
	return s
}

func (s *UpdatePrivateNetwrokResponseBodyResult) SetWhiteIpGroup(v []*UpdatePrivateNetwrokResponseBodyResultWhiteIpGroup) *UpdatePrivateNetwrokResponseBodyResult {
	s.WhiteIpGroup = v
	return s
}

func (s *UpdatePrivateNetwrokResponseBodyResult) Validate() error {
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

type UpdatePrivateNetwrokResponseBodyResultWhiteIpGroup struct {
	// example:
	//
	// default
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s UpdatePrivateNetwrokResponseBodyResultWhiteIpGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateNetwrokResponseBodyResultWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *UpdatePrivateNetwrokResponseBodyResultWhiteIpGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdatePrivateNetwrokResponseBodyResultWhiteIpGroup) GetIps() []*string {
	return s.Ips
}

func (s *UpdatePrivateNetwrokResponseBodyResultWhiteIpGroup) SetGroupName(v string) *UpdatePrivateNetwrokResponseBodyResultWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *UpdatePrivateNetwrokResponseBodyResultWhiteIpGroup) SetIps(v []*string) *UpdatePrivateNetwrokResponseBodyResultWhiteIpGroup {
	s.Ips = v
	return s
}

func (s *UpdatePrivateNetwrokResponseBodyResultWhiteIpGroup) Validate() error {
	return dara.Validate(s)
}
