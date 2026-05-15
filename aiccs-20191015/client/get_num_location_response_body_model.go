// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNumLocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetNumLocationResponseBody
	GetCode() *string
	SetData(v string) *GetNumLocationResponseBody
	GetData() *string
	SetMessage(v string) *GetNumLocationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetNumLocationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNumLocationResponseBody
	GetSuccess() *bool
}

type GetNumLocationResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNumLocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNumLocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetNumLocationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetNumLocationResponseBody) GetData() *string {
	return s.Data
}

func (s *GetNumLocationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetNumLocationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNumLocationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNumLocationResponseBody) SetCode(v string) *GetNumLocationResponseBody {
	s.Code = &v
	return s
}

func (s *GetNumLocationResponseBody) SetData(v string) *GetNumLocationResponseBody {
	s.Data = &v
	return s
}

func (s *GetNumLocationResponseBody) SetMessage(v string) *GetNumLocationResponseBody {
	s.Message = &v
	return s
}

func (s *GetNumLocationResponseBody) SetRequestId(v string) *GetNumLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNumLocationResponseBody) SetSuccess(v bool) *GetNumLocationResponseBody {
	s.Success = &v
	return s
}

func (s *GetNumLocationResponseBody) Validate() error {
	return dara.Validate(s)
}
