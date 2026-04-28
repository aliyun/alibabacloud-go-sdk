// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchSimilarImageClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchSimilarImageClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchSimilarImageClustersResponse
	GetStatusCode() *int32
	SetBody(v *SearchSimilarImageClustersResponseBody) *SearchSimilarImageClustersResponse
	GetBody() *SearchSimilarImageClustersResponseBody
}

type SearchSimilarImageClustersResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchSimilarImageClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchSimilarImageClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchSimilarImageClustersResponse) GoString() string {
	return s.String()
}

func (s *SearchSimilarImageClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchSimilarImageClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchSimilarImageClustersResponse) GetBody() *SearchSimilarImageClustersResponseBody {
	return s.Body
}

func (s *SearchSimilarImageClustersResponse) SetHeaders(v map[string]*string) *SearchSimilarImageClustersResponse {
	s.Headers = v
	return s
}

func (s *SearchSimilarImageClustersResponse) SetStatusCode(v int32) *SearchSimilarImageClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchSimilarImageClustersResponse) SetBody(v *SearchSimilarImageClustersResponseBody) *SearchSimilarImageClustersResponse {
	s.Body = v
	return s
}

func (s *SearchSimilarImageClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
