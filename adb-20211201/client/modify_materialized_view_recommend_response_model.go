// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaterializedViewRecommendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMaterializedViewRecommendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMaterializedViewRecommendResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMaterializedViewRecommendResponseBody) *ModifyMaterializedViewRecommendResponse
	GetBody() *ModifyMaterializedViewRecommendResponseBody
}

type ModifyMaterializedViewRecommendResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMaterializedViewRecommendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMaterializedViewRecommendResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaterializedViewRecommendResponse) GoString() string {
	return s.String()
}

func (s *ModifyMaterializedViewRecommendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMaterializedViewRecommendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMaterializedViewRecommendResponse) GetBody() *ModifyMaterializedViewRecommendResponseBody {
	return s.Body
}

func (s *ModifyMaterializedViewRecommendResponse) SetHeaders(v map[string]*string) *ModifyMaterializedViewRecommendResponse {
	s.Headers = v
	return s
}

func (s *ModifyMaterializedViewRecommendResponse) SetStatusCode(v int32) *ModifyMaterializedViewRecommendResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMaterializedViewRecommendResponse) SetBody(v *ModifyMaterializedViewRecommendResponseBody) *ModifyMaterializedViewRecommendResponse {
	s.Body = v
	return s
}

func (s *ModifyMaterializedViewRecommendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
