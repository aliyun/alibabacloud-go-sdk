// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateHttpApiResponseBody
	GetCode() *string
	SetData(v *CreateHttpApiResponseBodyData) *CreateHttpApiResponseBody
	GetData() *CreateHttpApiResponseBodyData
	SetMessage(v string) *CreateHttpApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateHttpApiResponseBody
	GetRequestId() *string
}

type CreateHttpApiResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The API information.
	Data *CreateHttpApiResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1994B10-C6A8-58FA-8347-6A08B0D4EFDE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateHttpApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHttpApiResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateHttpApiResponseBody) GetData() *CreateHttpApiResponseBodyData {
	return s.Data
}

func (s *CreateHttpApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHttpApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHttpApiResponseBody) SetCode(v string) *CreateHttpApiResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHttpApiResponseBody) SetData(v *CreateHttpApiResponseBodyData) *CreateHttpApiResponseBody {
	s.Data = v
	return s
}

func (s *CreateHttpApiResponseBody) SetMessage(v string) *CreateHttpApiResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHttpApiResponseBody) SetRequestId(v string) *CreateHttpApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHttpApiResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateHttpApiResponseBodyData struct {
	// The HTTP API ID.
	//
	// example:
	//
	// api-xxx
	HttpApiId *string `json:"httpApiId,omitempty" xml:"httpApiId,omitempty"`
	// The API name.
	//
	// example:
	//
	// test-api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateHttpApiResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateHttpApiResponseBodyData) GetHttpApiId() *string {
	return s.HttpApiId
}

func (s *CreateHttpApiResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateHttpApiResponseBodyData) SetHttpApiId(v string) *CreateHttpApiResponseBodyData {
	s.HttpApiId = &v
	return s
}

func (s *CreateHttpApiResponseBodyData) SetName(v string) *CreateHttpApiResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateHttpApiResponseBodyData) Validate() error {
	return dara.Validate(s)
}
