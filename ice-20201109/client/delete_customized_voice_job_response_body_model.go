// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomizedVoiceJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomizedVoiceJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCustomizedVoiceJobResponseBody
	GetSuccess() *bool
}

type DeleteCustomizedVoiceJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
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

func (s DeleteCustomizedVoiceJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomizedVoiceJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomizedVoiceJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomizedVoiceJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomizedVoiceJobResponseBody) SetRequestId(v string) *DeleteCustomizedVoiceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomizedVoiceJobResponseBody) SetSuccess(v bool) *DeleteCustomizedVoiceJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCustomizedVoiceJobResponseBody) Validate() error {
	return dara.Validate(s)
}
