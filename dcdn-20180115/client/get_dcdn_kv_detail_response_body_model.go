// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDcdnKvDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExpirationTtl(v string) *GetDcdnKvDetailResponseBody
	GetExpirationTtl() *string
	SetRequestId(v string) *GetDcdnKvDetailResponseBody
	GetRequestId() *string
	SetValue(v string) *GetDcdnKvDetailResponseBody
	GetValue() *string
}

type GetDcdnKvDetailResponseBody struct {
	// example:
	//
	// 3600
	ExpirationTtl *string `json:"ExpirationTtl,omitempty" xml:"ExpirationTtl,omitempty"`
	// example:
	//
	// EDBD3EB3-97DA-5465-AEF5-8DCA5DC5E395
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDcdnKvDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDcdnKvDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDcdnKvDetailResponseBody) GetExpirationTtl() *string {
	return s.ExpirationTtl
}

func (s *GetDcdnKvDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDcdnKvDetailResponseBody) GetValue() *string {
	return s.Value
}

func (s *GetDcdnKvDetailResponseBody) SetExpirationTtl(v string) *GetDcdnKvDetailResponseBody {
	s.ExpirationTtl = &v
	return s
}

func (s *GetDcdnKvDetailResponseBody) SetRequestId(v string) *GetDcdnKvDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDcdnKvDetailResponseBody) SetValue(v string) *GetDcdnKvDetailResponseBody {
	s.Value = &v
	return s
}

func (s *GetDcdnKvDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
