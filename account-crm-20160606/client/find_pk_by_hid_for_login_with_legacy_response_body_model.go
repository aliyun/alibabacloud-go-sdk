// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindPkByHidForLoginWithLegacyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FindPkByHidForLoginWithLegacyResponseBody
	GetCode() *string
	SetData(v *FindPkByHidForLoginWithLegacyResponseBodyData) *FindPkByHidForLoginWithLegacyResponseBody
	GetData() *FindPkByHidForLoginWithLegacyResponseBodyData
	SetMessage(v string) *FindPkByHidForLoginWithLegacyResponseBody
	GetMessage() *string
	SetRequestId(v string) *FindPkByHidForLoginWithLegacyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FindPkByHidForLoginWithLegacyResponseBody
	GetSuccess() *bool
}

type FindPkByHidForLoginWithLegacyResponseBody struct {
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *FindPkByHidForLoginWithLegacyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FindPkByHidForLoginWithLegacyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FindPkByHidForLoginWithLegacyResponseBody) GoString() string {
	return s.String()
}

func (s *FindPkByHidForLoginWithLegacyResponseBody) GetCode() *string {
	return s.Code
}

func (s *FindPkByHidForLoginWithLegacyResponseBody) GetData() *FindPkByHidForLoginWithLegacyResponseBodyData {
	return s.Data
}

func (s *FindPkByHidForLoginWithLegacyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FindPkByHidForLoginWithLegacyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FindPkByHidForLoginWithLegacyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FindPkByHidForLoginWithLegacyResponseBody) SetCode(v string) *FindPkByHidForLoginWithLegacyResponseBody {
	s.Code = &v
	return s
}

func (s *FindPkByHidForLoginWithLegacyResponseBody) SetData(v *FindPkByHidForLoginWithLegacyResponseBodyData) *FindPkByHidForLoginWithLegacyResponseBody {
	s.Data = v
	return s
}

func (s *FindPkByHidForLoginWithLegacyResponseBody) SetMessage(v string) *FindPkByHidForLoginWithLegacyResponseBody {
	s.Message = &v
	return s
}

func (s *FindPkByHidForLoginWithLegacyResponseBody) SetRequestId(v string) *FindPkByHidForLoginWithLegacyResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindPkByHidForLoginWithLegacyResponseBody) SetSuccess(v bool) *FindPkByHidForLoginWithLegacyResponseBody {
	s.Success = &v
	return s
}

func (s *FindPkByHidForLoginWithLegacyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FindPkByHidForLoginWithLegacyResponseBodyData struct {
	Hid *string `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Pk  *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s FindPkByHidForLoginWithLegacyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FindPkByHidForLoginWithLegacyResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindPkByHidForLoginWithLegacyResponseBodyData) GetHid() *string {
	return s.Hid
}

func (s *FindPkByHidForLoginWithLegacyResponseBodyData) GetPk() *string {
	return s.Pk
}

func (s *FindPkByHidForLoginWithLegacyResponseBodyData) SetHid(v string) *FindPkByHidForLoginWithLegacyResponseBodyData {
	s.Hid = &v
	return s
}

func (s *FindPkByHidForLoginWithLegacyResponseBodyData) SetPk(v string) *FindPkByHidForLoginWithLegacyResponseBodyData {
	s.Pk = &v
	return s
}

func (s *FindPkByHidForLoginWithLegacyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
