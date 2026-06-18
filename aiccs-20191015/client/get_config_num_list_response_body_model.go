// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigNumListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetConfigNumListResponseBody
	GetCode() *string
	SetData(v []*string) *GetConfigNumListResponseBody
	GetData() []*string
	SetMessage(v string) *GetConfigNumListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetConfigNumListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetConfigNumListResponseBody
	GetSuccess() *bool
}

type GetConfigNumListResponseBody struct {
	// The status code. A return value of "Success" indicates that the request succeeded.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number list.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Description of the status code.
	//
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE339D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetConfigNumListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConfigNumListResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigNumListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetConfigNumListResponseBody) GetData() []*string {
	return s.Data
}

func (s *GetConfigNumListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConfigNumListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConfigNumListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetConfigNumListResponseBody) SetCode(v string) *GetConfigNumListResponseBody {
	s.Code = &v
	return s
}

func (s *GetConfigNumListResponseBody) SetData(v []*string) *GetConfigNumListResponseBody {
	s.Data = v
	return s
}

func (s *GetConfigNumListResponseBody) SetMessage(v string) *GetConfigNumListResponseBody {
	s.Message = &v
	return s
}

func (s *GetConfigNumListResponseBody) SetRequestId(v string) *GetConfigNumListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConfigNumListResponseBody) SetSuccess(v bool) *GetConfigNumListResponseBody {
	s.Success = &v
	return s
}

func (s *GetConfigNumListResponseBody) Validate() error {
	return dara.Validate(s)
}
