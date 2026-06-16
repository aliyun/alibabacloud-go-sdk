// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextStoreAPIKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateContextStoreAPIKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateContextStoreAPIKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateContextStoreAPIKeyResponseBody) *CreateContextStoreAPIKeyResponse
	GetBody() *CreateContextStoreAPIKeyResponseBody
}

type CreateContextStoreAPIKeyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContextStoreAPIKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContextStoreAPIKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateContextStoreAPIKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateContextStoreAPIKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateContextStoreAPIKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateContextStoreAPIKeyResponse) GetBody() *CreateContextStoreAPIKeyResponseBody {
	return s.Body
}

func (s *CreateContextStoreAPIKeyResponse) SetHeaders(v map[string]*string) *CreateContextStoreAPIKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateContextStoreAPIKeyResponse) SetStatusCode(v int32) *CreateContextStoreAPIKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContextStoreAPIKeyResponse) SetBody(v *CreateContextStoreAPIKeyResponseBody) *CreateContextStoreAPIKeyResponse {
	s.Body = v
	return s
}

func (s *CreateContextStoreAPIKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
