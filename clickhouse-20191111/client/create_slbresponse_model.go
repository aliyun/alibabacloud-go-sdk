// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSLBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSLBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSLBResponse
	GetStatusCode() *int32
	SetBody(v *CreateSLBResponseBody) *CreateSLBResponse
	GetBody() *CreateSLBResponseBody
}

type CreateSLBResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSLBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSLBResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSLBResponse) GoString() string {
	return s.String()
}

func (s *CreateSLBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSLBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSLBResponse) GetBody() *CreateSLBResponseBody {
	return s.Body
}

func (s *CreateSLBResponse) SetHeaders(v map[string]*string) *CreateSLBResponse {
	s.Headers = v
	return s
}

func (s *CreateSLBResponse) SetStatusCode(v int32) *CreateSLBResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSLBResponse) SetBody(v *CreateSLBResponseBody) *CreateSLBResponse {
	s.Body = v
	return s
}

func (s *CreateSLBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
