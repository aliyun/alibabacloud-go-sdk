// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelCatalogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelCatalogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelCatalogResponse
	GetStatusCode() *int32
	SetBody(v *ListModelCatalogResponseBody) *ListModelCatalogResponse
	GetBody() *ListModelCatalogResponseBody
}

type ListModelCatalogResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelCatalogResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelCatalogResponse) GoString() string {
	return s.String()
}

func (s *ListModelCatalogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelCatalogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelCatalogResponse) GetBody() *ListModelCatalogResponseBody {
	return s.Body
}

func (s *ListModelCatalogResponse) SetHeaders(v map[string]*string) *ListModelCatalogResponse {
	s.Headers = v
	return s
}

func (s *ListModelCatalogResponse) SetStatusCode(v int32) *ListModelCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelCatalogResponse) SetBody(v *ListModelCatalogResponseBody) *ListModelCatalogResponse {
	s.Body = v
	return s
}

func (s *ListModelCatalogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
