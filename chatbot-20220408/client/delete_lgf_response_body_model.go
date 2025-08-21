// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLgfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLgfId(v int64) *DeleteLgfResponseBody
	GetLgfId() *int64
	SetRequestId(v string) *DeleteLgfResponseBody
	GetRequestId() *string
}

type DeleteLgfResponseBody struct {
	// LGF ID
	//
	// example:
	//
	// 2342424
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
	// example:
	//
	// dgw2342424qw42
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLgfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLgfResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLgfResponseBody) GetLgfId() *int64 {
	return s.LgfId
}

func (s *DeleteLgfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLgfResponseBody) SetLgfId(v int64) *DeleteLgfResponseBody {
	s.LgfId = &v
	return s
}

func (s *DeleteLgfResponseBody) SetRequestId(v string) *DeleteLgfResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLgfResponseBody) Validate() error {
	return dara.Validate(s)
}
