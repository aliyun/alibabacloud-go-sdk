// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllDepartmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAllDepartmentResponseBody
	GetCode() *string
	SetData(v []*GetAllDepartmentResponseBodyData) *GetAllDepartmentResponseBody
	GetData() []*GetAllDepartmentResponseBodyData
	SetHttpStatusCode(v int32) *GetAllDepartmentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAllDepartmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAllDepartmentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAllDepartmentResponseBody
	GetSuccess() *bool
}

type GetAllDepartmentResponseBody struct {
	// example:
	//
	// 200
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetAllDepartmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 98B032F5-6473-4EAC-8BA8-C28993513A1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAllDepartmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAllDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllDepartmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAllDepartmentResponseBody) GetData() []*GetAllDepartmentResponseBodyData {
	return s.Data
}

func (s *GetAllDepartmentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAllDepartmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAllDepartmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAllDepartmentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAllDepartmentResponseBody) SetCode(v string) *GetAllDepartmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetAllDepartmentResponseBody) SetData(v []*GetAllDepartmentResponseBodyData) *GetAllDepartmentResponseBody {
	s.Data = v
	return s
}

func (s *GetAllDepartmentResponseBody) SetHttpStatusCode(v int32) *GetAllDepartmentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAllDepartmentResponseBody) SetMessage(v string) *GetAllDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetAllDepartmentResponseBody) SetRequestId(v string) *GetAllDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllDepartmentResponseBody) SetSuccess(v bool) *GetAllDepartmentResponseBody {
	s.Success = &v
	return s
}

func (s *GetAllDepartmentResponseBody) Validate() error {
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

type GetAllDepartmentResponseBodyData struct {
	// example:
	//
	// 10
	DepartmentId   *int64  `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
}

func (s GetAllDepartmentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAllDepartmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAllDepartmentResponseBodyData) GetDepartmentId() *int64 {
	return s.DepartmentId
}

func (s *GetAllDepartmentResponseBodyData) GetDepartmentName() *string {
	return s.DepartmentName
}

func (s *GetAllDepartmentResponseBodyData) SetDepartmentId(v int64) *GetAllDepartmentResponseBodyData {
	s.DepartmentId = &v
	return s
}

func (s *GetAllDepartmentResponseBodyData) SetDepartmentName(v string) *GetAllDepartmentResponseBodyData {
	s.DepartmentName = &v
	return s
}

func (s *GetAllDepartmentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
