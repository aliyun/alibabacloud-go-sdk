// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMetaDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMetaDBResponse
	GetStatusCode() *int32
	SetBody(v *ListMetaDBResponseBody) *ListMetaDBResponse
	GetBody() *ListMetaDBResponseBody
}

type ListMetaDBResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMetaDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMetaDBResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMetaDBResponse) GoString() string {
	return s.String()
}

func (s *ListMetaDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMetaDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMetaDBResponse) GetBody() *ListMetaDBResponseBody {
	return s.Body
}

func (s *ListMetaDBResponse) SetHeaders(v map[string]*string) *ListMetaDBResponse {
	s.Headers = v
	return s
}

func (s *ListMetaDBResponse) SetStatusCode(v int32) *ListMetaDBResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMetaDBResponse) SetBody(v *ListMetaDBResponseBody) *ListMetaDBResponse {
	s.Body = v
	return s
}

func (s *ListMetaDBResponse) Validate() error {
	return dara.Validate(s)
}
