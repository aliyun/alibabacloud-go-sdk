// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEnableAllFwSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PutEnableAllFwSwitchRequest
	GetInstanceId() *string
	SetLang(v string) *PutEnableAllFwSwitchRequest
	GetLang() *string
	SetSourceIp(v string) *PutEnableAllFwSwitchRequest
	GetSourceIp() *string
}

type PutEnableAllFwSwitchRequest struct {
	// The instance ID of your Cloud Firewall.
	//
	// example:
	//
	// i-2ze8v2x5kd9qyvp2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	//
	// The source IP address of the request.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s PutEnableAllFwSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s PutEnableAllFwSwitchRequest) GoString() string {
	return s.String()
}

func (s *PutEnableAllFwSwitchRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PutEnableAllFwSwitchRequest) GetLang() *string {
	return s.Lang
}

func (s *PutEnableAllFwSwitchRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *PutEnableAllFwSwitchRequest) SetInstanceId(v string) *PutEnableAllFwSwitchRequest {
	s.InstanceId = &v
	return s
}

func (s *PutEnableAllFwSwitchRequest) SetLang(v string) *PutEnableAllFwSwitchRequest {
	s.Lang = &v
	return s
}

func (s *PutEnableAllFwSwitchRequest) SetSourceIp(v string) *PutEnableAllFwSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *PutEnableAllFwSwitchRequest) Validate() error {
	return dara.Validate(s)
}
