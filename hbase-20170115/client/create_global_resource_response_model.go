// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGlobalResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGlobalResourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateGlobalResourceResponseBody) *CreateGlobalResourceResponse
	GetBody() *CreateGlobalResourceResponseBody
}

type CreateGlobalResourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGlobalResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGlobalResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateGlobalResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGlobalResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGlobalResourceResponse) GetBody() *CreateGlobalResourceResponseBody {
	return s.Body
}

func (s *CreateGlobalResourceResponse) SetHeaders(v map[string]*string) *CreateGlobalResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateGlobalResourceResponse) SetStatusCode(v int32) *CreateGlobalResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGlobalResourceResponse) SetBody(v *CreateGlobalResourceResponseBody) *CreateGlobalResourceResponse {
	s.Body = v
	return s
}

func (s *CreateGlobalResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
