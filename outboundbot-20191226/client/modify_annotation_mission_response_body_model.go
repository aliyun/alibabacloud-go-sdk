// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAnnotationMissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyAnnotationMissionResponseBody
	GetCode() *string
	SetData(v *ModifyAnnotationMissionResponseBodyData) *ModifyAnnotationMissionResponseBody
	GetData() *ModifyAnnotationMissionResponseBodyData
	SetHttpStatusCode(v int32) *ModifyAnnotationMissionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyAnnotationMissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyAnnotationMissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyAnnotationMissionResponseBody
	GetSuccess() *bool
}

type ModifyAnnotationMissionResponseBody struct {
	// Status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data
	Data *ModifyAnnotationMissionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// The operation is not allowed. User state (DIALING) does not meet expectations (READY).
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the invocation succeeded. true: The invocation succeeded. false: Failed to invoke.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAnnotationMissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAnnotationMissionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAnnotationMissionResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyAnnotationMissionResponseBody) GetData() *ModifyAnnotationMissionResponseBodyData {
	return s.Data
}

func (s *ModifyAnnotationMissionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyAnnotationMissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyAnnotationMissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAnnotationMissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyAnnotationMissionResponseBody) SetCode(v string) *ModifyAnnotationMissionResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyAnnotationMissionResponseBody) SetData(v *ModifyAnnotationMissionResponseBodyData) *ModifyAnnotationMissionResponseBody {
	s.Data = v
	return s
}

func (s *ModifyAnnotationMissionResponseBody) SetHttpStatusCode(v int32) *ModifyAnnotationMissionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyAnnotationMissionResponseBody) SetMessage(v string) *ModifyAnnotationMissionResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAnnotationMissionResponseBody) SetRequestId(v string) *ModifyAnnotationMissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAnnotationMissionResponseBody) SetSuccess(v bool) *ModifyAnnotationMissionResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyAnnotationMissionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAnnotationMissionResponseBodyData struct {
	// Error message.
	//
	// example:
	//
	// The operation is not allowed. User state (DIALING) does not meet expectations (READY).
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAnnotationMissionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyAnnotationMissionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyAnnotationMissionResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ModifyAnnotationMissionResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyAnnotationMissionResponseBodyData) SetMessage(v string) *ModifyAnnotationMissionResponseBodyData {
	s.Message = &v
	return s
}

func (s *ModifyAnnotationMissionResponseBodyData) SetSuccess(v bool) *ModifyAnnotationMissionResponseBodyData {
	s.Success = &v
	return s
}

func (s *ModifyAnnotationMissionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
