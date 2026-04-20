// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMmAppBindingRagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MmAppBindingRagResponseBody
	GetCode() *string
	SetData(v *MmAppBindingRagResponseBodyData) *MmAppBindingRagResponseBody
	GetData() *MmAppBindingRagResponseBodyData
	SetHttpStatusCode(v int32) *MmAppBindingRagResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *MmAppBindingRagResponseBody
	GetMessage() *string
	SetRequestId(v string) *MmAppBindingRagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MmAppBindingRagResponseBody
	GetSuccess() *bool
}

type MmAppBindingRagResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *MmAppBindingRagResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s MmAppBindingRagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MmAppBindingRagResponseBody) GoString() string {
	return s.String()
}

func (s *MmAppBindingRagResponseBody) GetCode() *string {
	return s.Code
}

func (s *MmAppBindingRagResponseBody) GetData() *MmAppBindingRagResponseBodyData {
	return s.Data
}

func (s *MmAppBindingRagResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *MmAppBindingRagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MmAppBindingRagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MmAppBindingRagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MmAppBindingRagResponseBody) SetCode(v string) *MmAppBindingRagResponseBody {
	s.Code = &v
	return s
}

func (s *MmAppBindingRagResponseBody) SetData(v *MmAppBindingRagResponseBodyData) *MmAppBindingRagResponseBody {
	s.Data = v
	return s
}

func (s *MmAppBindingRagResponseBody) SetHttpStatusCode(v int32) *MmAppBindingRagResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *MmAppBindingRagResponseBody) SetMessage(v string) *MmAppBindingRagResponseBody {
	s.Message = &v
	return s
}

func (s *MmAppBindingRagResponseBody) SetRequestId(v string) *MmAppBindingRagResponseBody {
	s.RequestId = &v
	return s
}

func (s *MmAppBindingRagResponseBody) SetSuccess(v bool) *MmAppBindingRagResponseBody {
	s.Success = &v
	return s
}

func (s *MmAppBindingRagResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MmAppBindingRagResponseBodyData struct {
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MmAppBindingRagResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MmAppBindingRagResponseBodyData) GoString() string {
	return s.String()
}

func (s *MmAppBindingRagResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *MmAppBindingRagResponseBodyData) SetSuccess(v bool) *MmAppBindingRagResponseBodyData {
	s.Success = &v
	return s
}

func (s *MmAppBindingRagResponseBodyData) Validate() error {
	return dara.Validate(s)
}
