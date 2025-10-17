// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaterializedViewRecommendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMaterializedViewRecommendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMaterializedViewRecommendResponse
	GetStatusCode() *int32
	SetBody(v *CreateMaterializedViewRecommendResponseBody) *CreateMaterializedViewRecommendResponse
	GetBody() *CreateMaterializedViewRecommendResponseBody
}

type CreateMaterializedViewRecommendResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMaterializedViewRecommendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMaterializedViewRecommendResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMaterializedViewRecommendResponse) GoString() string {
	return s.String()
}

func (s *CreateMaterializedViewRecommendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMaterializedViewRecommendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMaterializedViewRecommendResponse) GetBody() *CreateMaterializedViewRecommendResponseBody {
	return s.Body
}

func (s *CreateMaterializedViewRecommendResponse) SetHeaders(v map[string]*string) *CreateMaterializedViewRecommendResponse {
	s.Headers = v
	return s
}

func (s *CreateMaterializedViewRecommendResponse) SetStatusCode(v int32) *CreateMaterializedViewRecommendResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMaterializedViewRecommendResponse) SetBody(v *CreateMaterializedViewRecommendResponseBody) *CreateMaterializedViewRecommendResponse {
	s.Body = v
	return s
}

func (s *CreateMaterializedViewRecommendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
