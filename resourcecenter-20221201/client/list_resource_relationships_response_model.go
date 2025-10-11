// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceRelationshipsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceRelationshipsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceRelationshipsResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceRelationshipsResponseBody) *ListResourceRelationshipsResponse
	GetBody() *ListResourceRelationshipsResponseBody
}

type ListResourceRelationshipsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceRelationshipsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceRelationshipsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceRelationshipsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceRelationshipsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceRelationshipsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceRelationshipsResponse) GetBody() *ListResourceRelationshipsResponseBody {
	return s.Body
}

func (s *ListResourceRelationshipsResponse) SetHeaders(v map[string]*string) *ListResourceRelationshipsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceRelationshipsResponse) SetStatusCode(v int32) *ListResourceRelationshipsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceRelationshipsResponse) SetBody(v *ListResourceRelationshipsResponseBody) *ListResourceRelationshipsResponse {
	s.Body = v
	return s
}

func (s *ListResourceRelationshipsResponse) Validate() error {
	return dara.Validate(s)
}
