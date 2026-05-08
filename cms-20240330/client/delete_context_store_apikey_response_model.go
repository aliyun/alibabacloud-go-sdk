// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextStoreAPIKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContextStoreAPIKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContextStoreAPIKeyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContextStoreAPIKeyResponseBody) *DeleteContextStoreAPIKeyResponse
	GetBody() *DeleteContextStoreAPIKeyResponseBody
}

type DeleteContextStoreAPIKeyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContextStoreAPIKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContextStoreAPIKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextStoreAPIKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteContextStoreAPIKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContextStoreAPIKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContextStoreAPIKeyResponse) GetBody() *DeleteContextStoreAPIKeyResponseBody {
	return s.Body
}

func (s *DeleteContextStoreAPIKeyResponse) SetHeaders(v map[string]*string) *DeleteContextStoreAPIKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteContextStoreAPIKeyResponse) SetStatusCode(v int32) *DeleteContextStoreAPIKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContextStoreAPIKeyResponse) SetBody(v *DeleteContextStoreAPIKeyResponseBody) *DeleteContextStoreAPIKeyResponse {
	s.Body = v
	return s
}

func (s *DeleteContextStoreAPIKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
