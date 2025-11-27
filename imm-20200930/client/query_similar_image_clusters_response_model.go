// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySimilarImageClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySimilarImageClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySimilarImageClustersResponse
	GetStatusCode() *int32
	SetBody(v *QuerySimilarImageClustersResponseBody) *QuerySimilarImageClustersResponse
	GetBody() *QuerySimilarImageClustersResponseBody
}

type QuerySimilarImageClustersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySimilarImageClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySimilarImageClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySimilarImageClustersResponse) GoString() string {
	return s.String()
}

func (s *QuerySimilarImageClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySimilarImageClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySimilarImageClustersResponse) GetBody() *QuerySimilarImageClustersResponseBody {
	return s.Body
}

func (s *QuerySimilarImageClustersResponse) SetHeaders(v map[string]*string) *QuerySimilarImageClustersResponse {
	s.Headers = v
	return s
}

func (s *QuerySimilarImageClustersResponse) SetStatusCode(v int32) *QuerySimilarImageClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySimilarImageClustersResponse) SetBody(v *QuerySimilarImageClustersResponseBody) *QuerySimilarImageClustersResponse {
	s.Body = v
	return s
}

func (s *QuerySimilarImageClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
