// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppRagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMmAppRagResponseBody
	GetCode() *string
	SetData(v *UpdateMmAppRagResponseBodyData) *UpdateMmAppRagResponseBody
	GetData() *UpdateMmAppRagResponseBodyData
	SetHttpStatusCode(v int32) *UpdateMmAppRagResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateMmAppRagResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMmAppRagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMmAppRagResponseBody
	GetSuccess() *bool
}

type UpdateMmAppRagResponseBody struct {
	// example:
	//
	// 200
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateMmAppRagResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 31033EC0-6968-5610-8328-708B59508E5A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMmAppRagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRagResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRagResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMmAppRagResponseBody) GetData() *UpdateMmAppRagResponseBodyData {
	return s.Data
}

func (s *UpdateMmAppRagResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMmAppRagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMmAppRagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMmAppRagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMmAppRagResponseBody) SetCode(v string) *UpdateMmAppRagResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMmAppRagResponseBody) SetData(v *UpdateMmAppRagResponseBodyData) *UpdateMmAppRagResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMmAppRagResponseBody) SetHttpStatusCode(v int32) *UpdateMmAppRagResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMmAppRagResponseBody) SetMessage(v string) *UpdateMmAppRagResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMmAppRagResponseBody) SetRequestId(v string) *UpdateMmAppRagResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMmAppRagResponseBody) SetSuccess(v bool) *UpdateMmAppRagResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMmAppRagResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMmAppRagResponseBodyData struct {
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMmAppRagResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRagResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRagResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMmAppRagResponseBodyData) SetSuccess(v bool) *UpdateMmAppRagResponseBodyData {
	s.Success = &v
	return s
}

func (s *UpdateMmAppRagResponseBodyData) Validate() error {
	return dara.Validate(s)
}
