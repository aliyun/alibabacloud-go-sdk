// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEntitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEntitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEntitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListEntitiesResponseBody) *ListEntitiesResponse
	GetBody() *ListEntitiesResponseBody
}

type ListEntitiesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEntitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListEntitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEntitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEntitiesResponse) GetBody() *ListEntitiesResponseBody {
	return s.Body
}

func (s *ListEntitiesResponse) SetHeaders(v map[string]*string) *ListEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListEntitiesResponse) SetStatusCode(v int32) *ListEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEntitiesResponse) SetBody(v *ListEntitiesResponseBody) *ListEntitiesResponse {
	s.Body = v
	return s
}

func (s *ListEntitiesResponse) Validate() error {
	return dara.Validate(s)
}
