// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *RemoveApplicationResponseBody
	GetData() *string
	SetErrorCode(v string) *RemoveApplicationResponseBody
	GetErrorCode() *string
	SetMessage(v string) *RemoveApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveApplicationResponseBody
	GetSuccess() *bool
}

type RemoveApplicationResponseBody struct {
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// abcde-12345
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveApplicationResponseBody) GetData() *string {
	return s.Data
}

func (s *RemoveApplicationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RemoveApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveApplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveApplicationResponseBody) SetData(v string) *RemoveApplicationResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveApplicationResponseBody) SetErrorCode(v string) *RemoveApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RemoveApplicationResponseBody) SetMessage(v string) *RemoveApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveApplicationResponseBody) SetRequestId(v string) *RemoveApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveApplicationResponseBody) SetSuccess(v bool) *RemoveApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
