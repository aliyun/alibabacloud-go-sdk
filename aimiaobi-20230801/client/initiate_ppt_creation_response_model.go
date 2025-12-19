// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiatePptCreationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitiatePptCreationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitiatePptCreationResponse
	GetStatusCode() *int32
	SetBody(v *InitiatePptCreationResponseBody) *InitiatePptCreationResponse
	GetBody() *InitiatePptCreationResponseBody
}

type InitiatePptCreationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitiatePptCreationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitiatePptCreationResponse) String() string {
	return dara.Prettify(s)
}

func (s InitiatePptCreationResponse) GoString() string {
	return s.String()
}

func (s *InitiatePptCreationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitiatePptCreationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitiatePptCreationResponse) GetBody() *InitiatePptCreationResponseBody {
	return s.Body
}

func (s *InitiatePptCreationResponse) SetHeaders(v map[string]*string) *InitiatePptCreationResponse {
	s.Headers = v
	return s
}

func (s *InitiatePptCreationResponse) SetStatusCode(v int32) *InitiatePptCreationResponse {
	s.StatusCode = &v
	return s
}

func (s *InitiatePptCreationResponse) SetBody(v *InitiatePptCreationResponseBody) *InitiatePptCreationResponse {
	s.Body = v
	return s
}

func (s *InitiatePptCreationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
