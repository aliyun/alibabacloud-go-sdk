// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveConfigDataSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveConfigDataSetResponseBody
	GetCode() *string
	SetMessage(v string) *SaveConfigDataSetResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveConfigDataSetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveConfigDataSetResponseBody
	GetSuccess() *bool
}

type SaveConfigDataSetResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveConfigDataSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveConfigDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *SaveConfigDataSetResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveConfigDataSetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveConfigDataSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveConfigDataSetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveConfigDataSetResponseBody) SetCode(v string) *SaveConfigDataSetResponseBody {
	s.Code = &v
	return s
}

func (s *SaveConfigDataSetResponseBody) SetMessage(v string) *SaveConfigDataSetResponseBody {
	s.Message = &v
	return s
}

func (s *SaveConfigDataSetResponseBody) SetRequestId(v string) *SaveConfigDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveConfigDataSetResponseBody) SetSuccess(v bool) *SaveConfigDataSetResponseBody {
	s.Success = &v
	return s
}

func (s *SaveConfigDataSetResponseBody) Validate() error {
	return dara.Validate(s)
}
