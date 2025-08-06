// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iControlVodAppServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommand(v string) *ControlVodAppServiceRequest
	GetCommand() *string
	SetOwnerId(v int64) *ControlVodAppServiceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ControlVodAppServiceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ControlVodAppServiceRequest
	GetResourceOwnerId() *int64
	SetResourceRealOwnerId(v int64) *ControlVodAppServiceRequest
	GetResourceRealOwnerId() *int64
}

type ControlVodAppServiceRequest struct {
	Command              *string `json:"Command,omitempty" xml:"Command,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
}

func (s ControlVodAppServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ControlVodAppServiceRequest) GoString() string {
	return s.String()
}

func (s *ControlVodAppServiceRequest) GetCommand() *string {
	return s.Command
}

func (s *ControlVodAppServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ControlVodAppServiceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ControlVodAppServiceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ControlVodAppServiceRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *ControlVodAppServiceRequest) SetCommand(v string) *ControlVodAppServiceRequest {
	s.Command = &v
	return s
}

func (s *ControlVodAppServiceRequest) SetOwnerId(v int64) *ControlVodAppServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *ControlVodAppServiceRequest) SetResourceOwnerAccount(v string) *ControlVodAppServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ControlVodAppServiceRequest) SetResourceOwnerId(v int64) *ControlVodAppServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ControlVodAppServiceRequest) SetResourceRealOwnerId(v int64) *ControlVodAppServiceRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *ControlVodAppServiceRequest) Validate() error {
	return dara.Validate(s)
}
