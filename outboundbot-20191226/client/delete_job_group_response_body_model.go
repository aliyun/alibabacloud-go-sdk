// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteJobGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteJobGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteJobGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteJobGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteJobGroupResponseBody
	GetSuccess() *bool
}

type DeleteJobGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteJobGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteJobGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteJobGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteJobGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteJobGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteJobGroupResponseBody) SetCode(v string) *DeleteJobGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteJobGroupResponseBody) SetHttpStatusCode(v int32) *DeleteJobGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteJobGroupResponseBody) SetMessage(v string) *DeleteJobGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteJobGroupResponseBody) SetRequestId(v string) *DeleteJobGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteJobGroupResponseBody) SetSuccess(v bool) *DeleteJobGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteJobGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
