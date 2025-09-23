// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReleaseApplicationResponseBody
	GetCode() *string
	SetData(v int64) *ReleaseApplicationResponseBody
	GetData() *int64
	SetMessage(v string) *ReleaseApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReleaseApplicationResponseBody
	GetRequestId() *string
}

type ReleaseApplicationResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return value.
	//
	// example:
	//
	// 1
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	//
	// example:
	//
	// The resource does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BFB7F5C8-FE7A-06CA-9F87-ABBF6B848F0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReleaseApplicationResponseBody) GetData() *int64 {
	return s.Data
}

func (s *ReleaseApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReleaseApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseApplicationResponseBody) SetCode(v string) *ReleaseApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseApplicationResponseBody) SetData(v int64) *ReleaseApplicationResponseBody {
	s.Data = &v
	return s
}

func (s *ReleaseApplicationResponseBody) SetMessage(v string) *ReleaseApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseApplicationResponseBody) SetRequestId(v string) *ReleaseApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
