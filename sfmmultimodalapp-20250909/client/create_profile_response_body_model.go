// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateProfileResponseBody
	GetCode() *string
	SetData(v *CreateProfileResponseBodyData) *CreateProfileResponseBody
	GetData() *CreateProfileResponseBodyData
	SetHttpStatusCode(v int32) *CreateProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateProfileResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateProfileResponseBody
	GetSuccess() *bool
}

type CreateProfileResponseBody struct {
	// example:
	//
	// 200
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F62225F8-DDE4-526E-A88B-08906DED5060
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProfileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateProfileResponseBody) GetData() *CreateProfileResponseBodyData {
	return s.Data
}

func (s *CreateProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateProfileResponseBody) SetCode(v string) *CreateProfileResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProfileResponseBody) SetData(v *CreateProfileResponseBodyData) *CreateProfileResponseBody {
	s.Data = v
	return s
}

func (s *CreateProfileResponseBody) SetHttpStatusCode(v int32) *CreateProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateProfileResponseBody) SetMessage(v string) *CreateProfileResponseBody {
	s.Message = &v
	return s
}

func (s *CreateProfileResponseBody) SetRequestId(v string) *CreateProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProfileResponseBody) SetSuccess(v bool) *CreateProfileResponseBody {
	s.Success = &v
	return s
}

func (s *CreateProfileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProfileResponseBodyData struct {
	// example:
	//
	// 210235A52L5248C0000N
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ST-JobTemplate
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// schema id
	//
	// example:
	//
	// 33440
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s CreateProfileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateProfileResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProfileResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateProfileResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateProfileResponseBodyData) GetSchemaId() *string {
	return s.SchemaId
}

func (s *CreateProfileResponseBodyData) SetDescription(v string) *CreateProfileResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateProfileResponseBodyData) SetName(v string) *CreateProfileResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateProfileResponseBodyData) SetSchemaId(v string) *CreateProfileResponseBodyData {
	s.SchemaId = &v
	return s
}

func (s *CreateProfileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
