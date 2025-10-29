// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelLogstashDeletionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelLogstashDeletionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelLogstashDeletionResponse
	GetStatusCode() *int32
	SetBody(v *CancelLogstashDeletionResponseBody) *CancelLogstashDeletionResponse
	GetBody() *CancelLogstashDeletionResponseBody
}

type CancelLogstashDeletionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelLogstashDeletionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelLogstashDeletionResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelLogstashDeletionResponse) GoString() string {
	return s.String()
}

func (s *CancelLogstashDeletionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelLogstashDeletionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelLogstashDeletionResponse) GetBody() *CancelLogstashDeletionResponseBody {
	return s.Body
}

func (s *CancelLogstashDeletionResponse) SetHeaders(v map[string]*string) *CancelLogstashDeletionResponse {
	s.Headers = v
	return s
}

func (s *CancelLogstashDeletionResponse) SetStatusCode(v int32) *CancelLogstashDeletionResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelLogstashDeletionResponse) SetBody(v *CancelLogstashDeletionResponseBody) *CancelLogstashDeletionResponse {
	s.Body = v
	return s
}

func (s *CancelLogstashDeletionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
