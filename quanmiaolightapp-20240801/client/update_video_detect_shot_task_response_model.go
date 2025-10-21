// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoDetectShotTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVideoDetectShotTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVideoDetectShotTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVideoDetectShotTaskResponseBody) *UpdateVideoDetectShotTaskResponse
	GetBody() *UpdateVideoDetectShotTaskResponseBody
}

type UpdateVideoDetectShotTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVideoDetectShotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVideoDetectShotTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoDetectShotTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateVideoDetectShotTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVideoDetectShotTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVideoDetectShotTaskResponse) GetBody() *UpdateVideoDetectShotTaskResponseBody {
	return s.Body
}

func (s *UpdateVideoDetectShotTaskResponse) SetHeaders(v map[string]*string) *UpdateVideoDetectShotTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateVideoDetectShotTaskResponse) SetStatusCode(v int32) *UpdateVideoDetectShotTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVideoDetectShotTaskResponse) SetBody(v *UpdateVideoDetectShotTaskResponseBody) *UpdateVideoDetectShotTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateVideoDetectShotTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
