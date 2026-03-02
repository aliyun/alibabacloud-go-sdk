// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteProfileResponseBody
	GetCode() *string
	SetData(v *DeleteProfileResponseBodyData) *DeleteProfileResponseBody
	GetData() *DeleteProfileResponseBodyData
	SetHttpStatusCode(v int32) *DeleteProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteProfileResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteProfileResponseBody
	GetSuccess() *bool
}

type DeleteProfileResponseBody struct {
	// example:
	//
	// 200
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EC4762F9-8109-5DE0-A3E2-27957A4F4183
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProfileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteProfileResponseBody) GetData() *DeleteProfileResponseBodyData {
	return s.Data
}

func (s *DeleteProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteProfileResponseBody) SetCode(v string) *DeleteProfileResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteProfileResponseBody) SetData(v *DeleteProfileResponseBodyData) *DeleteProfileResponseBody {
	s.Data = v
	return s
}

func (s *DeleteProfileResponseBody) SetHttpStatusCode(v int32) *DeleteProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteProfileResponseBody) SetMessage(v string) *DeleteProfileResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteProfileResponseBody) SetRequestId(v string) *DeleteProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProfileResponseBody) SetSuccess(v bool) *DeleteProfileResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteProfileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteProfileResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// schema id
	//
	// example:
	//
	// f7e110b6d4359977d1
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s DeleteProfileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteProfileResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteProfileResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DeleteProfileResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DeleteProfileResponseBodyData) GetSchemaId() *string {
	return s.SchemaId
}

func (s *DeleteProfileResponseBodyData) SetDescription(v string) *DeleteProfileResponseBodyData {
	s.Description = &v
	return s
}

func (s *DeleteProfileResponseBodyData) SetName(v string) *DeleteProfileResponseBodyData {
	s.Name = &v
	return s
}

func (s *DeleteProfileResponseBodyData) SetSchemaId(v string) *DeleteProfileResponseBodyData {
	s.SchemaId = &v
	return s
}

func (s *DeleteProfileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
