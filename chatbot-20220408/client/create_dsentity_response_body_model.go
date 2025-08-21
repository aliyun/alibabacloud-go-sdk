// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDSEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntityId(v int64) *CreateDSEntityResponseBody
	GetEntityId() *int64
	SetRequestId(v string) *CreateDSEntityResponseBody
	GetRequestId() *string
}

type CreateDSEntityResponseBody struct {
	// example:
	//
	// 23436345
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// adfad2343f1f2r
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDSEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDSEntityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDSEntityResponseBody) GetEntityId() *int64 {
	return s.EntityId
}

func (s *CreateDSEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDSEntityResponseBody) SetEntityId(v int64) *CreateDSEntityResponseBody {
	s.EntityId = &v
	return s
}

func (s *CreateDSEntityResponseBody) SetRequestId(v string) *CreateDSEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDSEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
