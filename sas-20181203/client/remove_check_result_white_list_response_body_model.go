// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveCheckResultWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveCheckResultWhiteListResponseBody
	GetCode() *string
	SetMessage(v string) *RemoveCheckResultWhiteListResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveCheckResultWhiteListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveCheckResultWhiteListResponseBody
	GetSuccess() *bool
}

type RemoveCheckResultWhiteListResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 98C82076-E0D5-51DA-99F2-513F4XXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveCheckResultWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveCheckResultWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveCheckResultWhiteListResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveCheckResultWhiteListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveCheckResultWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveCheckResultWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveCheckResultWhiteListResponseBody) SetCode(v string) *RemoveCheckResultWhiteListResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveCheckResultWhiteListResponseBody) SetMessage(v string) *RemoveCheckResultWhiteListResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveCheckResultWhiteListResponseBody) SetRequestId(v string) *RemoveCheckResultWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveCheckResultWhiteListResponseBody) SetSuccess(v bool) *RemoveCheckResultWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveCheckResultWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
