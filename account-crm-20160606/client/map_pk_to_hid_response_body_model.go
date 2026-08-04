// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMapPkToHidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MapPkToHidResponseBody
	GetCode() *string
	SetData(v *MapPkToHidResponseBodyData) *MapPkToHidResponseBody
	GetData() *MapPkToHidResponseBodyData
	SetHttpCode(v string) *MapPkToHidResponseBody
	GetHttpCode() *string
	SetMessage(v string) *MapPkToHidResponseBody
	GetMessage() *string
	SetRequestId(v string) *MapPkToHidResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MapPkToHidResponseBody
	GetSuccess() *bool
}

type MapPkToHidResponseBody struct {
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *MapPkToHidResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpCode  *string                     `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MapPkToHidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MapPkToHidResponseBody) GoString() string {
	return s.String()
}

func (s *MapPkToHidResponseBody) GetCode() *string {
	return s.Code
}

func (s *MapPkToHidResponseBody) GetData() *MapPkToHidResponseBodyData {
	return s.Data
}

func (s *MapPkToHidResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *MapPkToHidResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MapPkToHidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MapPkToHidResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MapPkToHidResponseBody) SetCode(v string) *MapPkToHidResponseBody {
	s.Code = &v
	return s
}

func (s *MapPkToHidResponseBody) SetData(v *MapPkToHidResponseBodyData) *MapPkToHidResponseBody {
	s.Data = v
	return s
}

func (s *MapPkToHidResponseBody) SetHttpCode(v string) *MapPkToHidResponseBody {
	s.HttpCode = &v
	return s
}

func (s *MapPkToHidResponseBody) SetMessage(v string) *MapPkToHidResponseBody {
	s.Message = &v
	return s
}

func (s *MapPkToHidResponseBody) SetRequestId(v string) *MapPkToHidResponseBody {
	s.RequestId = &v
	return s
}

func (s *MapPkToHidResponseBody) SetSuccess(v bool) *MapPkToHidResponseBody {
	s.Success = &v
	return s
}

func (s *MapPkToHidResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MapPkToHidResponseBodyData struct {
	Hid        *string `json:"Hid,omitempty" xml:"Hid,omitempty"`
	MappingSrc *string `json:"MappingSrc,omitempty" xml:"MappingSrc,omitempty"`
	Pk         *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s MapPkToHidResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MapPkToHidResponseBodyData) GoString() string {
	return s.String()
}

func (s *MapPkToHidResponseBodyData) GetHid() *string {
	return s.Hid
}

func (s *MapPkToHidResponseBodyData) GetMappingSrc() *string {
	return s.MappingSrc
}

func (s *MapPkToHidResponseBodyData) GetPk() *string {
	return s.Pk
}

func (s *MapPkToHidResponseBodyData) SetHid(v string) *MapPkToHidResponseBodyData {
	s.Hid = &v
	return s
}

func (s *MapPkToHidResponseBodyData) SetMappingSrc(v string) *MapPkToHidResponseBodyData {
	s.MappingSrc = &v
	return s
}

func (s *MapPkToHidResponseBodyData) SetPk(v string) *MapPkToHidResponseBodyData {
	s.Pk = &v
	return s
}

func (s *MapPkToHidResponseBodyData) Validate() error {
	return dara.Validate(s)
}
