// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDSEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntityId(v int64) *UpdateDSEntityResponseBody
	GetEntityId() *int64
	SetRequestId(v string) *UpdateDSEntityResponseBody
	GetRequestId() *string
}

type UpdateDSEntityResponseBody struct {
	// example:
	//
	// 123
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// df23fgh4hyj67hn56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDSEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDSEntityResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDSEntityResponseBody) GetEntityId() *int64 {
	return s.EntityId
}

func (s *UpdateDSEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDSEntityResponseBody) SetEntityId(v int64) *UpdateDSEntityResponseBody {
	s.EntityId = &v
	return s
}

func (s *UpdateDSEntityResponseBody) SetRequestId(v string) *UpdateDSEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDSEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
