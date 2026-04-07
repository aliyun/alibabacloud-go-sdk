// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListZooKeeperSaslUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListZooKeeperSaslUserRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *ListZooKeeperSaslUserRequest
	GetInstanceId() *string
}

type ListZooKeeperSaslUserRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mse-cn-5bffa4e8630
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListZooKeeperSaslUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListZooKeeperSaslUserRequest) GoString() string {
	return s.String()
}

func (s *ListZooKeeperSaslUserRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListZooKeeperSaslUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListZooKeeperSaslUserRequest) SetAcceptLanguage(v string) *ListZooKeeperSaslUserRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListZooKeeperSaslUserRequest) SetInstanceId(v string) *ListZooKeeperSaslUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListZooKeeperSaslUserRequest) Validate() error {
	return dara.Validate(s)
}
