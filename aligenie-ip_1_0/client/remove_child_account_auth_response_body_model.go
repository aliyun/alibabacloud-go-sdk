// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveChildAccountAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *RemoveChildAccountAuthResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveChildAccountAuthResponseBody
	GetRequestId() *string
	SetResult(v bool) *RemoveChildAccountAuthResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *RemoveChildAccountAuthResponseBody
	GetStatusCode() *int32
}

type RemoveChildAccountAuthResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F12B***F34E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s RemoveChildAccountAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveChildAccountAuthResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveChildAccountAuthResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveChildAccountAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveChildAccountAuthResponseBody) GetResult() *bool {
	return s.Result
}

func (s *RemoveChildAccountAuthResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveChildAccountAuthResponseBody) SetMessage(v string) *RemoveChildAccountAuthResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveChildAccountAuthResponseBody) SetRequestId(v string) *RemoveChildAccountAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveChildAccountAuthResponseBody) SetResult(v bool) *RemoveChildAccountAuthResponseBody {
	s.Result = &v
	return s
}

func (s *RemoveChildAccountAuthResponseBody) SetStatusCode(v int32) *RemoveChildAccountAuthResponseBody {
	s.StatusCode = &v
	return s
}

func (s *RemoveChildAccountAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
