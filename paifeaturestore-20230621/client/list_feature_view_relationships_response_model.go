// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureViewRelationshipsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFeatureViewRelationshipsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFeatureViewRelationshipsResponse
	GetStatusCode() *int32
	SetBody(v *ListFeatureViewRelationshipsResponseBody) *ListFeatureViewRelationshipsResponse
	GetBody() *ListFeatureViewRelationshipsResponseBody
}

type ListFeatureViewRelationshipsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeatureViewRelationshipsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeatureViewRelationshipsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureViewRelationshipsResponse) GoString() string {
	return s.String()
}

func (s *ListFeatureViewRelationshipsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFeatureViewRelationshipsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFeatureViewRelationshipsResponse) GetBody() *ListFeatureViewRelationshipsResponseBody {
	return s.Body
}

func (s *ListFeatureViewRelationshipsResponse) SetHeaders(v map[string]*string) *ListFeatureViewRelationshipsResponse {
	s.Headers = v
	return s
}

func (s *ListFeatureViewRelationshipsResponse) SetStatusCode(v int32) *ListFeatureViewRelationshipsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeatureViewRelationshipsResponse) SetBody(v *ListFeatureViewRelationshipsResponseBody) *ListFeatureViewRelationshipsResponse {
	s.Body = v
	return s
}

func (s *ListFeatureViewRelationshipsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
