// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDeletionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelDeletionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelDeletionResponse
	GetStatusCode() *int32
	SetBody(v *CancelDeletionResponseBody) *CancelDeletionResponse
	GetBody() *CancelDeletionResponseBody
}

type CancelDeletionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelDeletionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelDeletionResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelDeletionResponse) GoString() string {
	return s.String()
}

func (s *CancelDeletionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelDeletionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelDeletionResponse) GetBody() *CancelDeletionResponseBody {
	return s.Body
}

func (s *CancelDeletionResponse) SetHeaders(v map[string]*string) *CancelDeletionResponse {
	s.Headers = v
	return s
}

func (s *CancelDeletionResponse) SetStatusCode(v int32) *CancelDeletionResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelDeletionResponse) SetBody(v *CancelDeletionResponseBody) *CancelDeletionResponse {
	s.Body = v
	return s
}

func (s *CancelDeletionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
