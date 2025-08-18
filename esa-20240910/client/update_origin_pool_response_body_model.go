// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOriginPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *UpdateOriginPoolResponseBody
	GetId() *int64
	SetRequestId(v string) *UpdateOriginPoolResponseBody
	GetRequestId() *string
}

type UpdateOriginPoolResponseBody struct {
	// Source address pool ID.
	//
	// example:
	//
	// 1038520525196928
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOriginPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOriginPoolResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOriginPoolResponseBody) GetId() *int64 {
	return s.Id
}

func (s *UpdateOriginPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOriginPoolResponseBody) SetId(v int64) *UpdateOriginPoolResponseBody {
	s.Id = &v
	return s
}

func (s *UpdateOriginPoolResponseBody) SetRequestId(v string) *UpdateOriginPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOriginPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
