// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDSEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntityId(v int64) *DeleteDSEntityResponseBody
	GetEntityId() *int64
	SetRequestId(v string) *DeleteDSEntityResponseBody
	GetRequestId() *string
}

type DeleteDSEntityResponseBody struct {
	// example:
	//
	// 123
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// 4dfghf56235asdf452
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDSEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDSEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDSEntityResponseBody) GetEntityId() *int64 {
	return s.EntityId
}

func (s *DeleteDSEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDSEntityResponseBody) SetEntityId(v int64) *DeleteDSEntityResponseBody {
	s.EntityId = &v
	return s
}

func (s *DeleteDSEntityResponseBody) SetRequestId(v string) *DeleteDSEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDSEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
