// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartTranscodeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartTranscodeTaskResponseBody
	GetRequestId() *string
}

type RestartTranscodeTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartTranscodeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartTranscodeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RestartTranscodeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartTranscodeTaskResponseBody) SetRequestId(v string) *RestartTranscodeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartTranscodeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
