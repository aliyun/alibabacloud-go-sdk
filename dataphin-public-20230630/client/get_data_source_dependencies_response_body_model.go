// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataSourceDependenciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataSourceDependenciesResponseBody
	GetCode() *string
	SetDependencyList(v []*GetDataSourceDependenciesResponseBodyDependencyList) *GetDataSourceDependenciesResponseBody
	GetDependencyList() []*GetDataSourceDependenciesResponseBodyDependencyList
	SetHttpStatusCode(v int32) *GetDataSourceDependenciesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataSourceDependenciesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDataSourceDependenciesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataSourceDependenciesResponseBody
	GetSuccess() *bool
}

type GetDataSourceDependenciesResponseBody struct {
	// example:
	//
	// OK
	Code           *string                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	DependencyList []*GetDataSourceDependenciesResponseBodyDependencyList `json:"DependencyList,omitempty" xml:"DependencyList,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataSourceDependenciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceDependenciesResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSourceDependenciesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataSourceDependenciesResponseBody) GetDependencyList() []*GetDataSourceDependenciesResponseBodyDependencyList {
	return s.DependencyList
}

func (s *GetDataSourceDependenciesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataSourceDependenciesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataSourceDependenciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataSourceDependenciesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataSourceDependenciesResponseBody) SetCode(v string) *GetDataSourceDependenciesResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataSourceDependenciesResponseBody) SetDependencyList(v []*GetDataSourceDependenciesResponseBodyDependencyList) *GetDataSourceDependenciesResponseBody {
	s.DependencyList = v
	return s
}

func (s *GetDataSourceDependenciesResponseBody) SetHttpStatusCode(v int32) *GetDataSourceDependenciesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataSourceDependenciesResponseBody) SetMessage(v string) *GetDataSourceDependenciesResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataSourceDependenciesResponseBody) SetRequestId(v string) *GetDataSourceDependenciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataSourceDependenciesResponseBody) SetSuccess(v bool) *GetDataSourceDependenciesResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataSourceDependenciesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataSourceDependenciesResponseBodyDependencyList struct {
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// jytest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// pipeline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// pipeline
	TypeCode *string `json:"TypeCode,omitempty" xml:"TypeCode,omitempty"`
}

func (s GetDataSourceDependenciesResponseBodyDependencyList) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceDependenciesResponseBodyDependencyList) GoString() string {
	return s.String()
}

func (s *GetDataSourceDependenciesResponseBodyDependencyList) GetId() *string {
	return s.Id
}

func (s *GetDataSourceDependenciesResponseBodyDependencyList) GetName() *string {
	return s.Name
}

func (s *GetDataSourceDependenciesResponseBodyDependencyList) GetType() *string {
	return s.Type
}

func (s *GetDataSourceDependenciesResponseBodyDependencyList) GetTypeCode() *string {
	return s.TypeCode
}

func (s *GetDataSourceDependenciesResponseBodyDependencyList) SetId(v string) *GetDataSourceDependenciesResponseBodyDependencyList {
	s.Id = &v
	return s
}

func (s *GetDataSourceDependenciesResponseBodyDependencyList) SetName(v string) *GetDataSourceDependenciesResponseBodyDependencyList {
	s.Name = &v
	return s
}

func (s *GetDataSourceDependenciesResponseBodyDependencyList) SetType(v string) *GetDataSourceDependenciesResponseBodyDependencyList {
	s.Type = &v
	return s
}

func (s *GetDataSourceDependenciesResponseBodyDependencyList) SetTypeCode(v string) *GetDataSourceDependenciesResponseBodyDependencyList {
	s.TypeCode = &v
	return s
}

func (s *GetDataSourceDependenciesResponseBodyDependencyList) Validate() error {
	return dara.Validate(s)
}
