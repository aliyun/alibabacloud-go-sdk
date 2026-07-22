// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEnableFwSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAbnormalResourceStatusList(v []*PutEnableFwSwitchResponseBodyAbnormalResourceStatusList) *PutEnableFwSwitchResponseBody
	GetAbnormalResourceStatusList() []*PutEnableFwSwitchResponseBodyAbnormalResourceStatusList
	SetDryRun(v bool) *PutEnableFwSwitchResponseBody
	GetDryRun() *bool
	SetRequestId(v string) *PutEnableFwSwitchResponseBody
	GetRequestId() *string
}

type PutEnableFwSwitchResponseBody struct {
	// The status information list of assets that are not synchronized.
	AbnormalResourceStatusList []*PutEnableFwSwitchResponseBodyAbnormalResourceStatusList `json:"AbnormalResourceStatusList,omitempty" xml:"AbnormalResourceStatusList,omitempty" type:"Repeated"`
	// Indicates that this is a successful dry run response. A value of true indicates that only the dry run was completed and no real changes were made. This field is not returned or is set to false for real calls.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B2841452-CB8D-4F7D-B247-38E1CF7334F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutEnableFwSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutEnableFwSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *PutEnableFwSwitchResponseBody) GetAbnormalResourceStatusList() []*PutEnableFwSwitchResponseBodyAbnormalResourceStatusList {
	return s.AbnormalResourceStatusList
}

func (s *PutEnableFwSwitchResponseBody) GetDryRun() *bool {
	return s.DryRun
}

func (s *PutEnableFwSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutEnableFwSwitchResponseBody) SetAbnormalResourceStatusList(v []*PutEnableFwSwitchResponseBodyAbnormalResourceStatusList) *PutEnableFwSwitchResponseBody {
	s.AbnormalResourceStatusList = v
	return s
}

func (s *PutEnableFwSwitchResponseBody) SetDryRun(v bool) *PutEnableFwSwitchResponseBody {
	s.DryRun = &v
	return s
}

func (s *PutEnableFwSwitchResponseBody) SetRequestId(v string) *PutEnableFwSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutEnableFwSwitchResponseBody) Validate() error {
	if s.AbnormalResourceStatusList != nil {
		for _, item := range s.AbnormalResourceStatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutEnableFwSwitchResponseBodyAbnormalResourceStatusList struct {
	// The message when the asset is not synchronized. Valid values:
	//
	// - cloudfirewall do not sync this ip address: Cloud Firewall has not synchronized this asset IP address.
	//
	// example:
	//
	// cloudfirewall do not sync this ip address
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The asset IP address.
	//
	// example:
	//
	// 203.0.113.0
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The status of the asset that is not synchronized. Valid values:
	//
	// - ip_not_sync: The asset is not synchronized.
	//
	// example:
	//
	// ip_not_sync
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PutEnableFwSwitchResponseBodyAbnormalResourceStatusList) String() string {
	return dara.Prettify(s)
}

func (s PutEnableFwSwitchResponseBodyAbnormalResourceStatusList) GoString() string {
	return s.String()
}

func (s *PutEnableFwSwitchResponseBodyAbnormalResourceStatusList) GetMsg() *string {
	return s.Msg
}

func (s *PutEnableFwSwitchResponseBodyAbnormalResourceStatusList) GetResource() *string {
	return s.Resource
}

func (s *PutEnableFwSwitchResponseBodyAbnormalResourceStatusList) GetStatus() *string {
	return s.Status
}

func (s *PutEnableFwSwitchResponseBodyAbnormalResourceStatusList) SetMsg(v string) *PutEnableFwSwitchResponseBodyAbnormalResourceStatusList {
	s.Msg = &v
	return s
}

func (s *PutEnableFwSwitchResponseBodyAbnormalResourceStatusList) SetResource(v string) *PutEnableFwSwitchResponseBodyAbnormalResourceStatusList {
	s.Resource = &v
	return s
}

func (s *PutEnableFwSwitchResponseBodyAbnormalResourceStatusList) SetStatus(v string) *PutEnableFwSwitchResponseBodyAbnormalResourceStatusList {
	s.Status = &v
	return s
}

func (s *PutEnableFwSwitchResponseBodyAbnormalResourceStatusList) Validate() error {
	return dara.Validate(s)
}
