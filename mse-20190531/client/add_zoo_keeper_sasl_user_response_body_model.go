// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddZooKeeperSaslUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *AddZooKeeperSaslUserResponseBody
	GetData() *bool
	SetRequestId(v string) *AddZooKeeperSaslUserResponseBody
	GetRequestId() *string
}

type AddZooKeeperSaslUserResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 5B170A0D-2C5D-4CF8-B808-03966B86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddZooKeeperSaslUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddZooKeeperSaslUserResponseBody) GoString() string {
	return s.String()
}

func (s *AddZooKeeperSaslUserResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddZooKeeperSaslUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddZooKeeperSaslUserResponseBody) SetData(v bool) *AddZooKeeperSaslUserResponseBody {
	s.Data = &v
	return s
}

func (s *AddZooKeeperSaslUserResponseBody) SetRequestId(v string) *AddZooKeeperSaslUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddZooKeeperSaslUserResponseBody) Validate() error {
	return dara.Validate(s)
}
