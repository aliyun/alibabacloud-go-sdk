// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMmAppMemoryResponseBody
	GetCode() *string
	SetData(v *UpdateMmAppMemoryResponseBodyData) *UpdateMmAppMemoryResponseBody
	GetData() *UpdateMmAppMemoryResponseBodyData
	SetHttpStatusCode(v int32) *UpdateMmAppMemoryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateMmAppMemoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMmAppMemoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMmAppMemoryResponseBody
	GetSuccess() *bool
}

type UpdateMmAppMemoryResponseBody struct {
	// example:
	//
	// 200
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateMmAppMemoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82296D89-6895-574B-8AA1-64959957CB41
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMmAppMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMmAppMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMmAppMemoryResponseBody) GetData() *UpdateMmAppMemoryResponseBodyData {
	return s.Data
}

func (s *UpdateMmAppMemoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMmAppMemoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMmAppMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMmAppMemoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMmAppMemoryResponseBody) SetCode(v string) *UpdateMmAppMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMmAppMemoryResponseBody) SetData(v *UpdateMmAppMemoryResponseBodyData) *UpdateMmAppMemoryResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMmAppMemoryResponseBody) SetHttpStatusCode(v int32) *UpdateMmAppMemoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMmAppMemoryResponseBody) SetMessage(v string) *UpdateMmAppMemoryResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMmAppMemoryResponseBody) SetRequestId(v string) *UpdateMmAppMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMmAppMemoryResponseBody) SetSuccess(v bool) *UpdateMmAppMemoryResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMmAppMemoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMmAppMemoryResponseBodyData struct {
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMmAppMemoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppMemoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateMmAppMemoryResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMmAppMemoryResponseBodyData) SetSuccess(v bool) *UpdateMmAppMemoryResponseBodyData {
	s.Success = &v
	return s
}

func (s *UpdateMmAppMemoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
