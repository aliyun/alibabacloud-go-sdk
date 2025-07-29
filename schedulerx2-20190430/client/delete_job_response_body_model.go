// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteJobResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteJobResponseBody
	GetSuccess() *bool
}

type DeleteJobResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information returned.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the job was deleted. Valid values:
	//
	// 	- **true**: The job was deleted.
	//
	// 	- **false**: The job was not deleted.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteJobResponseBody) SetCode(v int32) *DeleteJobResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteJobResponseBody) SetMessage(v string) *DeleteJobResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteJobResponseBody) SetRequestId(v string) *DeleteJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteJobResponseBody) SetSuccess(v bool) *DeleteJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteJobResponseBody) Validate() error {
	return dara.Validate(s)
}
