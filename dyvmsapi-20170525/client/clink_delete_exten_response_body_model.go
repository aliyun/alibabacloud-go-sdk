// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDeleteExtenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkDeleteExtenResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkDeleteExtenResponseBody
	GetCode() *string
	SetData(v *ClinkDeleteExtenResponseBodyData) *ClinkDeleteExtenResponseBody
	GetData() *ClinkDeleteExtenResponseBodyData
	SetMessage(v string) *ClinkDeleteExtenResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkDeleteExtenResponseBody
	GetRequestId() *string
}

type ClinkDeleteExtenResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkDeleteExtenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkDeleteExtenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkDeleteExtenResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkDeleteExtenResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkDeleteExtenResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkDeleteExtenResponseBody) GetData() *ClinkDeleteExtenResponseBodyData {
	return s.Data
}

func (s *ClinkDeleteExtenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkDeleteExtenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkDeleteExtenResponseBody) SetAccessDeniedDetail(v string) *ClinkDeleteExtenResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkDeleteExtenResponseBody) SetCode(v string) *ClinkDeleteExtenResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkDeleteExtenResponseBody) SetData(v *ClinkDeleteExtenResponseBodyData) *ClinkDeleteExtenResponseBody {
	s.Data = v
	return s
}

func (s *ClinkDeleteExtenResponseBody) SetMessage(v string) *ClinkDeleteExtenResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkDeleteExtenResponseBody) SetRequestId(v string) *ClinkDeleteExtenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkDeleteExtenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkDeleteExtenResponseBodyData struct {
	// 请求 id
	//
	// example:
	//
	// xxx
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
	// 话机号码
	//
	// example:
	//
	// 333
	Exten *string `json:"Exten,omitempty" xml:"Exten,omitempty"`
}

func (s ClinkDeleteExtenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkDeleteExtenResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkDeleteExtenResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkDeleteExtenResponseBodyData) GetExten() *string {
	return s.Exten
}

func (s *ClinkDeleteExtenResponseBodyData) SetClinkRequestId(v string) *ClinkDeleteExtenResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkDeleteExtenResponseBodyData) SetExten(v string) *ClinkDeleteExtenResponseBodyData {
	s.Exten = &v
	return s
}

func (s *ClinkDeleteExtenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
