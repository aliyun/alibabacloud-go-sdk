// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllSystemContractResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAllSystemContractResponseBody
	GetCode() *string
	SetData(v []*ListAllSystemContractResponseBodyData) *ListAllSystemContractResponseBody
	GetData() []*ListAllSystemContractResponseBodyData
	SetHttpStatusCode(v int32) *ListAllSystemContractResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAllSystemContractResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAllSystemContractResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAllSystemContractResponseBody
	GetSuccess() *bool
}

type ListAllSystemContractResponseBody struct {
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListAllSystemContractResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAllSystemContractResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllSystemContractResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllSystemContractResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAllSystemContractResponseBody) GetData() []*ListAllSystemContractResponseBodyData {
	return s.Data
}

func (s *ListAllSystemContractResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAllSystemContractResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAllSystemContractResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllSystemContractResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAllSystemContractResponseBody) SetCode(v string) *ListAllSystemContractResponseBody {
	s.Code = &v
	return s
}

func (s *ListAllSystemContractResponseBody) SetData(v []*ListAllSystemContractResponseBodyData) *ListAllSystemContractResponseBody {
	s.Data = v
	return s
}

func (s *ListAllSystemContractResponseBody) SetHttpStatusCode(v int32) *ListAllSystemContractResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAllSystemContractResponseBody) SetMessage(v string) *ListAllSystemContractResponseBody {
	s.Message = &v
	return s
}

func (s *ListAllSystemContractResponseBody) SetRequestId(v string) *ListAllSystemContractResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllSystemContractResponseBody) SetSuccess(v bool) *ListAllSystemContractResponseBody {
	s.Success = &v
	return s
}

func (s *ListAllSystemContractResponseBody) Validate() error {
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

type ListAllSystemContractResponseBodyData struct {
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SystemContractId *string `json:"SystemContractId,omitempty" xml:"SystemContractId,omitempty"`
}

func (s ListAllSystemContractResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAllSystemContractResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAllSystemContractResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListAllSystemContractResponseBodyData) GetSystemContractId() *string {
	return s.SystemContractId
}

func (s *ListAllSystemContractResponseBodyData) SetName(v string) *ListAllSystemContractResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListAllSystemContractResponseBodyData) SetSystemContractId(v string) *ListAllSystemContractResponseBodyData {
	s.SystemContractId = &v
	return s
}

func (s *ListAllSystemContractResponseBodyData) Validate() error {
	return dara.Validate(s)
}
