// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveZooKeeperSaslUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *RemoveZooKeeperSaslUserRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *RemoveZooKeeperSaslUserRequest
	GetInstanceId() *string
	SetReload(v bool) *RemoveZooKeeperSaslUserRequest
	GetReload() *bool
	SetUserNames(v []*string) *RemoveZooKeeperSaslUserRequest
	GetUserNames() []*string
}

type RemoveZooKeeperSaslUserRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mse_prepaid_public_cn-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	Reload *bool `json:"Reload,omitempty" xml:"Reload,omitempty"`
	// This parameter is required.
	UserNames []*string `json:"UserNames,omitempty" xml:"UserNames,omitempty" type:"Repeated"`
}

func (s RemoveZooKeeperSaslUserRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveZooKeeperSaslUserRequest) GoString() string {
	return s.String()
}

func (s *RemoveZooKeeperSaslUserRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *RemoveZooKeeperSaslUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveZooKeeperSaslUserRequest) GetReload() *bool {
	return s.Reload
}

func (s *RemoveZooKeeperSaslUserRequest) GetUserNames() []*string {
	return s.UserNames
}

func (s *RemoveZooKeeperSaslUserRequest) SetAcceptLanguage(v string) *RemoveZooKeeperSaslUserRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *RemoveZooKeeperSaslUserRequest) SetInstanceId(v string) *RemoveZooKeeperSaslUserRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveZooKeeperSaslUserRequest) SetReload(v bool) *RemoveZooKeeperSaslUserRequest {
	s.Reload = &v
	return s
}

func (s *RemoveZooKeeperSaslUserRequest) SetUserNames(v []*string) *RemoveZooKeeperSaslUserRequest {
	s.UserNames = v
	return s
}

func (s *RemoveZooKeeperSaslUserRequest) Validate() error {
	return dara.Validate(s)
}
