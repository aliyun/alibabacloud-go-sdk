// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWarningConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteWarningConfigResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteWarningConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteWarningConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteWarningConfigResponseBody
	GetSuccess() *bool
}

type DeleteWarningConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4D55C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWarningConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWarningConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWarningConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteWarningConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteWarningConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWarningConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteWarningConfigResponseBody) SetCode(v string) *DeleteWarningConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWarningConfigResponseBody) SetMessage(v string) *DeleteWarningConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWarningConfigResponseBody) SetRequestId(v string) *DeleteWarningConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWarningConfigResponseBody) SetSuccess(v bool) *DeleteWarningConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteWarningConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
