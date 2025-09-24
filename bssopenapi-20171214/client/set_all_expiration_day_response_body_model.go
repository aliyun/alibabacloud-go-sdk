// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAllExpirationDayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetAllExpirationDayResponseBody
	GetCode() *string
	SetMessage(v string) *SetAllExpirationDayResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetAllExpirationDayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetAllExpirationDayResponseBody
	GetSuccess() *bool
}

type SetAllExpirationDayResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// The message that is returned
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// The ID of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetAllExpirationDayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAllExpirationDayResponseBody) GoString() string {
	return s.String()
}

func (s *SetAllExpirationDayResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetAllExpirationDayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetAllExpirationDayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAllExpirationDayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetAllExpirationDayResponseBody) SetCode(v string) *SetAllExpirationDayResponseBody {
	s.Code = &v
	return s
}

func (s *SetAllExpirationDayResponseBody) SetMessage(v string) *SetAllExpirationDayResponseBody {
	s.Message = &v
	return s
}

func (s *SetAllExpirationDayResponseBody) SetRequestId(v string) *SetAllExpirationDayResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAllExpirationDayResponseBody) SetSuccess(v bool) *SetAllExpirationDayResponseBody {
	s.Success = &v
	return s
}

func (s *SetAllExpirationDayResponseBody) Validate() error {
	return dara.Validate(s)
}
