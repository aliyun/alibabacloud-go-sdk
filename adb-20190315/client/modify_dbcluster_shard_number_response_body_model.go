// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterShardNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterShardNumberResponseBody
	GetRequestId() *string
}

type ModifyDBClusterShardNumberResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 25B56BC7-4978-40B3-9E48-4B7067******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterShardNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterShardNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterShardNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterShardNumberResponseBody) SetRequestId(v string) *ModifyDBClusterShardNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterShardNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
