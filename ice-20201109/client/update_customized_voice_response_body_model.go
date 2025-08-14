// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomizedVoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCustomizedVoiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCustomizedVoiceResponseBody
	GetSuccess() *bool
}

type UpdateCustomizedVoiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4E84BE44-58A7-****-****-FBEBEA16EF94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
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

func (s UpdateCustomizedVoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomizedVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomizedVoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomizedVoiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCustomizedVoiceResponseBody) SetRequestId(v string) *UpdateCustomizedVoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomizedVoiceResponseBody) SetSuccess(v bool) *UpdateCustomizedVoiceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCustomizedVoiceResponseBody) Validate() error {
	return dara.Validate(s)
}
