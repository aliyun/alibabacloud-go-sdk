// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveTranscodeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveTranscodeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveTranscodeJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveTranscodeJobResponseBody) *DeleteLiveTranscodeJobResponse
	GetBody() *DeleteLiveTranscodeJobResponseBody
}

type DeleteLiveTranscodeJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveTranscodeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveTranscodeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveTranscodeJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveTranscodeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveTranscodeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveTranscodeJobResponse) GetBody() *DeleteLiveTranscodeJobResponseBody {
	return s.Body
}

func (s *DeleteLiveTranscodeJobResponse) SetHeaders(v map[string]*string) *DeleteLiveTranscodeJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveTranscodeJobResponse) SetStatusCode(v int32) *DeleteLiveTranscodeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveTranscodeJobResponse) SetBody(v *DeleteLiveTranscodeJobResponseBody) *DeleteLiveTranscodeJobResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveTranscodeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
