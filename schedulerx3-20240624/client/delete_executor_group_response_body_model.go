// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExecutorGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteExecutorGroupResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteExecutorGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteExecutorGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteExecutorGroupResponseBody
	GetSuccess() *bool
}

type DeleteExecutorGroupResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3835AA29-2298-5434-BC53-9CC377CDFD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteExecutorGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExecutorGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExecutorGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteExecutorGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteExecutorGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExecutorGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteExecutorGroupResponseBody) SetCode(v int32) *DeleteExecutorGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteExecutorGroupResponseBody) SetMessage(v string) *DeleteExecutorGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteExecutorGroupResponseBody) SetRequestId(v string) *DeleteExecutorGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExecutorGroupResponseBody) SetSuccess(v bool) *DeleteExecutorGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteExecutorGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
