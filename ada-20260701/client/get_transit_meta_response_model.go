// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTransitMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTransitMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTransitMetaResponse
	GetStatusCode() *int32
	SetBody(v *GetTransitMetaResponseBody) *GetTransitMetaResponse
	GetBody() *GetTransitMetaResponseBody
}

type GetTransitMetaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTransitMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTransitMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTransitMetaResponse) GoString() string {
	return s.String()
}

func (s *GetTransitMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTransitMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTransitMetaResponse) GetBody() *GetTransitMetaResponseBody {
	return s.Body
}

func (s *GetTransitMetaResponse) SetHeaders(v map[string]*string) *GetTransitMetaResponse {
	s.Headers = v
	return s
}

func (s *GetTransitMetaResponse) SetStatusCode(v int32) *GetTransitMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTransitMetaResponse) SetBody(v *GetTransitMetaResponseBody) *GetTransitMetaResponse {
	s.Body = v
	return s
}

func (s *GetTransitMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
