// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSwimmingLaneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteSwimmingLaneResponseBody
	GetCode() *int32
	SetData(v int32) *DeleteSwimmingLaneResponseBody
	GetData() *int32
	SetMessage(v string) *DeleteSwimmingLaneResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSwimmingLaneResponseBody
	GetRequestId() *string
}

type DeleteSwimmingLaneResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the lane is deleted.
	//
	// example:
	//
	// 1
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4264F69C-686C-4107-B493-0599C8d1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSwimmingLaneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSwimmingLaneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSwimmingLaneResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteSwimmingLaneResponseBody) GetData() *int32 {
	return s.Data
}

func (s *DeleteSwimmingLaneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSwimmingLaneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSwimmingLaneResponseBody) SetCode(v int32) *DeleteSwimmingLaneResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSwimmingLaneResponseBody) SetData(v int32) *DeleteSwimmingLaneResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteSwimmingLaneResponseBody) SetMessage(v string) *DeleteSwimmingLaneResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSwimmingLaneResponseBody) SetRequestId(v string) *DeleteSwimmingLaneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSwimmingLaneResponseBody) Validate() error {
	return dara.Validate(s)
}
