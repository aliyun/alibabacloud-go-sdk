// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOriginPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateOriginPoolResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateOriginPoolResponseBody
	GetRequestId() *string
}

type CreateOriginPoolResponseBody struct {
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOriginPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOriginPoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOriginPoolResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateOriginPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOriginPoolResponseBody) SetId(v int64) *CreateOriginPoolResponseBody {
	s.Id = &v
	return s
}

func (s *CreateOriginPoolResponseBody) SetRequestId(v string) *CreateOriginPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOriginPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
