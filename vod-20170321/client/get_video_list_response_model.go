// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoListResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoListResponseBody) *GetVideoListResponse
	GetBody() *GetVideoListResponseBody
}

type GetVideoListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoListResponse) GoString() string {
	return s.String()
}

func (s *GetVideoListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoListResponse) GetBody() *GetVideoListResponseBody {
	return s.Body
}

func (s *GetVideoListResponse) SetHeaders(v map[string]*string) *GetVideoListResponse {
	s.Headers = v
	return s
}

func (s *GetVideoListResponse) SetStatusCode(v int32) *GetVideoListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoListResponse) SetBody(v *GetVideoListResponseBody) *GetVideoListResponse {
	s.Body = v
	return s
}

func (s *GetVideoListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
