// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdleInstanceCullerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIdleInstanceCullerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIdleInstanceCullerResponse
	GetStatusCode() *int32
	SetBody(v *CreateIdleInstanceCullerResponseBody) *CreateIdleInstanceCullerResponse
	GetBody() *CreateIdleInstanceCullerResponseBody
}

type CreateIdleInstanceCullerResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIdleInstanceCullerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIdleInstanceCullerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIdleInstanceCullerResponse) GoString() string {
	return s.String()
}

func (s *CreateIdleInstanceCullerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIdleInstanceCullerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIdleInstanceCullerResponse) GetBody() *CreateIdleInstanceCullerResponseBody {
	return s.Body
}

func (s *CreateIdleInstanceCullerResponse) SetHeaders(v map[string]*string) *CreateIdleInstanceCullerResponse {
	s.Headers = v
	return s
}

func (s *CreateIdleInstanceCullerResponse) SetStatusCode(v int32) *CreateIdleInstanceCullerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIdleInstanceCullerResponse) SetBody(v *CreateIdleInstanceCullerResponseBody) *CreateIdleInstanceCullerResponse {
	s.Body = v
	return s
}

func (s *CreateIdleInstanceCullerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
