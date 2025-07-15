// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCategoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCategoriesResponseBody
	GetCode() *string
	SetData(v string) *ListCategoriesResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *ListCategoriesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCategoriesResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListCategoriesResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListCategoriesResponseBody
	GetRequestId() *string
}

type ListCategoriesResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// DE803553-8AA9-4B9D-9E4E-A82BC69EDCEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCategoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCategoriesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCategoriesResponseBody) GetData() *string {
	return s.Data
}

func (s *ListCategoriesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCategoriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCategoriesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListCategoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCategoriesResponseBody) SetCode(v string) *ListCategoriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCategoriesResponseBody) SetData(v string) *ListCategoriesResponseBody {
	s.Data = &v
	return s
}

func (s *ListCategoriesResponseBody) SetHttpStatusCode(v int32) *ListCategoriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCategoriesResponseBody) SetMessage(v string) *ListCategoriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCategoriesResponseBody) SetParams(v []*string) *ListCategoriesResponseBody {
	s.Params = v
	return s
}

func (s *ListCategoriesResponseBody) SetRequestId(v string) *ListCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCategoriesResponseBody) Validate() error {
	return dara.Validate(s)
}
