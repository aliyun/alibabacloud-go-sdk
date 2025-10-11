// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiAccountResourceRelationshipsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMultiAccountResourceRelationshipsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMultiAccountResourceRelationshipsResponse
	GetStatusCode() *int32
	SetBody(v *ListMultiAccountResourceRelationshipsResponseBody) *ListMultiAccountResourceRelationshipsResponse
	GetBody() *ListMultiAccountResourceRelationshipsResponseBody
}

type ListMultiAccountResourceRelationshipsResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultiAccountResourceRelationshipsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultiAccountResourceRelationshipsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountResourceRelationshipsResponse) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceRelationshipsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMultiAccountResourceRelationshipsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMultiAccountResourceRelationshipsResponse) GetBody() *ListMultiAccountResourceRelationshipsResponseBody {
	return s.Body
}

func (s *ListMultiAccountResourceRelationshipsResponse) SetHeaders(v map[string]*string) *ListMultiAccountResourceRelationshipsResponse {
	s.Headers = v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponse) SetStatusCode(v int32) *ListMultiAccountResourceRelationshipsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponse) SetBody(v *ListMultiAccountResourceRelationshipsResponseBody) *ListMultiAccountResourceRelationshipsResponse {
	s.Body = v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponse) Validate() error {
	return dara.Validate(s)
}
