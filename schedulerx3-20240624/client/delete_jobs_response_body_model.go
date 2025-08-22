// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteJobsResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteJobsResponseBody
	GetSuccess() *bool
}

type DeleteJobsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 91143E1D-E235-5BE0-9364-C2EE28FFB5A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteJobsResponseBody) SetCode(v int32) *DeleteJobsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteJobsResponseBody) SetMessage(v string) *DeleteJobsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteJobsResponseBody) SetRequestId(v string) *DeleteJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteJobsResponseBody) SetSuccess(v bool) *DeleteJobsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
