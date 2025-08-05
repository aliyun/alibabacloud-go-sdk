// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiDestinationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateApiDestinationResponseBody
	GetCode() *string
	SetDate(v *CreateApiDestinationResponseBodyDate) *CreateApiDestinationResponseBody
	GetDate() *CreateApiDestinationResponseBodyDate
	SetMessage(v string) *CreateApiDestinationResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateApiDestinationResponseBody
	GetRequestId() *string
}

type CreateApiDestinationResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned if the API destination is created.
	Date *CreateApiDestinationResponseBodyDate `json:"Date,omitempty" xml:"Date,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5DAF96FB-A4B6-548C-B999-0BFDCB2261B9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApiDestinationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApiDestinationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiDestinationResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateApiDestinationResponseBody) GetDate() *CreateApiDestinationResponseBodyDate {
	return s.Date
}

func (s *CreateApiDestinationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateApiDestinationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApiDestinationResponseBody) SetCode(v string) *CreateApiDestinationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApiDestinationResponseBody) SetDate(v *CreateApiDestinationResponseBodyDate) *CreateApiDestinationResponseBody {
	s.Date = v
	return s
}

func (s *CreateApiDestinationResponseBody) SetMessage(v string) *CreateApiDestinationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApiDestinationResponseBody) SetRequestId(v string) *CreateApiDestinationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApiDestinationResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateApiDestinationResponseBodyDate struct {
	// The name of the API destination.
	//
	// example:
	//
	// ApiDestinationName
	ApiDestinationName *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
}

func (s CreateApiDestinationResponseBodyDate) String() string {
	return dara.Prettify(s)
}

func (s CreateApiDestinationResponseBodyDate) GoString() string {
	return s.String()
}

func (s *CreateApiDestinationResponseBodyDate) GetApiDestinationName() *string {
	return s.ApiDestinationName
}

func (s *CreateApiDestinationResponseBodyDate) SetApiDestinationName(v string) *CreateApiDestinationResponseBodyDate {
	s.ApiDestinationName = &v
	return s
}

func (s *CreateApiDestinationResponseBodyDate) Validate() error {
	return dara.Validate(s)
}
