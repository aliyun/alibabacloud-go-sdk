// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceEndpointAclPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointType(v string) *DeleteInstanceEndpointAclPolicyRequest
	GetEndpointType() *string
	SetEntries(v []*AccessControlEntry) *DeleteInstanceEndpointAclPolicyRequest
	GetEntries() []*AccessControlEntry
	SetEntry(v string) *DeleteInstanceEndpointAclPolicyRequest
	GetEntry() *string
	SetInstanceId(v string) *DeleteInstanceEndpointAclPolicyRequest
	GetInstanceId() *string
	SetModuleName(v string) *DeleteInstanceEndpointAclPolicyRequest
	GetModuleName() *string
}

type DeleteInstanceEndpointAclPolicyRequest struct {
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
	// The IP CIDR block.
	//
	// example:
	//
	// 127.0.0.1/32
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The module for which the access policy is set. Valid values:
	//
	// - `Registry`: access to the image repository
	//
	// - `Chart`: access to Helm Chart
	//
	// example:
	//
	// Chart
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s DeleteInstanceEndpointAclPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceEndpointAclPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceEndpointAclPolicyRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *DeleteInstanceEndpointAclPolicyRequest) GetEntries() []*AccessControlEntry {
	return s.Entries
}

func (s *DeleteInstanceEndpointAclPolicyRequest) GetEntry() *string {
	return s.Entry
}

func (s *DeleteInstanceEndpointAclPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInstanceEndpointAclPolicyRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *DeleteInstanceEndpointAclPolicyRequest) SetEndpointType(v string) *DeleteInstanceEndpointAclPolicyRequest {
	s.EndpointType = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyRequest) SetEntries(v []*AccessControlEntry) *DeleteInstanceEndpointAclPolicyRequest {
	s.Entries = v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyRequest) SetEntry(v string) *DeleteInstanceEndpointAclPolicyRequest {
	s.Entry = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyRequest) SetInstanceId(v string) *DeleteInstanceEndpointAclPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyRequest) SetModuleName(v string) *DeleteInstanceEndpointAclPolicyRequest {
	s.ModuleName = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyRequest) Validate() error {
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
