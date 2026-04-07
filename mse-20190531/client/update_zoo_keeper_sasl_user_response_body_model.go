// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateZooKeeperSaslUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateZooKeeperSaslUserResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateZooKeeperSaslUserResponseBody
	GetRequestId() *string
}

type UpdateZooKeeperSaslUserResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 316F5F64-F73D-42DC-8632-01E308B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateZooKeeperSaslUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateZooKeeperSaslUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateZooKeeperSaslUserResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateZooKeeperSaslUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateZooKeeperSaslUserResponseBody) SetData(v bool) *UpdateZooKeeperSaslUserResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateZooKeeperSaslUserResponseBody) SetRequestId(v string) *UpdateZooKeeperSaslUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateZooKeeperSaslUserResponseBody) Validate() error {
	return dara.Validate(s)
}
