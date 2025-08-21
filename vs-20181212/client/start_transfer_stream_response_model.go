// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTransferStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartTransferStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartTransferStreamResponse
	GetStatusCode() *int32
	SetBody(v *StartTransferStreamResponseBody) *StartTransferStreamResponse
	GetBody() *StartTransferStreamResponseBody
}

type StartTransferStreamResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartTransferStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartTransferStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s StartTransferStreamResponse) GoString() string {
	return s.String()
}

func (s *StartTransferStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartTransferStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartTransferStreamResponse) GetBody() *StartTransferStreamResponseBody {
	return s.Body
}

func (s *StartTransferStreamResponse) SetHeaders(v map[string]*string) *StartTransferStreamResponse {
	s.Headers = v
	return s
}

func (s *StartTransferStreamResponse) SetStatusCode(v int32) *StartTransferStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTransferStreamResponse) SetBody(v *StartTransferStreamResponseBody) *StartTransferStreamResponse {
	s.Body = v
	return s
}

func (s *StartTransferStreamResponse) Validate() error {
	return dara.Validate(s)
}
