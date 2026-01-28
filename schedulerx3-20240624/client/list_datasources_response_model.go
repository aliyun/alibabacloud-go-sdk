// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDatasourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDatasourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListDatasourcesResponseBody) *ListDatasourcesResponse
	GetBody() *ListDatasourcesResponseBody
}

type ListDatasourcesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatasourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatasourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourcesResponse) GoString() string {
	return s.String()
}

func (s *ListDatasourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDatasourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDatasourcesResponse) GetBody() *ListDatasourcesResponseBody {
	return s.Body
}

func (s *ListDatasourcesResponse) SetHeaders(v map[string]*string) *ListDatasourcesResponse {
	s.Headers = v
	return s
}

func (s *ListDatasourcesResponse) SetStatusCode(v int32) *ListDatasourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatasourcesResponse) SetBody(v *ListDatasourcesResponseBody) *ListDatasourcesResponse {
	s.Body = v
	return s
}

func (s *ListDatasourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
