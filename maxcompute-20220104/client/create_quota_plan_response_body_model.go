// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQuotaPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateQuotaPlanResponseBody
	GetData() *string
	SetRequestId(v string) *CreateQuotaPlanResponseBody
	GetRequestId() *string
}

type CreateQuotaPlanResponseBody struct {
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateQuotaPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQuotaPlanResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateQuotaPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQuotaPlanResponseBody) SetData(v string) *CreateQuotaPlanResponseBody {
	s.Data = &v
	return s
}

func (s *CreateQuotaPlanResponseBody) SetRequestId(v string) *CreateQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQuotaPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
