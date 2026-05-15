// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityProjectListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityProjectListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityProjectListResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityProjectListResponseBody) *GetQualityProjectListResponse
	GetBody() *GetQualityProjectListResponseBody
}

type GetQualityProjectListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityProjectListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityProjectListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityProjectListResponse) GoString() string {
	return s.String()
}

func (s *GetQualityProjectListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityProjectListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityProjectListResponse) GetBody() *GetQualityProjectListResponseBody {
	return s.Body
}

func (s *GetQualityProjectListResponse) SetHeaders(v map[string]*string) *GetQualityProjectListResponse {
	s.Headers = v
	return s
}

func (s *GetQualityProjectListResponse) SetStatusCode(v int32) *GetQualityProjectListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityProjectListResponse) SetBody(v *GetQualityProjectListResponseBody) *GetQualityProjectListResponse {
	s.Body = v
	return s
}

func (s *GetQualityProjectListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
