// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupOutput interface {
	dara.Model
	String() string
	GoString() string
	SetNewResourceGroupId(v string) *ChangeResourceGroupOutput
	GetNewResourceGroupId() *string
	SetOldResourceGroupId(v string) *ChangeResourceGroupOutput
	GetOldResourceGroupId() *string
	SetResourceId(v string) *ChangeResourceGroupOutput
	GetResourceId() *string
}

type ChangeResourceGroupOutput struct {
	NewResourceGroupId *string `json:"newResourceGroupId,omitempty" xml:"newResourceGroupId,omitempty"`
	OldResourceGroupId *string `json:"oldResourceGroupId,omitempty" xml:"oldResourceGroupId,omitempty"`
	ResourceId         *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
}

func (s ChangeResourceGroupOutput) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupOutput) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupOutput) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *ChangeResourceGroupOutput) GetOldResourceGroupId() *string {
	return s.OldResourceGroupId
}

func (s *ChangeResourceGroupOutput) GetResourceId() *string {
	return s.ResourceId
}

func (s *ChangeResourceGroupOutput) SetNewResourceGroupId(v string) *ChangeResourceGroupOutput {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupOutput) SetOldResourceGroupId(v string) *ChangeResourceGroupOutput {
	s.OldResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupOutput) SetResourceId(v string) *ChangeResourceGroupOutput {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupOutput) Validate() error {
	return dara.Validate(s)
}
