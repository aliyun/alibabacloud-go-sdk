// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMemoryResponseBody
	GetCode() *string
	SetData(v *DeleteMemoryResponseBodyData) *DeleteMemoryResponseBody
	GetData() *DeleteMemoryResponseBodyData
	SetHttpStatusCode(v int32) *DeleteMemoryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteMemoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMemoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMemoryResponseBody
	GetSuccess() *bool
}

type DeleteMemoryResponseBody struct {
	// example:
	//
	// SUCCEED
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteMemoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5979B783-0AF5-547E-A549-8659F8A2A12A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMemoryResponseBody) GetData() *DeleteMemoryResponseBodyData {
	return s.Data
}

func (s *DeleteMemoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteMemoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMemoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMemoryResponseBody) SetCode(v string) *DeleteMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMemoryResponseBody) SetData(v *DeleteMemoryResponseBodyData) *DeleteMemoryResponseBody {
	s.Data = v
	return s
}

func (s *DeleteMemoryResponseBody) SetHttpStatusCode(v int32) *DeleteMemoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteMemoryResponseBody) SetMessage(v string) *DeleteMemoryResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMemoryResponseBody) SetRequestId(v string) *DeleteMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMemoryResponseBody) SetSuccess(v bool) *DeleteMemoryResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMemoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMemoryResponseBodyData struct {
	// example:
	//
	// A90B1930-D2CC-57ED-A2F6-079466EB16F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMemoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteMemoryResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMemoryResponseBodyData) SetRequestId(v string) *DeleteMemoryResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *DeleteMemoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
