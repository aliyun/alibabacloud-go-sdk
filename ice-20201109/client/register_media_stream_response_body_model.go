// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterMediaStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *RegisterMediaStreamResponseBody
	GetMediaId() *string
	SetRequestId(v string) *RegisterMediaStreamResponseBody
	GetRequestId() *string
}

type RegisterMediaStreamResponseBody struct {
	// The ID of the media asset.
	//
	// example:
	//
	// 5e778ec0027b71ed80a8909598506302
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterMediaStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterMediaStreamResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterMediaStreamResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *RegisterMediaStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterMediaStreamResponseBody) SetMediaId(v string) *RegisterMediaStreamResponseBody {
	s.MediaId = &v
	return s
}

func (s *RegisterMediaStreamResponseBody) SetRequestId(v string) *RegisterMediaStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterMediaStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
