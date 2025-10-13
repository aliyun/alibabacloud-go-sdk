// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEntityStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEntityStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEntityStoreResponse
	GetStatusCode() *int32
	SetBody(v *GetEntityStoreResponseBody) *GetEntityStoreResponse
	GetBody() *GetEntityStoreResponseBody
}

type GetEntityStoreResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEntityStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEntityStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEntityStoreResponse) GoString() string {
	return s.String()
}

func (s *GetEntityStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEntityStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEntityStoreResponse) GetBody() *GetEntityStoreResponseBody {
	return s.Body
}

func (s *GetEntityStoreResponse) SetHeaders(v map[string]*string) *GetEntityStoreResponse {
	s.Headers = v
	return s
}

func (s *GetEntityStoreResponse) SetStatusCode(v int32) *GetEntityStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEntityStoreResponse) SetBody(v *GetEntityStoreResponseBody) *GetEntityStoreResponse {
	s.Body = v
	return s
}

func (s *GetEntityStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
