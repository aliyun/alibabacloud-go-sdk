// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranscodeTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTranscodeTaskStatusResponseBody
	GetRequestId() *string
}

type GetTranscodeTaskStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTranscodeTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranscodeTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTranscodeTaskStatusResponseBody) SetRequestId(v string) *GetTranscodeTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTranscodeTaskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
