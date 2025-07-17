// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteIndexResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteIndexResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteIndexResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteIndexResponseBody
	GetStatus() *string
	SetSuccess(v bool) *DeleteIndexResponseBody
	GetSuccess() *bool
}

type DeleteIndexResponseBody struct {
	// HTTP status code
	//
	// example:
	//
	// Index.InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indications whether the API call is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndexResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIndexResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteIndexResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIndexResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteIndexResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteIndexResponseBody) SetCode(v string) *DeleteIndexResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteIndexResponseBody) SetMessage(v string) *DeleteIndexResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteIndexResponseBody) SetRequestId(v string) *DeleteIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIndexResponseBody) SetStatus(v string) *DeleteIndexResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteIndexResponseBody) SetSuccess(v bool) *DeleteIndexResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
