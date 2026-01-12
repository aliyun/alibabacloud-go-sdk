// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHdrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHdrResponseBody
	GetCode() *string
	SetData(v []*ListHdrResponseBodyData) *ListHdrResponseBody
	GetData() []*ListHdrResponseBodyData
	SetErrorName(v string) *ListHdrResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *ListHdrResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *ListHdrResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHdrResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListHdrResponseBody
	GetSuccess() *bool
}

type ListHdrResponseBody struct {
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListHdrResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorName *string                    `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                     `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListHdrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHdrResponseBody) GoString() string {
	return s.String()
}

func (s *ListHdrResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHdrResponseBody) GetData() []*ListHdrResponseBodyData {
	return s.Data
}

func (s *ListHdrResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *ListHdrResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListHdrResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHdrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHdrResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListHdrResponseBody) SetCode(v string) *ListHdrResponseBody {
	s.Code = &v
	return s
}

func (s *ListHdrResponseBody) SetData(v []*ListHdrResponseBodyData) *ListHdrResponseBody {
	s.Data = v
	return s
}

func (s *ListHdrResponseBody) SetErrorName(v string) *ListHdrResponseBody {
	s.ErrorName = &v
	return s
}

func (s *ListHdrResponseBody) SetHttpCode(v int32) *ListHdrResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListHdrResponseBody) SetMessage(v string) *ListHdrResponseBody {
	s.Message = &v
	return s
}

func (s *ListHdrResponseBody) SetRequestId(v string) *ListHdrResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHdrResponseBody) SetSuccess(v bool) *ListHdrResponseBody {
	s.Success = &v
	return s
}

func (s *ListHdrResponseBody) Validate() error {
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

type ListHdrResponseBodyData struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ListHdrResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListHdrResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHdrResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListHdrResponseBodyData) GetPath() *string {
	return s.Path
}

func (s *ListHdrResponseBodyData) SetName(v string) *ListHdrResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListHdrResponseBodyData) SetPath(v string) *ListHdrResponseBodyData {
	s.Path = &v
	return s
}

func (s *ListHdrResponseBodyData) Validate() error {
	return dara.Validate(s)
}
