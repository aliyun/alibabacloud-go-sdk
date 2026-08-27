// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecommendNextActionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecommendNextActionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecommendNextActionsResponse
	GetStatusCode() *int32
	SetBody(v *RecommendNextActionsResponseBody) *RecommendNextActionsResponse
	GetBody() *RecommendNextActionsResponseBody
}

type RecommendNextActionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecommendNextActionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecommendNextActionsResponse) String() string {
	return dara.Prettify(s)
}

func (s RecommendNextActionsResponse) GoString() string {
	return s.String()
}

func (s *RecommendNextActionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecommendNextActionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecommendNextActionsResponse) GetBody() *RecommendNextActionsResponseBody {
	return s.Body
}

func (s *RecommendNextActionsResponse) SetHeaders(v map[string]*string) *RecommendNextActionsResponse {
	s.Headers = v
	return s
}

func (s *RecommendNextActionsResponse) SetStatusCode(v int32) *RecommendNextActionsResponse {
	s.StatusCode = &v
	return s
}

func (s *RecommendNextActionsResponse) SetBody(v *RecommendNextActionsResponseBody) *RecommendNextActionsResponse {
	s.Body = v
	return s
}

func (s *RecommendNextActionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
