// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDisableAllFwSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PutDisableAllFwSwitchRequest
	GetInstanceId() *string
	SetLang(v string) *PutDisableAllFwSwitchRequest
	GetLang() *string
	SetSourceIp(v string) *PutDisableAllFwSwitchRequest
	GetSourceIp() *string
}

type PutDisableAllFwSwitchRequest struct {
	// The instance ID of your Cloud Firewall.
	//
	// example:
	//
	// i-2ze8v2x5kd9qyvp2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the content within the request and response. Valid values: Valid values:
	//
	// 	- **zh*	- (default): Chinese
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

func (s PutDisableAllFwSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s PutDisableAllFwSwitchRequest) GoString() string {
	return s.String()
}

func (s *PutDisableAllFwSwitchRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PutDisableAllFwSwitchRequest) GetLang() *string {
	return s.Lang
}

func (s *PutDisableAllFwSwitchRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *PutDisableAllFwSwitchRequest) SetInstanceId(v string) *PutDisableAllFwSwitchRequest {
	s.InstanceId = &v
	return s
}

func (s *PutDisableAllFwSwitchRequest) SetLang(v string) *PutDisableAllFwSwitchRequest {
	s.Lang = &v
	return s
}

func (s *PutDisableAllFwSwitchRequest) SetSourceIp(v string) *PutDisableAllFwSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *PutDisableAllFwSwitchRequest) Validate() error {
	return dara.Validate(s)
}
