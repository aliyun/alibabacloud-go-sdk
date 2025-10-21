// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoDetectShotTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitVideoDetectShotTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitVideoDetectShotTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitVideoDetectShotTaskResponseBody) *SubmitVideoDetectShotTaskResponse
	GetBody() *SubmitVideoDetectShotTaskResponseBody
}

type SubmitVideoDetectShotTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitVideoDetectShotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitVideoDetectShotTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoDetectShotTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitVideoDetectShotTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitVideoDetectShotTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitVideoDetectShotTaskResponse) GetBody() *SubmitVideoDetectShotTaskResponseBody {
	return s.Body
}

func (s *SubmitVideoDetectShotTaskResponse) SetHeaders(v map[string]*string) *SubmitVideoDetectShotTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitVideoDetectShotTaskResponse) SetStatusCode(v int32) *SubmitVideoDetectShotTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitVideoDetectShotTaskResponse) SetBody(v *SubmitVideoDetectShotTaskResponseBody) *SubmitVideoDetectShotTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitVideoDetectShotTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
