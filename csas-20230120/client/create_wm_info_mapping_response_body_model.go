// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWmInfoMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateWmInfoMappingResponseBodyData) *CreateWmInfoMappingResponseBody
	GetData() *CreateWmInfoMappingResponseBodyData
	SetRequestId(v string) *CreateWmInfoMappingResponseBody
	GetRequestId() *string
}

type CreateWmInfoMappingResponseBody struct {
	Data *CreateWmInfoMappingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 7E9D7ACD-53D5-56EF-A913-79D148D06299
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWmInfoMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWmInfoMappingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWmInfoMappingResponseBody) GetData() *CreateWmInfoMappingResponseBodyData {
	return s.Data
}

func (s *CreateWmInfoMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWmInfoMappingResponseBody) SetData(v *CreateWmInfoMappingResponseBodyData) *CreateWmInfoMappingResponseBody {
	s.Data = v
	return s
}

func (s *CreateWmInfoMappingResponseBody) SetRequestId(v string) *CreateWmInfoMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWmInfoMappingResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateWmInfoMappingResponseBodyData struct {
	// example:
	//
	// 123***
	WmInfoUint *int64 `json:"WmInfoUint,omitempty" xml:"WmInfoUint,omitempty"`
}

func (s CreateWmInfoMappingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateWmInfoMappingResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWmInfoMappingResponseBodyData) GetWmInfoUint() *int64 {
	return s.WmInfoUint
}

func (s *CreateWmInfoMappingResponseBodyData) SetWmInfoUint(v int64) *CreateWmInfoMappingResponseBodyData {
	s.WmInfoUint = &v
	return s
}

func (s *CreateWmInfoMappingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
