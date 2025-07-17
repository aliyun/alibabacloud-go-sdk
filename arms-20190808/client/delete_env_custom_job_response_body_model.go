// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnvCustomJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteEnvCustomJobResponseBody
	GetCode() *int32
	SetData(v string) *DeleteEnvCustomJobResponseBody
	GetData() *string
	SetMessage(v string) *DeleteEnvCustomJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEnvCustomJobResponseBody
	GetRequestId() *string
}

type DeleteEnvCustomJobResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the operation.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2FC13182-B9AF-4E6B-BE51-72669B7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEnvCustomJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnvCustomJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnvCustomJobResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteEnvCustomJobResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteEnvCustomJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEnvCustomJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEnvCustomJobResponseBody) SetCode(v int32) *DeleteEnvCustomJobResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEnvCustomJobResponseBody) SetData(v string) *DeleteEnvCustomJobResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteEnvCustomJobResponseBody) SetMessage(v string) *DeleteEnvCustomJobResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEnvCustomJobResponseBody) SetRequestId(v string) *DeleteEnvCustomJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEnvCustomJobResponseBody) Validate() error {
	return dara.Validate(s)
}
