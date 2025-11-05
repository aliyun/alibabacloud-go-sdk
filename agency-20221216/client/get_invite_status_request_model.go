// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInviteStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInviteStatusList(v []*GetInviteStatusRequestInviteStatusList) *GetInviteStatusRequest
	GetInviteStatusList() []*GetInviteStatusRequestInviteStatusList
}

type GetInviteStatusRequest struct {
	// inviteId list</br>
	//
	// `Sub-levels <= 5`
	//
	// This parameter is required.
	InviteStatusList []*GetInviteStatusRequestInviteStatusList `json:"InviteStatusList,omitempty" xml:"InviteStatusList,omitempty" type:"Repeated"`
}

func (s GetInviteStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInviteStatusRequest) GoString() string {
	return s.String()
}

func (s *GetInviteStatusRequest) GetInviteStatusList() []*GetInviteStatusRequestInviteStatusList {
	return s.InviteStatusList
}

func (s *GetInviteStatusRequest) SetInviteStatusList(v []*GetInviteStatusRequestInviteStatusList) *GetInviteStatusRequest {
	s.InviteStatusList = v
	return s
}

func (s *GetInviteStatusRequest) Validate() error {
	if s.InviteStatusList != nil {
		for _, item := range s.InviteStatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInviteStatusRequestInviteStatusList struct {
	// Invitation ID, From interface InviteSubAccount
	//
	// example:
	//
	// 123
	InviteId *int64 `json:"InviteId,omitempty" xml:"InviteId,omitempty"`
}

func (s GetInviteStatusRequestInviteStatusList) String() string {
	return dara.Prettify(s)
}

func (s GetInviteStatusRequestInviteStatusList) GoString() string {
	return s.String()
}

func (s *GetInviteStatusRequestInviteStatusList) GetInviteId() *int64 {
	return s.InviteId
}

func (s *GetInviteStatusRequestInviteStatusList) SetInviteId(v int64) *GetInviteStatusRequestInviteStatusList {
	s.InviteId = &v
	return s
}

func (s *GetInviteStatusRequestInviteStatusList) Validate() error {
	return dara.Validate(s)
}
