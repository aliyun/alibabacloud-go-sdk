// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMapPkFromHidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MapPkFromHidResponseBody
	GetCode() *string
	SetData(v *MapPkFromHidResponseBodyData) *MapPkFromHidResponseBody
	GetData() *MapPkFromHidResponseBodyData
	SetHttpCode(v string) *MapPkFromHidResponseBody
	GetHttpCode() *string
	SetMessage(v string) *MapPkFromHidResponseBody
	GetMessage() *string
	SetRequestId(v string) *MapPkFromHidResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MapPkFromHidResponseBody
	GetSuccess() *bool
}

type MapPkFromHidResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *MapPkFromHidResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpCode  *string                       `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MapPkFromHidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MapPkFromHidResponseBody) GoString() string {
	return s.String()
}

func (s *MapPkFromHidResponseBody) GetCode() *string {
	return s.Code
}

func (s *MapPkFromHidResponseBody) GetData() *MapPkFromHidResponseBodyData {
	return s.Data
}

func (s *MapPkFromHidResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *MapPkFromHidResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MapPkFromHidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MapPkFromHidResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MapPkFromHidResponseBody) SetCode(v string) *MapPkFromHidResponseBody {
	s.Code = &v
	return s
}

func (s *MapPkFromHidResponseBody) SetData(v *MapPkFromHidResponseBodyData) *MapPkFromHidResponseBody {
	s.Data = v
	return s
}

func (s *MapPkFromHidResponseBody) SetHttpCode(v string) *MapPkFromHidResponseBody {
	s.HttpCode = &v
	return s
}

func (s *MapPkFromHidResponseBody) SetMessage(v string) *MapPkFromHidResponseBody {
	s.Message = &v
	return s
}

func (s *MapPkFromHidResponseBody) SetRequestId(v string) *MapPkFromHidResponseBody {
	s.RequestId = &v
	return s
}

func (s *MapPkFromHidResponseBody) SetSuccess(v bool) *MapPkFromHidResponseBody {
	s.Success = &v
	return s
}

func (s *MapPkFromHidResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MapPkFromHidResponseBodyData struct {
	Hid        *string `json:"Hid,omitempty" xml:"Hid,omitempty"`
	MappingSrc *string `json:"MappingSrc,omitempty" xml:"MappingSrc,omitempty"`
	Pk         *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s MapPkFromHidResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MapPkFromHidResponseBodyData) GoString() string {
	return s.String()
}

func (s *MapPkFromHidResponseBodyData) GetHid() *string {
	return s.Hid
}

func (s *MapPkFromHidResponseBodyData) GetMappingSrc() *string {
	return s.MappingSrc
}

func (s *MapPkFromHidResponseBodyData) GetPk() *string {
	return s.Pk
}

func (s *MapPkFromHidResponseBodyData) SetHid(v string) *MapPkFromHidResponseBodyData {
	s.Hid = &v
	return s
}

func (s *MapPkFromHidResponseBodyData) SetMappingSrc(v string) *MapPkFromHidResponseBodyData {
	s.MappingSrc = &v
	return s
}

func (s *MapPkFromHidResponseBodyData) SetPk(v string) *MapPkFromHidResponseBodyData {
	s.Pk = &v
	return s
}

func (s *MapPkFromHidResponseBodyData) Validate() error {
	return dara.Validate(s)
}
