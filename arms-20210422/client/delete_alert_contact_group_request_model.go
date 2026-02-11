// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertContactGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupId(v int64) *DeleteAlertContactGroupRequest
	GetContactGroupId() *int64
	SetRegionId(v string) *DeleteAlertContactGroupRequest
	GetRegionId() *string
}

type DeleteAlertContactGroupRequest struct {
	// This parameter is required.
	ContactGroupId *int64 `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAlertContactGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertContactGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactGroupRequest) GetContactGroupId() *int64 {
	return s.ContactGroupId
}

func (s *DeleteAlertContactGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAlertContactGroupRequest) SetContactGroupId(v int64) *DeleteAlertContactGroupRequest {
	s.ContactGroupId = &v
	return s
}

func (s *DeleteAlertContactGroupRequest) SetRegionId(v string) *DeleteAlertContactGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAlertContactGroupRequest) Validate() error {
	return dara.Validate(s)
}
