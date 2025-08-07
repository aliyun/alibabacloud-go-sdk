// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadClassNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadClassNameResponseBody
	GetCode() *string
	SetData(v []*ReadClassNameResponseBodyData) *ReadClassNameResponseBody
	GetData() []*ReadClassNameResponseBodyData
	SetMessage(v string) *ReadClassNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadClassNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadClassNameResponseBody
	GetSuccess() *bool
}

type ReadClassNameResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ReadClassNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadClassNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadClassNameResponseBody) GoString() string {
	return s.String()
}

func (s *ReadClassNameResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadClassNameResponseBody) GetData() []*ReadClassNameResponseBodyData {
	return s.Data
}

func (s *ReadClassNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadClassNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadClassNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadClassNameResponseBody) SetCode(v string) *ReadClassNameResponseBody {
	s.Code = &v
	return s
}

func (s *ReadClassNameResponseBody) SetData(v []*ReadClassNameResponseBodyData) *ReadClassNameResponseBody {
	s.Data = v
	return s
}

func (s *ReadClassNameResponseBody) SetMessage(v string) *ReadClassNameResponseBody {
	s.Message = &v
	return s
}

func (s *ReadClassNameResponseBody) SetRequestId(v string) *ReadClassNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadClassNameResponseBody) SetSuccess(v bool) *ReadClassNameResponseBody {
	s.Success = &v
	return s
}

func (s *ReadClassNameResponseBody) Validate() error {
	return dara.Validate(s)
}

type ReadClassNameResponseBodyData struct {
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ReadClassNameResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadClassNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadClassNameResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ReadClassNameResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ReadClassNameResponseBodyData) SetId(v int64) *ReadClassNameResponseBodyData {
	s.Id = &v
	return s
}

func (s *ReadClassNameResponseBodyData) SetName(v string) *ReadClassNameResponseBodyData {
	s.Name = &v
	return s
}

func (s *ReadClassNameResponseBodyData) Validate() error {
	return dara.Validate(s)
}
