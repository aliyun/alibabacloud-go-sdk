// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopySchemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CopySchemeResponseBody
	GetCode() *string
	SetData(v string) *CopySchemeResponseBody
	GetData() *string
	SetMessage(v string) *CopySchemeResponseBody
	GetMessage() *string
	SetRequestId(v string) *CopySchemeResponseBody
	GetRequestId() *string
}

type CopySchemeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopySchemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopySchemeResponseBody) GoString() string {
	return s.String()
}

func (s *CopySchemeResponseBody) GetCode() *string {
	return s.Code
}

func (s *CopySchemeResponseBody) GetData() *string {
	return s.Data
}

func (s *CopySchemeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CopySchemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopySchemeResponseBody) SetCode(v string) *CopySchemeResponseBody {
	s.Code = &v
	return s
}

func (s *CopySchemeResponseBody) SetData(v string) *CopySchemeResponseBody {
	s.Data = &v
	return s
}

func (s *CopySchemeResponseBody) SetMessage(v string) *CopySchemeResponseBody {
	s.Message = &v
	return s
}

func (s *CopySchemeResponseBody) SetRequestId(v string) *CopySchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopySchemeResponseBody) Validate() error {
	return dara.Validate(s)
}
