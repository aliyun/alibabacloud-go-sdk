// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLoggingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartLoggingResponseBody
	GetRequestId() *string
}

type StartLoggingResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 145318BE-DEE1-4C57-AA7C-5BE7D34A6AE0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartLoggingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartLoggingResponseBody) GoString() string {
	return s.String()
}

func (s *StartLoggingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartLoggingResponseBody) SetRequestId(v string) *StartLoggingResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartLoggingResponseBody) Validate() error {
	return dara.Validate(s)
}
