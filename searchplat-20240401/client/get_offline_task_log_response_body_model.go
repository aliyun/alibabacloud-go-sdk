// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOfflineTaskLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetOfflineTaskLogResponseBody
	GetRequestId() *string
	SetResult(v *GetOfflineTaskLogResponseBodyResult) *GetOfflineTaskLogResponseBody
	GetResult() *GetOfflineTaskLogResponseBodyResult
}

type GetOfflineTaskLogResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1-2-3-4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result *GetOfflineTaskLogResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetOfflineTaskLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOfflineTaskLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetOfflineTaskLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOfflineTaskLogResponseBody) GetResult() *GetOfflineTaskLogResponseBodyResult {
	return s.Result
}

func (s *GetOfflineTaskLogResponseBody) SetRequestId(v string) *GetOfflineTaskLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOfflineTaskLogResponseBody) SetResult(v *GetOfflineTaskLogResponseBodyResult) *GetOfflineTaskLogResponseBody {
	s.Result = v
	return s
}

func (s *GetOfflineTaskLogResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOfflineTaskLogResponseBodyResult struct {
	// The network information.
	Network *GetOfflineTaskLogResponseBodyResultNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
}

func (s GetOfflineTaskLogResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetOfflineTaskLogResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOfflineTaskLogResponseBodyResult) GetNetwork() *GetOfflineTaskLogResponseBodyResultNetwork {
	return s.Network
}

func (s *GetOfflineTaskLogResponseBodyResult) SetNetwork(v *GetOfflineTaskLogResponseBodyResultNetwork) *GetOfflineTaskLogResponseBodyResult {
	s.Network = v
	return s
}

func (s *GetOfflineTaskLogResponseBodyResult) Validate() error {
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOfflineTaskLogResponseBodyResultNetwork struct {
	// The private ES information.
	PrivateEs *GetOfflineTaskLogResponseBodyResultNetworkPrivateEs `json:"privateEs,omitempty" xml:"privateEs,omitempty" type:"Struct"`
	// The public ES information.
	PublicEs *GetOfflineTaskLogResponseBodyResultNetworkPublicEs `json:"publicEs,omitempty" xml:"publicEs,omitempty" type:"Struct"`
}

func (s GetOfflineTaskLogResponseBodyResultNetwork) String() string {
	return dara.Prettify(s)
}

func (s GetOfflineTaskLogResponseBodyResultNetwork) GoString() string {
	return s.String()
}

func (s *GetOfflineTaskLogResponseBodyResultNetwork) GetPrivateEs() *GetOfflineTaskLogResponseBodyResultNetworkPrivateEs {
	return s.PrivateEs
}

func (s *GetOfflineTaskLogResponseBodyResultNetwork) GetPublicEs() *GetOfflineTaskLogResponseBodyResultNetworkPublicEs {
	return s.PublicEs
}

func (s *GetOfflineTaskLogResponseBodyResultNetwork) SetPrivateEs(v *GetOfflineTaskLogResponseBodyResultNetworkPrivateEs) *GetOfflineTaskLogResponseBodyResultNetwork {
	s.PrivateEs = v
	return s
}

func (s *GetOfflineTaskLogResponseBodyResultNetwork) SetPublicEs(v *GetOfflineTaskLogResponseBodyResultNetworkPublicEs) *GetOfflineTaskLogResponseBodyResultNetwork {
	s.PublicEs = v
	return s
}

func (s *GetOfflineTaskLogResponseBodyResultNetwork) Validate() error {
	if s.PrivateEs != nil {
		if err := s.PrivateEs.Validate(); err != nil {
			return err
		}
	}
	if s.PublicEs != nil {
		if err := s.PublicEs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOfflineTaskLogResponseBodyResultNetworkPrivateEs struct {
	// The domain name of the private ES.
	//
	// example:
	//
	// test.private.cn-hangzhou.log.elasticsearch.aliyuncs.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// Indicates whether private ES is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The IP whitelist groups.
	WhiteIpGroup []*GetOfflineTaskLogResponseBodyResultNetworkPrivateEsWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s GetOfflineTaskLogResponseBodyResultNetworkPrivateEs) String() string {
	return dara.Prettify(s)
}

func (s GetOfflineTaskLogResponseBodyResultNetworkPrivateEs) GoString() string {
	return s.String()
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPrivateEs) GetDomain() *string {
	return s.Domain
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPrivateEs) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPrivateEs) GetWhiteIpGroup() []*GetOfflineTaskLogResponseBodyResultNetworkPrivateEsWhiteIpGroup {
	return s.WhiteIpGroup
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPrivateEs) SetDomain(v string) *GetOfflineTaskLogResponseBodyResultNetworkPrivateEs {
	s.Domain = &v
	return s
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPrivateEs) SetEnabled(v bool) *GetOfflineTaskLogResponseBodyResultNetworkPrivateEs {
	s.Enabled = &v
	return s
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPrivateEs) SetWhiteIpGroup(v []*GetOfflineTaskLogResponseBodyResultNetworkPrivateEsWhiteIpGroup) *GetOfflineTaskLogResponseBodyResultNetworkPrivateEs {
	s.WhiteIpGroup = v
	return s
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPrivateEs) Validate() error {
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

type GetOfflineTaskLogResponseBodyResultNetworkPrivateEsWhiteIpGroup struct {
	// The group name.
	//
	// example:
	//
	// kevintest
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The list of IP addresses in the whitelist group.
	Ips []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s GetOfflineTaskLogResponseBodyResultNetworkPrivateEsWhiteIpGroup) String() string {
	return dara.Prettify(s)
}

func (s GetOfflineTaskLogResponseBodyResultNetworkPrivateEsWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPrivateEsWhiteIpGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPrivateEsWhiteIpGroup) GetIps() []*string {
	return s.Ips
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPrivateEsWhiteIpGroup) SetGroupName(v string) *GetOfflineTaskLogResponseBodyResultNetworkPrivateEsWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPrivateEsWhiteIpGroup) SetIps(v []*string) *GetOfflineTaskLogResponseBodyResultNetworkPrivateEsWhiteIpGroup {
	s.Ips = v
	return s
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPrivateEsWhiteIpGroup) Validate() error {
	return dara.Validate(s)
}

type GetOfflineTaskLogResponseBodyResultNetworkPublicEs struct {
	// The public domain name of ES.
	//
	// example:
	//
	// test.public.cn-hangzhou.log.elasticsearch.aliyuncs.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// Indicates whether public ES is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The IP whitelist groups.
	WhiteIpGroup []*GetOfflineTaskLogResponseBodyResultNetworkPublicEsWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s GetOfflineTaskLogResponseBodyResultNetworkPublicEs) String() string {
	return dara.Prettify(s)
}

func (s GetOfflineTaskLogResponseBodyResultNetworkPublicEs) GoString() string {
	return s.String()
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPublicEs) GetDomain() *string {
	return s.Domain
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPublicEs) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPublicEs) GetWhiteIpGroup() []*GetOfflineTaskLogResponseBodyResultNetworkPublicEsWhiteIpGroup {
	return s.WhiteIpGroup
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPublicEs) SetDomain(v string) *GetOfflineTaskLogResponseBodyResultNetworkPublicEs {
	s.Domain = &v
	return s
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPublicEs) SetEnabled(v bool) *GetOfflineTaskLogResponseBodyResultNetworkPublicEs {
	s.Enabled = &v
	return s
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPublicEs) SetWhiteIpGroup(v []*GetOfflineTaskLogResponseBodyResultNetworkPublicEsWhiteIpGroup) *GetOfflineTaskLogResponseBodyResultNetworkPublicEs {
	s.WhiteIpGroup = v
	return s
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPublicEs) Validate() error {
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

type GetOfflineTaskLogResponseBodyResultNetworkPublicEsWhiteIpGroup struct {
	// The group name.
	//
	// example:
	//
	// kevintest
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The list of IP addresses in the whitelist group.
	Ips []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s GetOfflineTaskLogResponseBodyResultNetworkPublicEsWhiteIpGroup) String() string {
	return dara.Prettify(s)
}

func (s GetOfflineTaskLogResponseBodyResultNetworkPublicEsWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPublicEsWhiteIpGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPublicEsWhiteIpGroup) GetIps() []*string {
	return s.Ips
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPublicEsWhiteIpGroup) SetGroupName(v string) *GetOfflineTaskLogResponseBodyResultNetworkPublicEsWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPublicEsWhiteIpGroup) SetIps(v []*string) *GetOfflineTaskLogResponseBodyResultNetworkPublicEsWhiteIpGroup {
	s.Ips = v
	return s
}

func (s *GetOfflineTaskLogResponseBodyResultNetworkPublicEsWhiteIpGroup) Validate() error {
	return dara.Validate(s)
}
