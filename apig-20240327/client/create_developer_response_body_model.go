// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeveloperResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDeveloperResponseBody
	GetCode() *string
	SetData(v *CreateDeveloperResponseBodyData) *CreateDeveloperResponseBody
	GetData() *CreateDeveloperResponseBodyData
	SetMessage(v string) *CreateDeveloperResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDeveloperResponseBody
	GetRequestId() *string
}

type CreateDeveloperResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *CreateDeveloperResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateDeveloperResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDeveloperResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeveloperResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDeveloperResponseBody) GetData() *CreateDeveloperResponseBodyData {
	return s.Data
}

func (s *CreateDeveloperResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDeveloperResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDeveloperResponseBody) SetCode(v string) *CreateDeveloperResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDeveloperResponseBody) SetData(v *CreateDeveloperResponseBodyData) *CreateDeveloperResponseBody {
	s.Data = v
	return s
}

func (s *CreateDeveloperResponseBody) SetMessage(v string) *CreateDeveloperResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDeveloperResponseBody) SetRequestId(v string) *CreateDeveloperResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeveloperResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDeveloperResponseBodyData struct {
	// The developer ID. You can use this ID to locate the developer in API operations such as SetHiMarketDeveloperStatus and DeleteHiMarketDeveloper.
	//
	// example:
	//
	// dev-xxxxxxxx
	DeveloperId *string `json:"developerId,omitempty" xml:"developerId,omitempty"`
}

func (s CreateDeveloperResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDeveloperResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDeveloperResponseBodyData) GetDeveloperId() *string {
	return s.DeveloperId
}

func (s *CreateDeveloperResponseBodyData) SetDeveloperId(v string) *CreateDeveloperResponseBodyData {
	s.DeveloperId = &v
	return s
}

func (s *CreateDeveloperResponseBodyData) Validate() error {
	return dara.Validate(s)
}
