// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMaterializedViewRecommendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunMaterializedViewRecommendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunMaterializedViewRecommendResponse
	GetStatusCode() *int32
	SetBody(v *RunMaterializedViewRecommendResponseBody) *RunMaterializedViewRecommendResponse
	GetBody() *RunMaterializedViewRecommendResponseBody
}

type RunMaterializedViewRecommendResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunMaterializedViewRecommendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunMaterializedViewRecommendResponse) String() string {
	return dara.Prettify(s)
}

func (s RunMaterializedViewRecommendResponse) GoString() string {
	return s.String()
}

func (s *RunMaterializedViewRecommendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunMaterializedViewRecommendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunMaterializedViewRecommendResponse) GetBody() *RunMaterializedViewRecommendResponseBody {
	return s.Body
}

func (s *RunMaterializedViewRecommendResponse) SetHeaders(v map[string]*string) *RunMaterializedViewRecommendResponse {
	s.Headers = v
	return s
}

func (s *RunMaterializedViewRecommendResponse) SetStatusCode(v int32) *RunMaterializedViewRecommendResponse {
	s.StatusCode = &v
	return s
}

func (s *RunMaterializedViewRecommendResponse) SetBody(v *RunMaterializedViewRecommendResponseBody) *RunMaterializedViewRecommendResponse {
	s.Body = v
	return s
}

func (s *RunMaterializedViewRecommendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
