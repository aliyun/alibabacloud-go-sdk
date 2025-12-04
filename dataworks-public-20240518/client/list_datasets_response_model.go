// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDatasetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDatasetsResponse
	GetStatusCode() *int32
	SetBody(v *ListDatasetsResponseBody) *ListDatasetsResponse
	GetBody() *ListDatasetsResponseBody
}

type ListDatasetsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatasetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatasetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponse) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDatasetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDatasetsResponse) GetBody() *ListDatasetsResponseBody {
	return s.Body
}

func (s *ListDatasetsResponse) SetHeaders(v map[string]*string) *ListDatasetsResponse {
	s.Headers = v
	return s
}

func (s *ListDatasetsResponse) SetStatusCode(v int32) *ListDatasetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatasetsResponse) SetBody(v *ListDatasetsResponseBody) *ListDatasetsResponse {
	s.Body = v
	return s
}

func (s *ListDatasetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
