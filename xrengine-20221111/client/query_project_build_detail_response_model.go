// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProjectBuildDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryProjectBuildDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryProjectBuildDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryProjectBuildDetailResponseBody) *QueryProjectBuildDetailResponse
	GetBody() *QueryProjectBuildDetailResponseBody
}

type QueryProjectBuildDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryProjectBuildDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryProjectBuildDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectBuildDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryProjectBuildDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryProjectBuildDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryProjectBuildDetailResponse) GetBody() *QueryProjectBuildDetailResponseBody {
	return s.Body
}

func (s *QueryProjectBuildDetailResponse) SetHeaders(v map[string]*string) *QueryProjectBuildDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryProjectBuildDetailResponse) SetStatusCode(v int32) *QueryProjectBuildDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProjectBuildDetailResponse) SetBody(v *QueryProjectBuildDetailResponseBody) *QueryProjectBuildDetailResponse {
	s.Body = v
	return s
}

func (s *QueryProjectBuildDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
