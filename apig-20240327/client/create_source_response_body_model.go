// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSourceResponseBody
	GetCode() *string
	SetData(v *CreateSourceResponseBodyData) *CreateSourceResponseBody
	GetData() *CreateSourceResponseBodyData
	SetMessage(v string) *CreateSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSourceResponseBody
	GetRequestId() *string
}

type CreateSourceResponseBody struct {
	// example:
	//
	// Ok
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateSourceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSourceResponseBody) GetData() *CreateSourceResponseBodyData {
	return s.Data
}

func (s *CreateSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSourceResponseBody) SetCode(v string) *CreateSourceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSourceResponseBody) SetData(v *CreateSourceResponseBodyData) *CreateSourceResponseBody {
	s.Data = v
	return s
}

func (s *CreateSourceResponseBody) SetMessage(v string) *CreateSourceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSourceResponseBody) SetRequestId(v string) *CreateSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSourceResponseBodyData struct {
	// example:
	//
	// src-crdddallhtgt***
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s CreateSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSourceResponseBodyData) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateSourceResponseBodyData) SetSourceId(v string) *CreateSourceResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *CreateSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
