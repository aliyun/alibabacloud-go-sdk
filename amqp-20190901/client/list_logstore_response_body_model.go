// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogstoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListLogstoreResponseBody
	GetCode() *int32
	SetData(v *ListLogstoreResponseBodyData) *ListLogstoreResponseBody
	GetData() *ListLogstoreResponseBodyData
	SetMessage(v string) *ListLogstoreResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLogstoreResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLogstoreResponseBody
	GetSuccess() *bool
}

type ListLogstoreResponseBody struct {
	Code      *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListLogstoreResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLogstoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLogstoreResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogstoreResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListLogstoreResponseBody) GetData() *ListLogstoreResponseBodyData {
	return s.Data
}

func (s *ListLogstoreResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLogstoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLogstoreResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLogstoreResponseBody) SetCode(v int32) *ListLogstoreResponseBody {
	s.Code = &v
	return s
}

func (s *ListLogstoreResponseBody) SetData(v *ListLogstoreResponseBodyData) *ListLogstoreResponseBody {
	s.Data = v
	return s
}

func (s *ListLogstoreResponseBody) SetMessage(v string) *ListLogstoreResponseBody {
	s.Message = &v
	return s
}

func (s *ListLogstoreResponseBody) SetRequestId(v string) *ListLogstoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogstoreResponseBody) SetSuccess(v bool) *ListLogstoreResponseBody {
	s.Success = &v
	return s
}

func (s *ListLogstoreResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLogstoreResponseBodyData struct {
	Logstores []*string `json:"Logstores,omitempty" xml:"Logstores,omitempty" type:"Repeated"`
}

func (s ListLogstoreResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLogstoreResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLogstoreResponseBodyData) GetLogstores() []*string {
	return s.Logstores
}

func (s *ListLogstoreResponseBodyData) SetLogstores(v []*string) *ListLogstoreResponseBodyData {
	s.Logstores = v
	return s
}

func (s *ListLogstoreResponseBodyData) Validate() error {
	return dara.Validate(s)
}
