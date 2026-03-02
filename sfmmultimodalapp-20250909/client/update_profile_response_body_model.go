// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateProfileResponseBody
	GetCode() *string
	SetData(v *UpdateProfileResponseBodyData) *UpdateProfileResponseBody
	GetData() *UpdateProfileResponseBodyData
	SetHttpStatusCode(v int32) *UpdateProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateProfileResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateProfileResponseBody
	GetSuccess() *bool
}

type UpdateProfileResponseBody struct {
	// example:
	//
	// 200
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProfileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateProfileResponseBody) GetData() *UpdateProfileResponseBodyData {
	return s.Data
}

func (s *UpdateProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateProfileResponseBody) SetCode(v string) *UpdateProfileResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateProfileResponseBody) SetData(v *UpdateProfileResponseBodyData) *UpdateProfileResponseBody {
	s.Data = v
	return s
}

func (s *UpdateProfileResponseBody) SetHttpStatusCode(v int32) *UpdateProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateProfileResponseBody) SetMessage(v string) *UpdateProfileResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateProfileResponseBody) SetRequestId(v string) *UpdateProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProfileResponseBody) SetSuccess(v bool) *UpdateProfileResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateProfileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateProfileResponseBodyData struct {
	// example:
	//
	// M21436N3000X
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ceshi
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 39536
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateProfileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateProfileResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateProfileResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateProfileResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateProfileResponseBodyData) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateProfileResponseBodyData) SetDescription(v string) *UpdateProfileResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateProfileResponseBodyData) SetName(v string) *UpdateProfileResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateProfileResponseBodyData) SetSchemaId(v string) *UpdateProfileResponseBodyData {
	s.SchemaId = &v
	return s
}

func (s *UpdateProfileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
