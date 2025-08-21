// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLgfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLgfId(v int64) *CreateLgfResponseBody
	GetLgfId() *int64
	SetRequestId(v string) *CreateLgfResponseBody
	GetRequestId() *string
}

type CreateLgfResponseBody struct {
	// LGF ID
	//
	// example:
	//
	// 123453433453
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
	// example:
	//
	// af5fg3sdf457j5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLgfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLgfResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLgfResponseBody) GetLgfId() *int64 {
	return s.LgfId
}

func (s *CreateLgfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLgfResponseBody) SetLgfId(v int64) *CreateLgfResponseBody {
	s.LgfId = &v
	return s
}

func (s *CreateLgfResponseBody) SetRequestId(v string) *CreateLgfResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLgfResponseBody) Validate() error {
	return dara.Validate(s)
}
