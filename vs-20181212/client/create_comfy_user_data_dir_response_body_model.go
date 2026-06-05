// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComfyUserDataDirResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateComfyUserDataDirResponseBody
	GetCode() *int64
	SetMessage(v string) *CreateComfyUserDataDirResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateComfyUserDataDirResponseBody
	GetRequestId() *string
}

type CreateComfyUserDataDirResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// conn failed!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateComfyUserDataDirResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateComfyUserDataDirResponseBody) GoString() string {
	return s.String()
}

func (s *CreateComfyUserDataDirResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateComfyUserDataDirResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateComfyUserDataDirResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateComfyUserDataDirResponseBody) SetCode(v int64) *CreateComfyUserDataDirResponseBody {
	s.Code = &v
	return s
}

func (s *CreateComfyUserDataDirResponseBody) SetMessage(v string) *CreateComfyUserDataDirResponseBody {
	s.Message = &v
	return s
}

func (s *CreateComfyUserDataDirResponseBody) SetRequestId(v string) *CreateComfyUserDataDirResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateComfyUserDataDirResponseBody) Validate() error {
	return dara.Validate(s)
}
