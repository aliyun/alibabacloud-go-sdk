// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaterializedViewRecommendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMaterializedViewRecommendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMaterializedViewRecommendResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMaterializedViewRecommendResponseBody) *DeleteMaterializedViewRecommendResponse
	GetBody() *DeleteMaterializedViewRecommendResponseBody
}

type DeleteMaterializedViewRecommendResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMaterializedViewRecommendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMaterializedViewRecommendResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaterializedViewRecommendResponse) GoString() string {
	return s.String()
}

func (s *DeleteMaterializedViewRecommendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMaterializedViewRecommendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMaterializedViewRecommendResponse) GetBody() *DeleteMaterializedViewRecommendResponseBody {
	return s.Body
}

func (s *DeleteMaterializedViewRecommendResponse) SetHeaders(v map[string]*string) *DeleteMaterializedViewRecommendResponse {
	s.Headers = v
	return s
}

func (s *DeleteMaterializedViewRecommendResponse) SetStatusCode(v int32) *DeleteMaterializedViewRecommendResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMaterializedViewRecommendResponse) SetBody(v *DeleteMaterializedViewRecommendResponseBody) *DeleteMaterializedViewRecommendResponse {
	s.Body = v
	return s
}

func (s *DeleteMaterializedViewRecommendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
