// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigSequenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConfigSequenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConfigSequenceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConfigSequenceResponseBody) *UpdateConfigSequenceResponse
	GetBody() *UpdateConfigSequenceResponseBody
}

type UpdateConfigSequenceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConfigSequenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConfigSequenceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigSequenceResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigSequenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConfigSequenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConfigSequenceResponse) GetBody() *UpdateConfigSequenceResponseBody {
	return s.Body
}

func (s *UpdateConfigSequenceResponse) SetHeaders(v map[string]*string) *UpdateConfigSequenceResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigSequenceResponse) SetStatusCode(v int32) *UpdateConfigSequenceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConfigSequenceResponse) SetBody(v *UpdateConfigSequenceResponseBody) *UpdateConfigSequenceResponse {
	s.Body = v
	return s
}

func (s *UpdateConfigSequenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
