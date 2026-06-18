// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotlineNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteHotlineNumberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *DeleteHotlineNumberResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *DeleteHotlineNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHotlineNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteHotlineNumberResponseBody
	GetSuccess() *bool
}

type DeleteHotlineNumberResponseBody struct {
	// The status code. A value of Success indicates that the request succeeded.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 98B032F5-6473-4EAC-8BA8-C28993513A1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteHotlineNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotlineNumberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHotlineNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteHotlineNumberResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DeleteHotlineNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHotlineNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHotlineNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteHotlineNumberResponseBody) SetCode(v string) *DeleteHotlineNumberResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHotlineNumberResponseBody) SetHttpStatusCode(v int64) *DeleteHotlineNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteHotlineNumberResponseBody) SetMessage(v string) *DeleteHotlineNumberResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHotlineNumberResponseBody) SetRequestId(v string) *DeleteHotlineNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHotlineNumberResponseBody) SetSuccess(v bool) *DeleteHotlineNumberResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteHotlineNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
