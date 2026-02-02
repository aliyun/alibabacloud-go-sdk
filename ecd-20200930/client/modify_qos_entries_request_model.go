// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQosEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthAndroidId(v []*string) *ModifyQosEntriesRequest
	GetAuthAndroidId() []*string
	SetAuthDesktopId(v []*string) *ModifyQosEntriesRequest
	GetAuthDesktopId() []*string
	SetQosRuleId(v string) *ModifyQosEntriesRequest
	GetQosRuleId() *string
	SetRevokeAndroidId(v []*string) *ModifyQosEntriesRequest
	GetRevokeAndroidId() []*string
	SetRevokeDesktopId(v []*string) *ModifyQosEntriesRequest
	GetRevokeDesktopId() []*string
}

type ModifyQosEntriesRequest struct {
	AuthAndroidId []*string `json:"AuthAndroidId,omitempty" xml:"AuthAndroidId,omitempty" type:"Repeated"`
	AuthDesktopId []*string `json:"AuthDesktopId,omitempty" xml:"AuthDesktopId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// qos-5605u0gelk200****
	QosRuleId       *string   `json:"QosRuleId,omitempty" xml:"QosRuleId,omitempty"`
	RevokeAndroidId []*string `json:"RevokeAndroidId,omitempty" xml:"RevokeAndroidId,omitempty" type:"Repeated"`
	RevokeDesktopId []*string `json:"RevokeDesktopId,omitempty" xml:"RevokeDesktopId,omitempty" type:"Repeated"`
}

func (s ModifyQosEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyQosEntriesRequest) GoString() string {
	return s.String()
}

func (s *ModifyQosEntriesRequest) GetAuthAndroidId() []*string {
	return s.AuthAndroidId
}

func (s *ModifyQosEntriesRequest) GetAuthDesktopId() []*string {
	return s.AuthDesktopId
}

func (s *ModifyQosEntriesRequest) GetQosRuleId() *string {
	return s.QosRuleId
}

func (s *ModifyQosEntriesRequest) GetRevokeAndroidId() []*string {
	return s.RevokeAndroidId
}

func (s *ModifyQosEntriesRequest) GetRevokeDesktopId() []*string {
	return s.RevokeDesktopId
}

func (s *ModifyQosEntriesRequest) SetAuthAndroidId(v []*string) *ModifyQosEntriesRequest {
	s.AuthAndroidId = v
	return s
}

func (s *ModifyQosEntriesRequest) SetAuthDesktopId(v []*string) *ModifyQosEntriesRequest {
	s.AuthDesktopId = v
	return s
}

func (s *ModifyQosEntriesRequest) SetQosRuleId(v string) *ModifyQosEntriesRequest {
	s.QosRuleId = &v
	return s
}

func (s *ModifyQosEntriesRequest) SetRevokeAndroidId(v []*string) *ModifyQosEntriesRequest {
	s.RevokeAndroidId = v
	return s
}

func (s *ModifyQosEntriesRequest) SetRevokeDesktopId(v []*string) *ModifyQosEntriesRequest {
	s.RevokeDesktopId = v
	return s
}

func (s *ModifyQosEntriesRequest) Validate() error {
	return dara.Validate(s)
}
