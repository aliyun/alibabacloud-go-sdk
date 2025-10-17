// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivyComputeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateLivyComputeResponseBody
	GetCode() *string
	SetData(v *CreateLivyComputeResponseBodyData) *CreateLivyComputeResponseBody
	GetData() *CreateLivyComputeResponseBodyData
	SetMessage(v string) *CreateLivyComputeResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateLivyComputeResponseBody
	GetRequestId() *string
}

type CreateLivyComputeResponseBody struct {
	// example:
	//
	// 1000000
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateLivyComputeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateLivyComputeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLivyComputeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateLivyComputeResponseBody) GetData() *CreateLivyComputeResponseBodyData {
	return s.Data
}

func (s *CreateLivyComputeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateLivyComputeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLivyComputeResponseBody) SetCode(v string) *CreateLivyComputeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLivyComputeResponseBody) SetData(v *CreateLivyComputeResponseBodyData) *CreateLivyComputeResponseBody {
	s.Data = v
	return s
}

func (s *CreateLivyComputeResponseBody) SetMessage(v string) *CreateLivyComputeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLivyComputeResponseBody) SetRequestId(v string) *CreateLivyComputeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLivyComputeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLivyComputeResponseBodyData struct {
	// example:
	//
	// lc-i8xogcdfa4fk3yn1
	LivyComputeId *string `json:"livyComputeId,omitempty" xml:"livyComputeId,omitempty"`
}

func (s CreateLivyComputeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateLivyComputeResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeResponseBodyData) GetLivyComputeId() *string {
	return s.LivyComputeId
}

func (s *CreateLivyComputeResponseBodyData) SetLivyComputeId(v string) *CreateLivyComputeResponseBodyData {
	s.LivyComputeId = &v
	return s
}

func (s *CreateLivyComputeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
