// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnalysisConditionFavoriteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAnalysisConditionFavoriteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAnalysisConditionFavoriteResponse
	GetStatusCode() *int32
	SetBody(v *CreateAnalysisConditionFavoriteResponseBody) *CreateAnalysisConditionFavoriteResponse
	GetBody() *CreateAnalysisConditionFavoriteResponseBody
}

type CreateAnalysisConditionFavoriteResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAnalysisConditionFavoriteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAnalysisConditionFavoriteResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAnalysisConditionFavoriteResponse) GoString() string {
	return s.String()
}

func (s *CreateAnalysisConditionFavoriteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAnalysisConditionFavoriteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAnalysisConditionFavoriteResponse) GetBody() *CreateAnalysisConditionFavoriteResponseBody {
	return s.Body
}

func (s *CreateAnalysisConditionFavoriteResponse) SetHeaders(v map[string]*string) *CreateAnalysisConditionFavoriteResponse {
	s.Headers = v
	return s
}

func (s *CreateAnalysisConditionFavoriteResponse) SetStatusCode(v int32) *CreateAnalysisConditionFavoriteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAnalysisConditionFavoriteResponse) SetBody(v *CreateAnalysisConditionFavoriteResponseBody) *CreateAnalysisConditionFavoriteResponse {
	s.Body = v
	return s
}

func (s *CreateAnalysisConditionFavoriteResponse) Validate() error {
	return dara.Validate(s)
}
