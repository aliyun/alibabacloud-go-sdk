// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLineageRelationshipsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLineageRelationshipsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLineageRelationshipsResponse
	GetStatusCode() *int32
	SetBody(v *ListLineageRelationshipsResponseBody) *ListLineageRelationshipsResponse
	GetBody() *ListLineageRelationshipsResponseBody
}

type ListLineageRelationshipsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLineageRelationshipsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLineageRelationshipsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLineageRelationshipsResponse) GoString() string {
	return s.String()
}

func (s *ListLineageRelationshipsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLineageRelationshipsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLineageRelationshipsResponse) GetBody() *ListLineageRelationshipsResponseBody {
	return s.Body
}

func (s *ListLineageRelationshipsResponse) SetHeaders(v map[string]*string) *ListLineageRelationshipsResponse {
	s.Headers = v
	return s
}

func (s *ListLineageRelationshipsResponse) SetStatusCode(v int32) *ListLineageRelationshipsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLineageRelationshipsResponse) SetBody(v *ListLineageRelationshipsResponseBody) *ListLineageRelationshipsResponse {
	s.Body = v
	return s
}

func (s *ListLineageRelationshipsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
