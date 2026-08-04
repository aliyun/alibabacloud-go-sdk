// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMapToHavanaBindIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MapToHavanaBindIdResponseBody
	GetCode() *string
	SetData(v []*MapToHavanaBindIdResponseBodyData) *MapToHavanaBindIdResponseBody
	GetData() []*MapToHavanaBindIdResponseBodyData
	SetHttpCode(v string) *MapToHavanaBindIdResponseBody
	GetHttpCode() *string
	SetMessage(v string) *MapToHavanaBindIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *MapToHavanaBindIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MapToHavanaBindIdResponseBody
	GetSuccess() *bool
}

type MapToHavanaBindIdResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*MapToHavanaBindIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpCode  *string                              `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MapToHavanaBindIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MapToHavanaBindIdResponseBody) GoString() string {
	return s.String()
}

func (s *MapToHavanaBindIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *MapToHavanaBindIdResponseBody) GetData() []*MapToHavanaBindIdResponseBodyData {
	return s.Data
}

func (s *MapToHavanaBindIdResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *MapToHavanaBindIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MapToHavanaBindIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MapToHavanaBindIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MapToHavanaBindIdResponseBody) SetCode(v string) *MapToHavanaBindIdResponseBody {
	s.Code = &v
	return s
}

func (s *MapToHavanaBindIdResponseBody) SetData(v []*MapToHavanaBindIdResponseBodyData) *MapToHavanaBindIdResponseBody {
	s.Data = v
	return s
}

func (s *MapToHavanaBindIdResponseBody) SetHttpCode(v string) *MapToHavanaBindIdResponseBody {
	s.HttpCode = &v
	return s
}

func (s *MapToHavanaBindIdResponseBody) SetMessage(v string) *MapToHavanaBindIdResponseBody {
	s.Message = &v
	return s
}

func (s *MapToHavanaBindIdResponseBody) SetRequestId(v string) *MapToHavanaBindIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *MapToHavanaBindIdResponseBody) SetSuccess(v bool) *MapToHavanaBindIdResponseBody {
	s.Success = &v
	return s
}

func (s *MapToHavanaBindIdResponseBody) Validate() error {
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

type MapToHavanaBindIdResponseBodyData struct {
	BindHid           *string `json:"BindHid,omitempty" xml:"BindHid,omitempty"`
	HavanaBindId      *string `json:"HavanaBindId,omitempty" xml:"HavanaBindId,omitempty"`
	HavanaBindStation *string `json:"HavanaBindStation,omitempty" xml:"HavanaBindStation,omitempty"`
	Pk                *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s MapToHavanaBindIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MapToHavanaBindIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *MapToHavanaBindIdResponseBodyData) GetBindHid() *string {
	return s.BindHid
}

func (s *MapToHavanaBindIdResponseBodyData) GetHavanaBindId() *string {
	return s.HavanaBindId
}

func (s *MapToHavanaBindIdResponseBodyData) GetHavanaBindStation() *string {
	return s.HavanaBindStation
}

func (s *MapToHavanaBindIdResponseBodyData) GetPk() *string {
	return s.Pk
}

func (s *MapToHavanaBindIdResponseBodyData) SetBindHid(v string) *MapToHavanaBindIdResponseBodyData {
	s.BindHid = &v
	return s
}

func (s *MapToHavanaBindIdResponseBodyData) SetHavanaBindId(v string) *MapToHavanaBindIdResponseBodyData {
	s.HavanaBindId = &v
	return s
}

func (s *MapToHavanaBindIdResponseBodyData) SetHavanaBindStation(v string) *MapToHavanaBindIdResponseBodyData {
	s.HavanaBindStation = &v
	return s
}

func (s *MapToHavanaBindIdResponseBodyData) SetPk(v string) *MapToHavanaBindIdResponseBodyData {
	s.Pk = &v
	return s
}

func (s *MapToHavanaBindIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
