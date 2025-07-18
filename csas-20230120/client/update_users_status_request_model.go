// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUsersStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSaseUserIds(v []*string) *UpdateUsersStatusRequest
	GetSaseUserIds() []*string
	SetStatus(v string) *UpdateUsersStatusRequest
	GetStatus() *string
}

type UpdateUsersStatusRequest struct {
	// This parameter is required.
	SaseUserIds []*string `json:"SaseUserIds,omitempty" xml:"SaseUserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateUsersStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUsersStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateUsersStatusRequest) GetSaseUserIds() []*string {
	return s.SaseUserIds
}

func (s *UpdateUsersStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateUsersStatusRequest) SetSaseUserIds(v []*string) *UpdateUsersStatusRequest {
	s.SaseUserIds = v
	return s
}

func (s *UpdateUsersStatusRequest) SetStatus(v string) *UpdateUsersStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateUsersStatusRequest) Validate() error {
	return dara.Validate(s)
}
