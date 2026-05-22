// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKvDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExpirationTtl(v string) *GetKvDetailResponseBody
	GetExpirationTtl() *string
	SetRequestId(v string) *GetKvDetailResponseBody
	GetRequestId() *string
	SetValue(v string) *GetKvDetailResponseBody
	GetValue() *string
}

type GetKvDetailResponseBody struct {
	ExpirationTtl *string `json:"ExpirationTtl,omitempty" xml:"ExpirationTtl,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetKvDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKvDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetKvDetailResponseBody) GetExpirationTtl() *string {
	return s.ExpirationTtl
}

func (s *GetKvDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKvDetailResponseBody) GetValue() *string {
	return s.Value
}

func (s *GetKvDetailResponseBody) SetExpirationTtl(v string) *GetKvDetailResponseBody {
	s.ExpirationTtl = &v
	return s
}

func (s *GetKvDetailResponseBody) SetRequestId(v string) *GetKvDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKvDetailResponseBody) SetValue(v string) *GetKvDetailResponseBody {
	s.Value = &v
	return s
}

func (s *GetKvDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
