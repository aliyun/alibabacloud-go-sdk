// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMapFromHavanaBindIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MapFromHavanaBindIdResponseBody
	GetCode() *string
	SetData(v []*MapFromHavanaBindIdResponseBodyData) *MapFromHavanaBindIdResponseBody
	GetData() []*MapFromHavanaBindIdResponseBodyData
	SetHttpCode(v string) *MapFromHavanaBindIdResponseBody
	GetHttpCode() *string
	SetMessage(v string) *MapFromHavanaBindIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *MapFromHavanaBindIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MapFromHavanaBindIdResponseBody
	GetSuccess() *bool
}

type MapFromHavanaBindIdResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*MapFromHavanaBindIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpCode  *string                                `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MapFromHavanaBindIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MapFromHavanaBindIdResponseBody) GoString() string {
	return s.String()
}

func (s *MapFromHavanaBindIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *MapFromHavanaBindIdResponseBody) GetData() []*MapFromHavanaBindIdResponseBodyData {
	return s.Data
}

func (s *MapFromHavanaBindIdResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *MapFromHavanaBindIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MapFromHavanaBindIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MapFromHavanaBindIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MapFromHavanaBindIdResponseBody) SetCode(v string) *MapFromHavanaBindIdResponseBody {
	s.Code = &v
	return s
}

func (s *MapFromHavanaBindIdResponseBody) SetData(v []*MapFromHavanaBindIdResponseBodyData) *MapFromHavanaBindIdResponseBody {
	s.Data = v
	return s
}

func (s *MapFromHavanaBindIdResponseBody) SetHttpCode(v string) *MapFromHavanaBindIdResponseBody {
	s.HttpCode = &v
	return s
}

func (s *MapFromHavanaBindIdResponseBody) SetMessage(v string) *MapFromHavanaBindIdResponseBody {
	s.Message = &v
	return s
}

func (s *MapFromHavanaBindIdResponseBody) SetRequestId(v string) *MapFromHavanaBindIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *MapFromHavanaBindIdResponseBody) SetSuccess(v bool) *MapFromHavanaBindIdResponseBody {
	s.Success = &v
	return s
}

func (s *MapFromHavanaBindIdResponseBody) Validate() error {
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

type MapFromHavanaBindIdResponseBodyData struct {
	BindHid           *string `json:"BindHid,omitempty" xml:"BindHid,omitempty"`
	HavanaBindId      *string `json:"HavanaBindId,omitempty" xml:"HavanaBindId,omitempty"`
	HavanaBindStation *string `json:"HavanaBindStation,omitempty" xml:"HavanaBindStation,omitempty"`
	Pk                *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s MapFromHavanaBindIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MapFromHavanaBindIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *MapFromHavanaBindIdResponseBodyData) GetBindHid() *string {
	return s.BindHid
}

func (s *MapFromHavanaBindIdResponseBodyData) GetHavanaBindId() *string {
	return s.HavanaBindId
}

func (s *MapFromHavanaBindIdResponseBodyData) GetHavanaBindStation() *string {
	return s.HavanaBindStation
}

func (s *MapFromHavanaBindIdResponseBodyData) GetPk() *string {
	return s.Pk
}

func (s *MapFromHavanaBindIdResponseBodyData) SetBindHid(v string) *MapFromHavanaBindIdResponseBodyData {
	s.BindHid = &v
	return s
}

func (s *MapFromHavanaBindIdResponseBodyData) SetHavanaBindId(v string) *MapFromHavanaBindIdResponseBodyData {
	s.HavanaBindId = &v
	return s
}

func (s *MapFromHavanaBindIdResponseBodyData) SetHavanaBindStation(v string) *MapFromHavanaBindIdResponseBodyData {
	s.HavanaBindStation = &v
	return s
}

func (s *MapFromHavanaBindIdResponseBodyData) SetPk(v string) *MapFromHavanaBindIdResponseBodyData {
	s.Pk = &v
	return s
}

func (s *MapFromHavanaBindIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
