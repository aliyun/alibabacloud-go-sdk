// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageEventOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateImageEventOperationResponseBody
	GetCode() *string
	SetData(v *UpdateImageEventOperationResponseBodyData) *UpdateImageEventOperationResponseBody
	GetData() *UpdateImageEventOperationResponseBodyData
	SetMessage(v string) *UpdateImageEventOperationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateImageEventOperationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateImageEventOperationResponseBody
	GetSuccess() *bool
}

type UpdateImageEventOperationResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *UpdateImageEventOperationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-B12447DD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateImageEventOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageEventOperationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageEventOperationResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateImageEventOperationResponseBody) GetData() *UpdateImageEventOperationResponseBodyData {
	return s.Data
}

func (s *UpdateImageEventOperationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateImageEventOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateImageEventOperationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateImageEventOperationResponseBody) SetCode(v string) *UpdateImageEventOperationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateImageEventOperationResponseBody) SetData(v *UpdateImageEventOperationResponseBodyData) *UpdateImageEventOperationResponseBody {
	s.Data = v
	return s
}

func (s *UpdateImageEventOperationResponseBody) SetMessage(v string) *UpdateImageEventOperationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateImageEventOperationResponseBody) SetRequestId(v string) *UpdateImageEventOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImageEventOperationResponseBody) SetSuccess(v bool) *UpdateImageEventOperationResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateImageEventOperationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateImageEventOperationResponseBodyData struct {
	// The ID of the alert handling rule, which is the same as the value of the Id request parameter.
	//
	// example:
	//
	// 67429
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateImageEventOperationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageEventOperationResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateImageEventOperationResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *UpdateImageEventOperationResponseBodyData) SetId(v int64) *UpdateImageEventOperationResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateImageEventOperationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
