// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceEndpointAclPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateInstanceEndpointAclPolicyRequest
	GetComment() *string
	SetEndpointType(v string) *CreateInstanceEndpointAclPolicyRequest
	GetEndpointType() *string
	SetEntries(v []*AccessControlEntry) *CreateInstanceEndpointAclPolicyRequest
	GetEntries() []*AccessControlEntry
	SetEntry(v string) *CreateInstanceEndpointAclPolicyRequest
	GetEntry() *string
	SetInstanceId(v string) *CreateInstanceEndpointAclPolicyRequest
	GetInstanceId() *string
	SetModuleName(v string) *CreateInstanceEndpointAclPolicyRequest
	GetModuleName() *string
}

type CreateInstanceEndpointAclPolicyRequest struct {
	// Deprecated
	//
	// The description.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The endpoint type. Only Internet is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// internet
	EndpointType *string               `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	Entries      []*AccessControlEntry `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The IP address range that is allowed to access the instance.
	//
	// example:
	//
	// 192.168.1.1/32
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The module for which you want to set the access policy. Valid values:
	//
	// - `Registry`: access the image repository
	//
	// - `Chart`: access Helm Chart
	//
	// example:
	//
	// Registry
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s CreateInstanceEndpointAclPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceEndpointAclPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceEndpointAclPolicyRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateInstanceEndpointAclPolicyRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *CreateInstanceEndpointAclPolicyRequest) GetEntries() []*AccessControlEntry {
	return s.Entries
}

func (s *CreateInstanceEndpointAclPolicyRequest) GetEntry() *string {
	return s.Entry
}

func (s *CreateInstanceEndpointAclPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceEndpointAclPolicyRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *CreateInstanceEndpointAclPolicyRequest) SetComment(v string) *CreateInstanceEndpointAclPolicyRequest {
	s.Comment = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyRequest) SetEndpointType(v string) *CreateInstanceEndpointAclPolicyRequest {
	s.EndpointType = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyRequest) SetEntries(v []*AccessControlEntry) *CreateInstanceEndpointAclPolicyRequest {
	s.Entries = v
	return s
}

func (s *CreateInstanceEndpointAclPolicyRequest) SetEntry(v string) *CreateInstanceEndpointAclPolicyRequest {
	s.Entry = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyRequest) SetInstanceId(v string) *CreateInstanceEndpointAclPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyRequest) SetModuleName(v string) *CreateInstanceEndpointAclPolicyRequest {
	s.ModuleName = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyRequest) Validate() error {
	if s.Entries != nil {
		for _, item := range s.Entries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
