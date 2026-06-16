// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextStoreAPIKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContextStoreAPIKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContextStoreAPIKeyResponse
	GetStatusCode() *int32
	SetBody(v *GetContextStoreAPIKeyResponseBody) *GetContextStoreAPIKeyResponse
	GetBody() *GetContextStoreAPIKeyResponseBody
}

type GetContextStoreAPIKeyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContextStoreAPIKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContextStoreAPIKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreAPIKeyResponse) GoString() string {
	return s.String()
}

func (s *GetContextStoreAPIKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContextStoreAPIKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContextStoreAPIKeyResponse) GetBody() *GetContextStoreAPIKeyResponseBody {
	return s.Body
}

func (s *GetContextStoreAPIKeyResponse) SetHeaders(v map[string]*string) *GetContextStoreAPIKeyResponse {
	s.Headers = v
	return s
}

func (s *GetContextStoreAPIKeyResponse) SetStatusCode(v int32) *GetContextStoreAPIKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContextStoreAPIKeyResponse) SetBody(v *GetContextStoreAPIKeyResponseBody) *GetContextStoreAPIKeyResponse {
	s.Body = v
	return s
}

func (s *GetContextStoreAPIKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
