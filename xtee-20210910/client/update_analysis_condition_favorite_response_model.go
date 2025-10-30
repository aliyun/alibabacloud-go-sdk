// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAnalysisConditionFavoriteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAnalysisConditionFavoriteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAnalysisConditionFavoriteResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAnalysisConditionFavoriteResponseBody) *UpdateAnalysisConditionFavoriteResponse
	GetBody() *UpdateAnalysisConditionFavoriteResponseBody
}

type UpdateAnalysisConditionFavoriteResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAnalysisConditionFavoriteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAnalysisConditionFavoriteResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAnalysisConditionFavoriteResponse) GoString() string {
	return s.String()
}

func (s *UpdateAnalysisConditionFavoriteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAnalysisConditionFavoriteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAnalysisConditionFavoriteResponse) GetBody() *UpdateAnalysisConditionFavoriteResponseBody {
	return s.Body
}

func (s *UpdateAnalysisConditionFavoriteResponse) SetHeaders(v map[string]*string) *UpdateAnalysisConditionFavoriteResponse {
	s.Headers = v
	return s
}

func (s *UpdateAnalysisConditionFavoriteResponse) SetStatusCode(v int32) *UpdateAnalysisConditionFavoriteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAnalysisConditionFavoriteResponse) SetBody(v *UpdateAnalysisConditionFavoriteResponseBody) *UpdateAnalysisConditionFavoriteResponse {
	s.Body = v
	return s
}

func (s *UpdateAnalysisConditionFavoriteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
