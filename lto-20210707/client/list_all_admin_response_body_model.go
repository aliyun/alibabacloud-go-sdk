// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllAdminResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAllAdminResponseBody
	GetCode() *string
	SetData(v []*ListAllAdminResponseBodyData) *ListAllAdminResponseBody
	GetData() []*ListAllAdminResponseBodyData
	SetHttpStatusCode(v int32) *ListAllAdminResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAllAdminResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAllAdminResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAllAdminResponseBody
	GetSuccess() *bool
}

type ListAllAdminResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListAllAdminResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAllAdminResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllAdminResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllAdminResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAllAdminResponseBody) GetData() []*ListAllAdminResponseBodyData {
	return s.Data
}

func (s *ListAllAdminResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAllAdminResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAllAdminResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllAdminResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAllAdminResponseBody) SetCode(v string) *ListAllAdminResponseBody {
	s.Code = &v
	return s
}

func (s *ListAllAdminResponseBody) SetData(v []*ListAllAdminResponseBodyData) *ListAllAdminResponseBody {
	s.Data = v
	return s
}

func (s *ListAllAdminResponseBody) SetHttpStatusCode(v int32) *ListAllAdminResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAllAdminResponseBody) SetMessage(v string) *ListAllAdminResponseBody {
	s.Message = &v
	return s
}

func (s *ListAllAdminResponseBody) SetRequestId(v string) *ListAllAdminResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllAdminResponseBody) SetSuccess(v bool) *ListAllAdminResponseBody {
	s.Success = &v
	return s
}

func (s *ListAllAdminResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAllAdminResponseBodyData struct {
	AdminId *string `json:"AdminId,omitempty" xml:"AdminId,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAllAdminResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAllAdminResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAllAdminResponseBodyData) GetAdminId() *string {
	return s.AdminId
}

func (s *ListAllAdminResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListAllAdminResponseBodyData) SetAdminId(v string) *ListAllAdminResponseBodyData {
	s.AdminId = &v
	return s
}

func (s *ListAllAdminResponseBodyData) SetName(v string) *ListAllAdminResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListAllAdminResponseBodyData) Validate() error {
	return dara.Validate(s)
}
