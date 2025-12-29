// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChildAccountAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtentions(v map[string]interface{}) *ChildAccountAuthResponseBody
	GetExtentions() map[string]interface{}
	SetMessage(v string) *ChildAccountAuthResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChildAccountAuthResponseBody
	GetRequestId() *string
	SetResult(v bool) *ChildAccountAuthResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *ChildAccountAuthResponseBody
	GetStatusCode() *int32
}

type ChildAccountAuthResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3DB51A10-327C-58D3-91DF-3A5A471C51E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ChildAccountAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChildAccountAuthResponseBody) GoString() string {
	return s.String()
}

func (s *ChildAccountAuthResponseBody) GetExtentions() map[string]interface{} {
	return s.Extentions
}

func (s *ChildAccountAuthResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChildAccountAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChildAccountAuthResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ChildAccountAuthResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChildAccountAuthResponseBody) SetExtentions(v map[string]interface{}) *ChildAccountAuthResponseBody {
	s.Extentions = v
	return s
}

func (s *ChildAccountAuthResponseBody) SetMessage(v string) *ChildAccountAuthResponseBody {
	s.Message = &v
	return s
}

func (s *ChildAccountAuthResponseBody) SetRequestId(v string) *ChildAccountAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChildAccountAuthResponseBody) SetResult(v bool) *ChildAccountAuthResponseBody {
	s.Result = &v
	return s
}

func (s *ChildAccountAuthResponseBody) SetStatusCode(v int32) *ChildAccountAuthResponseBody {
	s.StatusCode = &v
	return s
}

func (s *ChildAccountAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
