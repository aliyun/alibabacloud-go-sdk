// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAnalysisConditionFavoriteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAnalysisConditionFavoriteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAnalysisConditionFavoriteResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAnalysisConditionFavoriteResponseBody) *DeleteAnalysisConditionFavoriteResponse
	GetBody() *DeleteAnalysisConditionFavoriteResponseBody
}

type DeleteAnalysisConditionFavoriteResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAnalysisConditionFavoriteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAnalysisConditionFavoriteResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAnalysisConditionFavoriteResponse) GoString() string {
	return s.String()
}

func (s *DeleteAnalysisConditionFavoriteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAnalysisConditionFavoriteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAnalysisConditionFavoriteResponse) GetBody() *DeleteAnalysisConditionFavoriteResponseBody {
	return s.Body
}

func (s *DeleteAnalysisConditionFavoriteResponse) SetHeaders(v map[string]*string) *DeleteAnalysisConditionFavoriteResponse {
	s.Headers = v
	return s
}

func (s *DeleteAnalysisConditionFavoriteResponse) SetStatusCode(v int32) *DeleteAnalysisConditionFavoriteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAnalysisConditionFavoriteResponse) SetBody(v *DeleteAnalysisConditionFavoriteResponseBody) *DeleteAnalysisConditionFavoriteResponse {
	s.Body = v
	return s
}

func (s *DeleteAnalysisConditionFavoriteResponse) Validate() error {
	return dara.Validate(s)
}
