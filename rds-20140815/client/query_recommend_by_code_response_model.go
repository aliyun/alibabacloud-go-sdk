// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecommendByCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecommendByCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecommendByCodeResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecommendByCodeResponseBody) *QueryRecommendByCodeResponse
	GetBody() *QueryRecommendByCodeResponseBody
}

type QueryRecommendByCodeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecommendByCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecommendByCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecommendByCodeResponse) GoString() string {
	return s.String()
}

func (s *QueryRecommendByCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecommendByCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecommendByCodeResponse) GetBody() *QueryRecommendByCodeResponseBody {
	return s.Body
}

func (s *QueryRecommendByCodeResponse) SetHeaders(v map[string]*string) *QueryRecommendByCodeResponse {
	s.Headers = v
	return s
}

func (s *QueryRecommendByCodeResponse) SetStatusCode(v int32) *QueryRecommendByCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecommendByCodeResponse) SetBody(v *QueryRecommendByCodeResponseBody) *QueryRecommendByCodeResponse {
	s.Body = v
	return s
}

func (s *QueryRecommendByCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
