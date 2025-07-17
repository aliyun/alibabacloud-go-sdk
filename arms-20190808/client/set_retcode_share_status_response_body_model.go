// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRetcodeShareStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *SetRetcodeShareStatusResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *SetRetcodeShareStatusResponseBody
	GetRequestId() *string
}

type SetRetcodeShareStatusResponseBody struct {
	// Indicates whether the call is successful. Valid values:
	//
	// 	- `true`: The call is successful.
	//
	// 	- `false`: The call fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 40B10E04-81E8-4643-970D-F1B38F2E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetRetcodeShareStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetRetcodeShareStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetRetcodeShareStatusResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *SetRetcodeShareStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetRetcodeShareStatusResponseBody) SetIsSuccess(v bool) *SetRetcodeShareStatusResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *SetRetcodeShareStatusResponseBody) SetRequestId(v string) *SetRetcodeShareStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetRetcodeShareStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
