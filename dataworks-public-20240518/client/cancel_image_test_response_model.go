// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelImageTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelImageTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelImageTestResponse
	GetStatusCode() *int32
	SetBody(v *CancelImageTestResponseBody) *CancelImageTestResponse
	GetBody() *CancelImageTestResponseBody
}

type CancelImageTestResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelImageTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelImageTestResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelImageTestResponse) GoString() string {
	return s.String()
}

func (s *CancelImageTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelImageTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelImageTestResponse) GetBody() *CancelImageTestResponseBody {
	return s.Body
}

func (s *CancelImageTestResponse) SetHeaders(v map[string]*string) *CancelImageTestResponse {
	s.Headers = v
	return s
}

func (s *CancelImageTestResponse) SetStatusCode(v int32) *CancelImageTestResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelImageTestResponse) SetBody(v *CancelImageTestResponseBody) *CancelImageTestResponse {
	s.Body = v
	return s
}

func (s *CancelImageTestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
