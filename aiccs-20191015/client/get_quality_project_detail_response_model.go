// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityProjectDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityProjectDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityProjectDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityProjectDetailResponseBody) *GetQualityProjectDetailResponse
	GetBody() *GetQualityProjectDetailResponseBody
}

type GetQualityProjectDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityProjectDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityProjectDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityProjectDetailResponse) GoString() string {
	return s.String()
}

func (s *GetQualityProjectDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityProjectDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityProjectDetailResponse) GetBody() *GetQualityProjectDetailResponseBody {
	return s.Body
}

func (s *GetQualityProjectDetailResponse) SetHeaders(v map[string]*string) *GetQualityProjectDetailResponse {
	s.Headers = v
	return s
}

func (s *GetQualityProjectDetailResponse) SetStatusCode(v int32) *GetQualityProjectDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityProjectDetailResponse) SetBody(v *GetQualityProjectDetailResponseBody) *GetQualityProjectDetailResponse {
	s.Body = v
	return s
}

func (s *GetQualityProjectDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
