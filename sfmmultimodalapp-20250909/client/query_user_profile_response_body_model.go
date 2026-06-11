// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryUserProfileResponseBody
	GetCode() *string
	SetData(v *QueryUserProfileResponseBodyData) *QueryUserProfileResponseBody
	GetData() *QueryUserProfileResponseBodyData
	SetHttpStatusCode(v int32) *QueryUserProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryUserProfileResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryUserProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryUserProfileResponseBody
	GetSuccess() *bool
}

type QueryUserProfileResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *QueryUserProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUserProfileResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryUserProfileResponseBody) GetData() *QueryUserProfileResponseBodyData {
	return s.Data
}

func (s *QueryUserProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryUserProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryUserProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUserProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUserProfileResponseBody) SetCode(v string) *QueryUserProfileResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUserProfileResponseBody) SetData(v *QueryUserProfileResponseBodyData) *QueryUserProfileResponseBody {
	s.Data = v
	return s
}

func (s *QueryUserProfileResponseBody) SetHttpStatusCode(v int32) *QueryUserProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryUserProfileResponseBody) SetMessage(v string) *QueryUserProfileResponseBody {
	s.Message = &v
	return s
}

func (s *QueryUserProfileResponseBody) SetRequestId(v string) *QueryUserProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserProfileResponseBody) SetSuccess(v bool) *QueryUserProfileResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUserProfileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryUserProfileResponseBodyData struct {
	Attributes  []*QueryUserProfileResponseBodyDataAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Description *string                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string                                       `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryUserProfileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryUserProfileResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryUserProfileResponseBodyData) GetAttributes() []*QueryUserProfileResponseBodyDataAttributes {
	return s.Attributes
}

func (s *QueryUserProfileResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *QueryUserProfileResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QueryUserProfileResponseBodyData) SetAttributes(v []*QueryUserProfileResponseBodyDataAttributes) *QueryUserProfileResponseBodyData {
	s.Attributes = v
	return s
}

func (s *QueryUserProfileResponseBodyData) SetDescription(v string) *QueryUserProfileResponseBodyData {
	s.Description = &v
	return s
}

func (s *QueryUserProfileResponseBodyData) SetName(v string) *QueryUserProfileResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryUserProfileResponseBodyData) Validate() error {
	if s.Attributes != nil {
		for _, item := range s.Attributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryUserProfileResponseBodyDataAttributes struct {
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryUserProfileResponseBodyDataAttributes) String() string {
	return dara.Prettify(s)
}

func (s QueryUserProfileResponseBodyDataAttributes) GoString() string {
	return s.String()
}

func (s *QueryUserProfileResponseBodyDataAttributes) GetId() *string {
	return s.Id
}

func (s *QueryUserProfileResponseBodyDataAttributes) GetName() *string {
	return s.Name
}

func (s *QueryUserProfileResponseBodyDataAttributes) GetValue() *string {
	return s.Value
}

func (s *QueryUserProfileResponseBodyDataAttributes) SetId(v string) *QueryUserProfileResponseBodyDataAttributes {
	s.Id = &v
	return s
}

func (s *QueryUserProfileResponseBodyDataAttributes) SetName(v string) *QueryUserProfileResponseBodyDataAttributes {
	s.Name = &v
	return s
}

func (s *QueryUserProfileResponseBodyDataAttributes) SetValue(v string) *QueryUserProfileResponseBodyDataAttributes {
	s.Value = &v
	return s
}

func (s *QueryUserProfileResponseBodyDataAttributes) Validate() error {
	return dara.Validate(s)
}
