// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveZooKeeperSaslUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *RemoveZooKeeperSaslUserResponseBody
	GetData() *bool
	SetRequestId(v string) *RemoveZooKeeperSaslUserResponseBody
	GetRequestId() *string
}

type RemoveZooKeeperSaslUserResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 56D9E600-6348-4260-B35F-583413F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveZooKeeperSaslUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveZooKeeperSaslUserResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveZooKeeperSaslUserResponseBody) GetData() *bool {
	return s.Data
}

func (s *RemoveZooKeeperSaslUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveZooKeeperSaslUserResponseBody) SetData(v bool) *RemoveZooKeeperSaslUserResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveZooKeeperSaslUserResponseBody) SetRequestId(v string) *RemoveZooKeeperSaslUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveZooKeeperSaslUserResponseBody) Validate() error {
	return dara.Validate(s)
}
