// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRetcodeAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRetcodeAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRetcodeAppResponse
	GetStatusCode() *int32
	SetBody(v *CreateRetcodeAppResponseBody) *CreateRetcodeAppResponse
	GetBody() *CreateRetcodeAppResponseBody
}

type CreateRetcodeAppResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRetcodeAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRetcodeAppResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRetcodeAppResponse) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRetcodeAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRetcodeAppResponse) GetBody() *CreateRetcodeAppResponseBody {
	return s.Body
}

func (s *CreateRetcodeAppResponse) SetHeaders(v map[string]*string) *CreateRetcodeAppResponse {
	s.Headers = v
	return s
}

func (s *CreateRetcodeAppResponse) SetStatusCode(v int32) *CreateRetcodeAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRetcodeAppResponse) SetBody(v *CreateRetcodeAppResponseBody) *CreateRetcodeAppResponse {
	s.Body = v
	return s
}

func (s *CreateRetcodeAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
