// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelResponse
	GetStatusCode() *int32
	SetBody(v *CancelResponseBody) *CancelResponse
	GetBody() *CancelResponseBody
}

type CancelResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelResponse) GoString() string {
	return s.String()
}

func (s *CancelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelResponse) GetBody() *CancelResponseBody {
	return s.Body
}

func (s *CancelResponse) SetHeaders(v map[string]*string) *CancelResponse {
	s.Headers = v
	return s
}

func (s *CancelResponse) SetStatusCode(v int32) *CancelResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelResponse) SetBody(v *CancelResponseBody) *CancelResponse {
	s.Body = v
	return s
}

func (s *CancelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
