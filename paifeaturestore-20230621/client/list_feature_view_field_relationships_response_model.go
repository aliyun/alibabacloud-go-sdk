// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureViewFieldRelationshipsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFeatureViewFieldRelationshipsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFeatureViewFieldRelationshipsResponse
	GetStatusCode() *int32
	SetBody(v *ListFeatureViewFieldRelationshipsResponseBody) *ListFeatureViewFieldRelationshipsResponse
	GetBody() *ListFeatureViewFieldRelationshipsResponseBody
}

type ListFeatureViewFieldRelationshipsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeatureViewFieldRelationshipsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeatureViewFieldRelationshipsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureViewFieldRelationshipsResponse) GoString() string {
	return s.String()
}

func (s *ListFeatureViewFieldRelationshipsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFeatureViewFieldRelationshipsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFeatureViewFieldRelationshipsResponse) GetBody() *ListFeatureViewFieldRelationshipsResponseBody {
	return s.Body
}

func (s *ListFeatureViewFieldRelationshipsResponse) SetHeaders(v map[string]*string) *ListFeatureViewFieldRelationshipsResponse {
	s.Headers = v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponse) SetStatusCode(v int32) *ListFeatureViewFieldRelationshipsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponse) SetBody(v *ListFeatureViewFieldRelationshipsResponseBody) *ListFeatureViewFieldRelationshipsResponse {
	s.Body = v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
