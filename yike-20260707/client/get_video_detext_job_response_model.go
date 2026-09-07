// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoDetextJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoDetextJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoDetextJobResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoDetextJobResponseBody) *GetVideoDetextJobResponse
	GetBody() *GetVideoDetextJobResponseBody
}

type GetVideoDetextJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoDetextJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoDetextJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDetextJobResponse) GoString() string {
	return s.String()
}

func (s *GetVideoDetextJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoDetextJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoDetextJobResponse) GetBody() *GetVideoDetextJobResponseBody {
	return s.Body
}

func (s *GetVideoDetextJobResponse) SetHeaders(v map[string]*string) *GetVideoDetextJobResponse {
	s.Headers = v
	return s
}

func (s *GetVideoDetextJobResponse) SetStatusCode(v int32) *GetVideoDetextJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoDetextJobResponse) SetBody(v *GetVideoDetextJobResponseBody) *GetVideoDetextJobResponse {
	s.Body = v
	return s
}

func (s *GetVideoDetextJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
