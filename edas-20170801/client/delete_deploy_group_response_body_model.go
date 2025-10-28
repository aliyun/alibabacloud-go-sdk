// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeployGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteDeployGroupResponseBody
	GetCode() *int32
	SetData(v string) *DeleteDeployGroupResponseBody
	GetData() *string
	SetMessage(v string) *DeleteDeployGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDeployGroupResponseBody
	GetRequestId() *string
}

type DeleteDeployGroupResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	//
	// example:
	//
	// 1
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4D9F-DR94-FD****************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDeployGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeployGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeployGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteDeployGroupResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteDeployGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDeployGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDeployGroupResponseBody) SetCode(v int32) *DeleteDeployGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDeployGroupResponseBody) SetData(v string) *DeleteDeployGroupResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDeployGroupResponseBody) SetMessage(v string) *DeleteDeployGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDeployGroupResponseBody) SetRequestId(v string) *DeleteDeployGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeployGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
