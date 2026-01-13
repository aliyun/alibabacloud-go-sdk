// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecallManagementConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkConfigs(v []*GetRecallManagementConfigResponseBodyNetworkConfigs) *GetRecallManagementConfigResponseBody
	GetNetworkConfigs() []*GetRecallManagementConfigResponseBodyNetworkConfigs
	SetRequestId(v string) *GetRecallManagementConfigResponseBody
	GetRequestId() *string
	SetUserName(v string) *GetRecallManagementConfigResponseBody
	GetUserName() *string
}

type GetRecallManagementConfigResponseBody struct {
	NetworkConfigs []*GetRecallManagementConfigResponseBodyNetworkConfigs `json:"NetworkConfigs,omitempty" xml:"NetworkConfigs,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// scene_test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetRecallManagementConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecallManagementConfigResponseBody) GetNetworkConfigs() []*GetRecallManagementConfigResponseBodyNetworkConfigs {
	return s.NetworkConfigs
}

func (s *GetRecallManagementConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRecallManagementConfigResponseBody) GetUserName() *string {
	return s.UserName
}

func (s *GetRecallManagementConfigResponseBody) SetNetworkConfigs(v []*GetRecallManagementConfigResponseBodyNetworkConfigs) *GetRecallManagementConfigResponseBody {
	s.NetworkConfigs = v
	return s
}

func (s *GetRecallManagementConfigResponseBody) SetRequestId(v string) *GetRecallManagementConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecallManagementConfigResponseBody) SetUserName(v string) *GetRecallManagementConfigResponseBody {
	s.UserName = &v
	return s
}

func (s *GetRecallManagementConfigResponseBody) Validate() error {
	if s.NetworkConfigs != nil {
		for _, item := range s.NetworkConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRecallManagementConfigResponseBodyNetworkConfigs struct {
	// example:
	//
	// http://xxx
	PrivateLinkAddress *string `json:"PrivateLinkAddress,omitempty" xml:"PrivateLinkAddress,omitempty"`
	// example:
	//
	// Connecting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// vpc-xxx
	VpcId      *string            `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchIds map[string]*string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty"`
}

func (s GetRecallManagementConfigResponseBodyNetworkConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementConfigResponseBodyNetworkConfigs) GoString() string {
	return s.String()
}

func (s *GetRecallManagementConfigResponseBodyNetworkConfigs) GetPrivateLinkAddress() *string {
	return s.PrivateLinkAddress
}

func (s *GetRecallManagementConfigResponseBodyNetworkConfigs) GetStatus() *string {
	return s.Status
}

func (s *GetRecallManagementConfigResponseBodyNetworkConfigs) GetVpcId() *string {
	return s.VpcId
}

func (s *GetRecallManagementConfigResponseBodyNetworkConfigs) GetVswitchIds() map[string]*string {
	return s.VswitchIds
}

func (s *GetRecallManagementConfigResponseBodyNetworkConfigs) SetPrivateLinkAddress(v string) *GetRecallManagementConfigResponseBodyNetworkConfigs {
	s.PrivateLinkAddress = &v
	return s
}

func (s *GetRecallManagementConfigResponseBodyNetworkConfigs) SetStatus(v string) *GetRecallManagementConfigResponseBodyNetworkConfigs {
	s.Status = &v
	return s
}

func (s *GetRecallManagementConfigResponseBodyNetworkConfigs) SetVpcId(v string) *GetRecallManagementConfigResponseBodyNetworkConfigs {
	s.VpcId = &v
	return s
}

func (s *GetRecallManagementConfigResponseBodyNetworkConfigs) SetVswitchIds(v map[string]*string) *GetRecallManagementConfigResponseBodyNetworkConfigs {
	s.VswitchIds = v
	return s
}

func (s *GetRecallManagementConfigResponseBodyNetworkConfigs) Validate() error {
	return dara.Validate(s)
}
