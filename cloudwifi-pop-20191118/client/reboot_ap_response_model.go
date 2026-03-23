// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootApResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebootApResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebootApResponse
	GetStatusCode() *int32
	SetBody(v *RebootApResponseBody) *RebootApResponse
	GetBody() *RebootApResponseBody
}

type RebootApResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootApResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootApResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootApResponse) GoString() string {
	return s.String()
}

func (s *RebootApResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebootApResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebootApResponse) GetBody() *RebootApResponseBody {
	return s.Body
}

func (s *RebootApResponse) SetHeaders(v map[string]*string) *RebootApResponse {
	s.Headers = v
	return s
}

func (s *RebootApResponse) SetStatusCode(v int32) *RebootApResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootApResponse) SetBody(v *RebootApResponseBody) *RebootApResponse {
	s.Body = v
	return s
}

func (s *RebootApResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
