// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeLiveStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResumeLiveStreamResponseBody
	GetRequestId() *string
}

type ResumeLiveStreamResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16BFE188-B193-4C3C-ADC5-79A7E31486EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeLiveStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeLiveStreamResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeLiveStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeLiveStreamResponseBody) SetRequestId(v string) *ResumeLiveStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeLiveStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
