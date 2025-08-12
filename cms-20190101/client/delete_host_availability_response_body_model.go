// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHostAvailabilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteHostAvailabilityResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteHostAvailabilityResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHostAvailabilityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteHostAvailabilityResponseBody
	GetSuccess() *bool
}

type DeleteHostAvailabilityResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// 57C782E6-B235-4842-AD2B-DB94961761EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteHostAvailabilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHostAvailabilityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHostAvailabilityResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteHostAvailabilityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHostAvailabilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHostAvailabilityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteHostAvailabilityResponseBody) SetCode(v string) *DeleteHostAvailabilityResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHostAvailabilityResponseBody) SetMessage(v string) *DeleteHostAvailabilityResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHostAvailabilityResponseBody) SetRequestId(v string) *DeleteHostAvailabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHostAvailabilityResponseBody) SetSuccess(v bool) *DeleteHostAvailabilityResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteHostAvailabilityResponseBody) Validate() error {
	return dara.Validate(s)
}
