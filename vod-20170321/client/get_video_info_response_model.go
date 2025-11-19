// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoInfoResponseBody) *GetVideoInfoResponse
	GetBody() *GetVideoInfoResponseBody
}

type GetVideoInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoInfoResponse) GoString() string {
	return s.String()
}

func (s *GetVideoInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoInfoResponse) GetBody() *GetVideoInfoResponseBody {
	return s.Body
}

func (s *GetVideoInfoResponse) SetHeaders(v map[string]*string) *GetVideoInfoResponse {
	s.Headers = v
	return s
}

func (s *GetVideoInfoResponse) SetStatusCode(v int32) *GetVideoInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoInfoResponse) SetBody(v *GetVideoInfoResponseBody) *GetVideoInfoResponse {
	s.Body = v
	return s
}

func (s *GetVideoInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
