// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoTaskResultResponseBody) *GetVideoTaskResultResponse
	GetBody() *GetVideoTaskResultResponseBody
}

type GetVideoTaskResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetVideoTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoTaskResultResponse) GetBody() *GetVideoTaskResultResponseBody {
	return s.Body
}

func (s *GetVideoTaskResultResponse) SetHeaders(v map[string]*string) *GetVideoTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetVideoTaskResultResponse) SetStatusCode(v int32) *GetVideoTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoTaskResultResponse) SetBody(v *GetVideoTaskResultResponseBody) *GetVideoTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetVideoTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
