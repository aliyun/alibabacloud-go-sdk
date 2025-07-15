// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAISubtitleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveAISubtitleResponseBody
	GetRequestId() *string
}

type UpdateLiveAISubtitleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// String	5c6a2a0df228-4a64- af62-20e91b96****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveAISubtitleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAISubtitleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveAISubtitleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveAISubtitleResponseBody) SetRequestId(v string) *UpdateLiveAISubtitleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveAISubtitleResponseBody) Validate() error {
	return dara.Validate(s)
}
