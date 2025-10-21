// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotUpdateJobResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotUpdateJobResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotUpdateJobResultResponse
	GetStatusCode() *int32
	SetBody(v *GetHotUpdateJobResultResponseBody) *GetHotUpdateJobResultResponse
	GetBody() *GetHotUpdateJobResultResponseBody
}

type GetHotUpdateJobResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotUpdateJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotUpdateJobResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotUpdateJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetHotUpdateJobResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotUpdateJobResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotUpdateJobResultResponse) GetBody() *GetHotUpdateJobResultResponseBody {
	return s.Body
}

func (s *GetHotUpdateJobResultResponse) SetHeaders(v map[string]*string) *GetHotUpdateJobResultResponse {
	s.Headers = v
	return s
}

func (s *GetHotUpdateJobResultResponse) SetStatusCode(v int32) *GetHotUpdateJobResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotUpdateJobResultResponse) SetBody(v *GetHotUpdateJobResultResponseBody) *GetHotUpdateJobResultResponse {
	s.Body = v
	return s
}

func (s *GetHotUpdateJobResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
