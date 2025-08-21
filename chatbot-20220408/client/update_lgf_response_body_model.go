// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLgfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLgfId(v int64) *UpdateLgfResponseBody
	GetLgfId() *int64
	SetRequestId(v string) *UpdateLgfResponseBody
	GetRequestId() *string
}

type UpdateLgfResponseBody struct {
	// example:
	//
	// 2342556223532
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
	// example:
	//
	// 289dfa131adf23wqe2r
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLgfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLgfResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLgfResponseBody) GetLgfId() *int64 {
	return s.LgfId
}

func (s *UpdateLgfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLgfResponseBody) SetLgfId(v int64) *UpdateLgfResponseBody {
	s.LgfId = &v
	return s
}

func (s *UpdateLgfResponseBody) SetRequestId(v string) *UpdateLgfResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLgfResponseBody) Validate() error {
	return dara.Validate(s)
}
