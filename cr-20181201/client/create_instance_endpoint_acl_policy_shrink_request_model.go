// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceEndpointAclPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateInstanceEndpointAclPolicyShrinkRequest
	GetComment() *string
	SetEndpointType(v string) *CreateInstanceEndpointAclPolicyShrinkRequest
	GetEndpointType() *string
	SetEntriesShrink(v string) *CreateInstanceEndpointAclPolicyShrinkRequest
	GetEntriesShrink() *string
	SetEntry(v string) *CreateInstanceEndpointAclPolicyShrinkRequest
	GetEntry() *string
	SetInstanceId(v string) *CreateInstanceEndpointAclPolicyShrinkRequest
	GetInstanceId() *string
	SetModuleName(v string) *CreateInstanceEndpointAclPolicyShrinkRequest
	GetModuleName() *string
}

type CreateInstanceEndpointAclPolicyShrinkRequest struct {
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
	EndpointType  *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	EntriesShrink *string `json:"Entries,omitempty" xml:"Entries,omitempty"`
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

func (s CreateInstanceEndpointAclPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceEndpointAclPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceEndpointAclPolicyShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateInstanceEndpointAclPolicyShrinkRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *CreateInstanceEndpointAclPolicyShrinkRequest) GetEntriesShrink() *string {
	return s.EntriesShrink
}

func (s *CreateInstanceEndpointAclPolicyShrinkRequest) GetEntry() *string {
	return s.Entry
}

func (s *CreateInstanceEndpointAclPolicyShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceEndpointAclPolicyShrinkRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *CreateInstanceEndpointAclPolicyShrinkRequest) SetComment(v string) *CreateInstanceEndpointAclPolicyShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyShrinkRequest) SetEndpointType(v string) *CreateInstanceEndpointAclPolicyShrinkRequest {
	s.EndpointType = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyShrinkRequest) SetEntriesShrink(v string) *CreateInstanceEndpointAclPolicyShrinkRequest {
	s.EntriesShrink = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyShrinkRequest) SetEntry(v string) *CreateInstanceEndpointAclPolicyShrinkRequest {
	s.Entry = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyShrinkRequest) SetInstanceId(v string) *CreateInstanceEndpointAclPolicyShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyShrinkRequest) SetModuleName(v string) *CreateInstanceEndpointAclPolicyShrinkRequest {
	s.ModuleName = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
