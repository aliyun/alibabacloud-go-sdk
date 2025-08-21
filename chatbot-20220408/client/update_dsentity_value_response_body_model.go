// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDSEntityValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntityValueId(v int64) *UpdateDSEntityValueResponseBody
	GetEntityValueId() *int64
	SetRequestId(v string) *UpdateDSEntityValueResponseBody
	GetRequestId() *string
}

type UpdateDSEntityValueResponseBody struct {
	// example:
	//
	// 2342377423
	EntityValueId *int64 `json:"EntityValueId,omitempty" xml:"EntityValueId,omitempty"`
	// example:
	//
	// sDag3g43wesf2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDSEntityValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDSEntityValueResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDSEntityValueResponseBody) GetEntityValueId() *int64 {
	return s.EntityValueId
}

func (s *UpdateDSEntityValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDSEntityValueResponseBody) SetEntityValueId(v int64) *UpdateDSEntityValueResponseBody {
	s.EntityValueId = &v
	return s
}

func (s *UpdateDSEntityValueResponseBody) SetRequestId(v string) *UpdateDSEntityValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDSEntityValueResponseBody) Validate() error {
	return dara.Validate(s)
}
