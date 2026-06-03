// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSchemeResponseBody
	GetCode() *string
	SetData(v string) *DeleteSchemeResponseBody
	GetData() *string
	SetMessage(v string) *DeleteSchemeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSchemeResponseBody
	GetRequestId() *string
}

type DeleteSchemeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSchemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSchemeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSchemeResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteSchemeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSchemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSchemeResponseBody) SetCode(v string) *DeleteSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSchemeResponseBody) SetData(v string) *DeleteSchemeResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteSchemeResponseBody) SetMessage(v string) *DeleteSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSchemeResponseBody) SetRequestId(v string) *DeleteSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSchemeResponseBody) Validate() error {
	return dara.Validate(s)
}
