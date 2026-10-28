// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelRunResponse
	GetStatusCode() *int32
	SetBody(v *CancelRunResponseBody) *CancelRunResponse
	GetBody() *CancelRunResponseBody
}

type CancelRunResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelRunResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelRunResponse) GoString() string {
	return s.String()
}

func (s *CancelRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelRunResponse) GetBody() *CancelRunResponseBody {
	return s.Body
}

func (s *CancelRunResponse) SetHeaders(v map[string]*string) *CancelRunResponse {
	s.Headers = v
	return s
}

func (s *CancelRunResponse) SetStatusCode(v int32) *CancelRunResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelRunResponse) SetBody(v *CancelRunResponseBody) *CancelRunResponse {
	s.Body = v
	return s
}

func (s *CancelRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
