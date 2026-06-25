// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkBlacklistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkBlacklistModel(v *GetNetworkBlacklistResponseBodyNetworkBlacklistModel) *GetNetworkBlacklistResponseBody
	GetNetworkBlacklistModel() *GetNetworkBlacklistResponseBodyNetworkBlacklistModel
	SetRequestId(v string) *GetNetworkBlacklistResponseBody
	GetRequestId() *string
}

type GetNetworkBlacklistResponseBody struct {
	// The network blacklist.
	NetworkBlacklistModel *GetNetworkBlacklistResponseBodyNetworkBlacklistModel `json:"NetworkBlacklistModel,omitempty" xml:"NetworkBlacklistModel,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A578AD3A-8E7C-54FE-A09F-B060941*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNetworkBlacklistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *GetNetworkBlacklistResponseBody) GetNetworkBlacklistModel() *GetNetworkBlacklistResponseBodyNetworkBlacklistModel {
	return s.NetworkBlacklistModel
}

func (s *GetNetworkBlacklistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNetworkBlacklistResponseBody) SetNetworkBlacklistModel(v *GetNetworkBlacklistResponseBodyNetworkBlacklistModel) *GetNetworkBlacklistResponseBody {
	s.NetworkBlacklistModel = v
	return s
}

func (s *GetNetworkBlacklistResponseBody) SetRequestId(v string) *GetNetworkBlacklistResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNetworkBlacklistResponseBody) Validate() error {
	if s.NetworkBlacklistModel != nil {
		if err := s.NetworkBlacklistModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNetworkBlacklistResponseBodyNetworkBlacklistModel struct {
	// The list of blacklisted domain names.
	DomainBlacklist []*string `json:"DomainBlacklist,omitempty" xml:"DomainBlacklist,omitempty" type:"Repeated"`
	// The list of blacklisted IP addresses.
	IpBlacklist []*string `json:"IpBlacklist,omitempty" xml:"IpBlacklist,omitempty" type:"Repeated"`
}

func (s GetNetworkBlacklistResponseBodyNetworkBlacklistModel) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkBlacklistResponseBodyNetworkBlacklistModel) GoString() string {
	return s.String()
}

func (s *GetNetworkBlacklistResponseBodyNetworkBlacklistModel) GetDomainBlacklist() []*string {
	return s.DomainBlacklist
}

func (s *GetNetworkBlacklistResponseBodyNetworkBlacklistModel) GetIpBlacklist() []*string {
	return s.IpBlacklist
}

func (s *GetNetworkBlacklistResponseBodyNetworkBlacklistModel) SetDomainBlacklist(v []*string) *GetNetworkBlacklistResponseBodyNetworkBlacklistModel {
	s.DomainBlacklist = v
	return s
}

func (s *GetNetworkBlacklistResponseBodyNetworkBlacklistModel) SetIpBlacklist(v []*string) *GetNetworkBlacklistResponseBodyNetworkBlacklistModel {
	s.IpBlacklist = v
	return s
}

func (s *GetNetworkBlacklistResponseBodyNetworkBlacklistModel) Validate() error {
	return dara.Validate(s)
}
