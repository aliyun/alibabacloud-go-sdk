// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNumberLocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetNumberLocationResponseBody
	GetCode() *string
	SetData(v *GetNumberLocationResponseBodyData) *GetNumberLocationResponseBody
	GetData() *GetNumberLocationResponseBodyData
	SetHttpStatusCode(v int32) *GetNumberLocationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetNumberLocationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetNumberLocationResponseBody
	GetRequestId() *string
}

type GetNumberLocationResponseBody struct {
	// example:
	//
	// OK
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetNumberLocationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 584AA2E3-9AC4-561B-BC8D-C74BA11B1387
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNumberLocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNumberLocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetNumberLocationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetNumberLocationResponseBody) GetData() *GetNumberLocationResponseBodyData {
	return s.Data
}

func (s *GetNumberLocationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetNumberLocationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetNumberLocationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNumberLocationResponseBody) SetCode(v string) *GetNumberLocationResponseBody {
	s.Code = &v
	return s
}

func (s *GetNumberLocationResponseBody) SetData(v *GetNumberLocationResponseBodyData) *GetNumberLocationResponseBody {
	s.Data = v
	return s
}

func (s *GetNumberLocationResponseBody) SetHttpStatusCode(v int32) *GetNumberLocationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNumberLocationResponseBody) SetMessage(v string) *GetNumberLocationResponseBody {
	s.Message = &v
	return s
}

func (s *GetNumberLocationResponseBody) SetRequestId(v string) *GetNumberLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNumberLocationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetNumberLocationResponseBodyData struct {
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// 1312121****
	Number   *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s GetNumberLocationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetNumberLocationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNumberLocationResponseBodyData) GetCity() *string {
	return s.City
}

func (s *GetNumberLocationResponseBodyData) GetNumber() *string {
	return s.Number
}

func (s *GetNumberLocationResponseBodyData) GetProvince() *string {
	return s.Province
}

func (s *GetNumberLocationResponseBodyData) SetCity(v string) *GetNumberLocationResponseBodyData {
	s.City = &v
	return s
}

func (s *GetNumberLocationResponseBodyData) SetNumber(v string) *GetNumberLocationResponseBodyData {
	s.Number = &v
	return s
}

func (s *GetNumberLocationResponseBodyData) SetProvince(v string) *GetNumberLocationResponseBodyData {
	s.Province = &v
	return s
}

func (s *GetNumberLocationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
