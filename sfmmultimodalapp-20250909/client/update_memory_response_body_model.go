// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMemoryResponseBody
	GetCode() *string
	SetData(v *UpdateMemoryResponseBodyData) *UpdateMemoryResponseBody
	GetData() *UpdateMemoryResponseBodyData
	SetHttpStatusCode(v int32) *UpdateMemoryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateMemoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMemoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMemoryResponseBody
	GetSuccess() *bool
}

type UpdateMemoryResponseBody struct {
	// example:
	//
	// 200
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateMemoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 37C73FCC-E9EB-57D3-9413-651F5FCAE1D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMemoryResponseBody) GetData() *UpdateMemoryResponseBodyData {
	return s.Data
}

func (s *UpdateMemoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMemoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMemoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMemoryResponseBody) SetCode(v string) *UpdateMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMemoryResponseBody) SetData(v *UpdateMemoryResponseBodyData) *UpdateMemoryResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMemoryResponseBody) SetHttpStatusCode(v int32) *UpdateMemoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMemoryResponseBody) SetMessage(v string) *UpdateMemoryResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMemoryResponseBody) SetRequestId(v string) *UpdateMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMemoryResponseBody) SetSuccess(v bool) *UpdateMemoryResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMemoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMemoryResponseBodyData struct {
	// example:
	//
	// 37C73FCC-E9EB-57D3-9413-651F5FCAE1D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMemoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateMemoryResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMemoryResponseBodyData) SetRequestId(v string) *UpdateMemoryResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *UpdateMemoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
