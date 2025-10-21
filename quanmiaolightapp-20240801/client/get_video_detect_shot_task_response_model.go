// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoDetectShotTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoDetectShotTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoDetectShotTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoDetectShotTaskResponseBody) *GetVideoDetectShotTaskResponse
	GetBody() *GetVideoDetectShotTaskResponseBody
}

type GetVideoDetectShotTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoDetectShotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoDetectShotTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDetectShotTaskResponse) GoString() string {
	return s.String()
}

func (s *GetVideoDetectShotTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoDetectShotTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoDetectShotTaskResponse) GetBody() *GetVideoDetectShotTaskResponseBody {
	return s.Body
}

func (s *GetVideoDetectShotTaskResponse) SetHeaders(v map[string]*string) *GetVideoDetectShotTaskResponse {
	s.Headers = v
	return s
}

func (s *GetVideoDetectShotTaskResponse) SetStatusCode(v int32) *GetVideoDetectShotTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoDetectShotTaskResponse) SetBody(v *GetVideoDetectShotTaskResponseBody) *GetVideoDetectShotTaskResponse {
	s.Body = v
	return s
}

func (s *GetVideoDetectShotTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
