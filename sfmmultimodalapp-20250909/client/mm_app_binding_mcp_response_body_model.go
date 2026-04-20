// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMmAppBindingMcpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MmAppBindingMcpResponseBody
	GetCode() *string
	SetData(v *MmAppBindingMcpResponseBodyData) *MmAppBindingMcpResponseBody
	GetData() *MmAppBindingMcpResponseBodyData
	SetHttpStatusCode(v int32) *MmAppBindingMcpResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *MmAppBindingMcpResponseBody
	GetMessage() *string
	SetRequestId(v string) *MmAppBindingMcpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MmAppBindingMcpResponseBody
	GetSuccess() *bool
}

type MmAppBindingMcpResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *MmAppBindingMcpResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 539759F7-A281-577D-9962-6E69AEBD3AB9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MmAppBindingMcpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MmAppBindingMcpResponseBody) GoString() string {
	return s.String()
}

func (s *MmAppBindingMcpResponseBody) GetCode() *string {
	return s.Code
}

func (s *MmAppBindingMcpResponseBody) GetData() *MmAppBindingMcpResponseBodyData {
	return s.Data
}

func (s *MmAppBindingMcpResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *MmAppBindingMcpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MmAppBindingMcpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MmAppBindingMcpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MmAppBindingMcpResponseBody) SetCode(v string) *MmAppBindingMcpResponseBody {
	s.Code = &v
	return s
}

func (s *MmAppBindingMcpResponseBody) SetData(v *MmAppBindingMcpResponseBodyData) *MmAppBindingMcpResponseBody {
	s.Data = v
	return s
}

func (s *MmAppBindingMcpResponseBody) SetHttpStatusCode(v int32) *MmAppBindingMcpResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *MmAppBindingMcpResponseBody) SetMessage(v string) *MmAppBindingMcpResponseBody {
	s.Message = &v
	return s
}

func (s *MmAppBindingMcpResponseBody) SetRequestId(v string) *MmAppBindingMcpResponseBody {
	s.RequestId = &v
	return s
}

func (s *MmAppBindingMcpResponseBody) SetSuccess(v bool) *MmAppBindingMcpResponseBody {
	s.Success = &v
	return s
}

func (s *MmAppBindingMcpResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MmAppBindingMcpResponseBodyData struct {
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MmAppBindingMcpResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MmAppBindingMcpResponseBodyData) GoString() string {
	return s.String()
}

func (s *MmAppBindingMcpResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *MmAppBindingMcpResponseBodyData) SetSuccess(v bool) *MmAppBindingMcpResponseBodyData {
	s.Success = &v
	return s
}

func (s *MmAppBindingMcpResponseBodyData) Validate() error {
	return dara.Validate(s)
}
