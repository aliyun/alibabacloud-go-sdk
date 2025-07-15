// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAISubtitleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveAISubtitleResponseBody
	GetRequestId() *string
}

type DeleteLiveAISubtitleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0df228-4a64-af62-20e91b96****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveAISubtitleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAISubtitleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveAISubtitleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveAISubtitleResponseBody) SetRequestId(v string) *DeleteLiveAISubtitleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveAISubtitleResponseBody) Validate() error {
	return dara.Validate(s)
}
