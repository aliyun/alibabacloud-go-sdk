// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrailStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsOrganizationTrail(v bool) *GetTrailStatusRequest
	GetIsOrganizationTrail() *bool
	SetName(v string) *GetTrailStatusRequest
	GetName() *string
}

type GetTrailStatusRequest struct {
	// Specifies whether to query the status of a multi-account trail. Valid values:
	//
	// 	- true: Query the status of a multi-account trail.
	//
	// 	- false: Query the status of a single-account trail. It is the default value.
	//
	// example:
	//
	// false
	IsOrganizationTrail *bool `json:"IsOrganizationTrail,omitempty" xml:"IsOrganizationTrail,omitempty"`
	// The name of the trail.
	//
	// The name must be 6 to 36 characters in length. The name must start with a lowercase letter and can contain lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// > The name must be unique within your Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// trail-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetTrailStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTrailStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTrailStatusRequest) GetIsOrganizationTrail() *bool {
	return s.IsOrganizationTrail
}

func (s *GetTrailStatusRequest) GetName() *string {
	return s.Name
}

func (s *GetTrailStatusRequest) SetIsOrganizationTrail(v bool) *GetTrailStatusRequest {
	s.IsOrganizationTrail = &v
	return s
}

func (s *GetTrailStatusRequest) SetName(v string) *GetTrailStatusRequest {
	s.Name = &v
	return s
}

func (s *GetTrailStatusRequest) Validate() error {
	return dara.Validate(s)
}
