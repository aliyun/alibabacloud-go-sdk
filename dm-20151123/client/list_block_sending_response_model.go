// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBlockSendingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBlockSendingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBlockSendingResponse
	GetStatusCode() *int32
	SetBody(v *ListBlockSendingResponseBody) *ListBlockSendingResponse
	GetBody() *ListBlockSendingResponseBody
}

type ListBlockSendingResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBlockSendingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBlockSendingResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBlockSendingResponse) GoString() string {
	return s.String()
}

func (s *ListBlockSendingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBlockSendingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBlockSendingResponse) GetBody() *ListBlockSendingResponseBody {
	return s.Body
}

func (s *ListBlockSendingResponse) SetHeaders(v map[string]*string) *ListBlockSendingResponse {
	s.Headers = v
	return s
}

func (s *ListBlockSendingResponse) SetStatusCode(v int32) *ListBlockSendingResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBlockSendingResponse) SetBody(v *ListBlockSendingResponseBody) *ListBlockSendingResponse {
	s.Body = v
	return s
}

func (s *ListBlockSendingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
