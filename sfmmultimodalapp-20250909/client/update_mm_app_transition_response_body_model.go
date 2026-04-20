// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppTransitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMmAppTransitionResponseBody
	GetCode() *string
	SetData(v *UpdateMmAppTransitionResponseBodyData) *UpdateMmAppTransitionResponseBody
	GetData() *UpdateMmAppTransitionResponseBodyData
	SetHttpStatusCode(v int32) *UpdateMmAppTransitionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateMmAppTransitionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMmAppTransitionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMmAppTransitionResponseBody
	GetSuccess() *bool
}

type UpdateMmAppTransitionResponseBody struct {
	// example:
	//
	// 200
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateMmAppTransitionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// EDD6123F-0122-5FBF-9A7E-097F319CF478
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMmAppTransitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppTransitionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMmAppTransitionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMmAppTransitionResponseBody) GetData() *UpdateMmAppTransitionResponseBodyData {
	return s.Data
}

func (s *UpdateMmAppTransitionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMmAppTransitionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMmAppTransitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMmAppTransitionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMmAppTransitionResponseBody) SetCode(v string) *UpdateMmAppTransitionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMmAppTransitionResponseBody) SetData(v *UpdateMmAppTransitionResponseBodyData) *UpdateMmAppTransitionResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMmAppTransitionResponseBody) SetHttpStatusCode(v int32) *UpdateMmAppTransitionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMmAppTransitionResponseBody) SetMessage(v string) *UpdateMmAppTransitionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMmAppTransitionResponseBody) SetRequestId(v string) *UpdateMmAppTransitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMmAppTransitionResponseBody) SetSuccess(v bool) *UpdateMmAppTransitionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMmAppTransitionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMmAppTransitionResponseBodyData struct {
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMmAppTransitionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppTransitionResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateMmAppTransitionResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMmAppTransitionResponseBodyData) SetSuccess(v bool) *UpdateMmAppTransitionResponseBodyData {
	s.Success = &v
	return s
}

func (s *UpdateMmAppTransitionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
