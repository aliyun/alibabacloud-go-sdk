// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDSEntityValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntityValueId(v int64) *CreateDSEntityValueResponseBody
	GetEntityValueId() *int64
	SetRequestId(v string) *CreateDSEntityValueResponseBody
	GetRequestId() *string
}

type CreateDSEntityValueResponseBody struct {
	// example:
	//
	// 2434543453
	EntityValueId *int64 `json:"EntityValueId,omitempty" xml:"EntityValueId,omitempty"`
	// example:
	//
	// g763hg48j3f3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDSEntityValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDSEntityValueResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDSEntityValueResponseBody) GetEntityValueId() *int64 {
	return s.EntityValueId
}

func (s *CreateDSEntityValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDSEntityValueResponseBody) SetEntityValueId(v int64) *CreateDSEntityValueResponseBody {
	s.EntityValueId = &v
	return s
}

func (s *CreateDSEntityValueResponseBody) SetRequestId(v string) *CreateDSEntityValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDSEntityValueResponseBody) Validate() error {
	return dara.Validate(s)
}
