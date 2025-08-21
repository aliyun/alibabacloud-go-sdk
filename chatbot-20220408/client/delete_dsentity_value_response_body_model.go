// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDSEntityValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntityValueId(v int64) *DeleteDSEntityValueResponseBody
	GetEntityValueId() *int64
	SetRequestId(v string) *DeleteDSEntityValueResponseBody
	GetRequestId() *string
}

type DeleteDSEntityValueResponseBody struct {
	// example:
	//
	// 3453453452
	EntityValueId *int64 `json:"EntityValueId,omitempty" xml:"EntityValueId,omitempty"`
	// example:
	//
	// dfdf2t3rfvb45y
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDSEntityValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDSEntityValueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDSEntityValueResponseBody) GetEntityValueId() *int64 {
	return s.EntityValueId
}

func (s *DeleteDSEntityValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDSEntityValueResponseBody) SetEntityValueId(v int64) *DeleteDSEntityValueResponseBody {
	s.EntityValueId = &v
	return s
}

func (s *DeleteDSEntityValueResponseBody) SetRequestId(v string) *DeleteDSEntityValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDSEntityValueResponseBody) Validate() error {
	return dara.Validate(s)
}
