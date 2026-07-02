// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetailByOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetDetailByOrderResponseBody
	GetCode() *int64
	SetData(v *GetDetailByOrderResponseBodyData) *GetDetailByOrderResponseBody
	GetData() *GetDetailByOrderResponseBodyData
	SetMessage(v string) *GetDetailByOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDetailByOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDetailByOrderResponseBody
	GetSuccess() *bool
}

type GetDetailByOrderResponseBody struct {
	Code      *int64                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetDetailByOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDetailByOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDetailByOrderResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetailByOrderResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetDetailByOrderResponseBody) GetData() *GetDetailByOrderResponseBodyData {
	return s.Data
}

func (s *GetDetailByOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDetailByOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDetailByOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDetailByOrderResponseBody) SetCode(v int64) *GetDetailByOrderResponseBody {
	s.Code = &v
	return s
}

func (s *GetDetailByOrderResponseBody) SetData(v *GetDetailByOrderResponseBodyData) *GetDetailByOrderResponseBody {
	s.Data = v
	return s
}

func (s *GetDetailByOrderResponseBody) SetMessage(v string) *GetDetailByOrderResponseBody {
	s.Message = &v
	return s
}

func (s *GetDetailByOrderResponseBody) SetRequestId(v string) *GetDetailByOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDetailByOrderResponseBody) SetSuccess(v bool) *GetDetailByOrderResponseBody {
	s.Success = &v
	return s
}

func (s *GetDetailByOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDetailByOrderResponseBodyData struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetDetailByOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDetailByOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDetailByOrderResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetDetailByOrderResponseBodyData) SetUrl(v string) *GetDetailByOrderResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetDetailByOrderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
