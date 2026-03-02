// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePbcInvokeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPbcInvokeId(v int64) *CreatePbcInvokeResponseBody
	GetPbcInvokeId() *int64
	SetRequestId(v string) *CreatePbcInvokeResponseBody
	GetRequestId() *string
}

type CreatePbcInvokeResponseBody struct {
	PbcInvokeId *int64  `json:"pbcInvokeId,omitempty" xml:"pbcInvokeId,omitempty"`
	RequestId   *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePbcInvokeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePbcInvokeResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePbcInvokeResponseBody) GetPbcInvokeId() *int64 {
	return s.PbcInvokeId
}

func (s *CreatePbcInvokeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePbcInvokeResponseBody) SetPbcInvokeId(v int64) *CreatePbcInvokeResponseBody {
	s.PbcInvokeId = &v
	return s
}

func (s *CreatePbcInvokeResponseBody) SetRequestId(v string) *CreatePbcInvokeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePbcInvokeResponseBody) Validate() error {
	return dara.Validate(s)
}
