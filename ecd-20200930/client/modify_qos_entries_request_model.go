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
	SetAuthDesktopGroupId(v []*string) *ModifyQosEntriesRequest
	GetAuthDesktopGroupId() []*string
	SetAuthDesktopId(v []*string) *ModifyQosEntriesRequest
	GetAuthDesktopId() []*string
	SetQosRuleId(v string) *ModifyQosEntriesRequest
	GetQosRuleId() *string
	SetRevokeAndroidId(v []*string) *ModifyQosEntriesRequest
	GetRevokeAndroidId() []*string
	SetRevokeDesktopGroupId(v []*string) *ModifyQosEntriesRequest
	GetRevokeDesktopGroupId() []*string
	SetRevokeDesktopId(v []*string) *ModifyQosEntriesRequest
	GetRevokeDesktopId() []*string
}

type ModifyQosEntriesRequest struct {
	// The list of cloud phone IDs to associate.
	AuthAndroidId []*string `json:"AuthAndroidId,omitempty" xml:"AuthAndroidId,omitempty" type:"Repeated"`
	// The ID of the cloud desktop pool to authorize.
	AuthDesktopGroupId []*string `json:"AuthDesktopGroupId,omitempty" xml:"AuthDesktopGroupId,omitempty" type:"Repeated"`
	// The list of cloud desktop IDs to associate.
	AuthDesktopId []*string `json:"AuthDesktopId,omitempty" xml:"AuthDesktopId,omitempty" type:"Repeated"`
	// The ID of the public network rate limiting rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-5605u0gelk200****
	QosRuleId *string `json:"QosRuleId,omitempty" xml:"QosRuleId,omitempty"`
	// The list of cloud phone IDs to disassociate.
	RevokeAndroidId []*string `json:"RevokeAndroidId,omitempty" xml:"RevokeAndroidId,omitempty" type:"Repeated"`
	// The ID of the cloud desktop pool to revoke authorization from.
	RevokeDesktopGroupId []*string `json:"RevokeDesktopGroupId,omitempty" xml:"RevokeDesktopGroupId,omitempty" type:"Repeated"`
	// The list of cloud desktop IDs to disassociate.
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

func (s *ModifyQosEntriesRequest) GetAuthDesktopGroupId() []*string {
	return s.AuthDesktopGroupId
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

func (s *ModifyQosEntriesRequest) GetRevokeDesktopGroupId() []*string {
	return s.RevokeDesktopGroupId
}

func (s *ModifyQosEntriesRequest) GetRevokeDesktopId() []*string {
	return s.RevokeDesktopId
}

func (s *ModifyQosEntriesRequest) SetAuthAndroidId(v []*string) *ModifyQosEntriesRequest {
	s.AuthAndroidId = v
	return s
}

func (s *ModifyQosEntriesRequest) SetAuthDesktopGroupId(v []*string) *ModifyQosEntriesRequest {
	s.AuthDesktopGroupId = v
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

func (s *ModifyQosEntriesRequest) SetRevokeDesktopGroupId(v []*string) *ModifyQosEntriesRequest {
	s.RevokeDesktopGroupId = v
	return s
}

func (s *ModifyQosEntriesRequest) SetRevokeDesktopId(v []*string) *ModifyQosEntriesRequest {
	s.RevokeDesktopId = v
	return s
}

func (s *ModifyQosEntriesRequest) Validate() error {
	return dara.Validate(s)
}
