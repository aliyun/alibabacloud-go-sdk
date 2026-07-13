// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelProviderResponse
	GetStatusCode() *int32
	SetBody(v *GetModelProviderResponseBody) *GetModelProviderResponse
	GetBody() *GetModelProviderResponseBody
}

type GetModelProviderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelProviderResponse) GoString() string {
	return s.String()
}

func (s *GetModelProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelProviderResponse) GetBody() *GetModelProviderResponseBody {
	return s.Body
}

func (s *GetModelProviderResponse) SetHeaders(v map[string]*string) *GetModelProviderResponse {
	s.Headers = v
	return s
}

func (s *GetModelProviderResponse) SetStatusCode(v int32) *GetModelProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelProviderResponse) SetBody(v *GetModelProviderResponseBody) *GetModelProviderResponse {
	s.Body = v
	return s
}

func (s *GetModelProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
