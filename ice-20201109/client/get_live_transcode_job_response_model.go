// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveTranscodeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLiveTranscodeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLiveTranscodeJobResponse
	GetStatusCode() *int32
	SetBody(v *GetLiveTranscodeJobResponseBody) *GetLiveTranscodeJobResponse
	GetBody() *GetLiveTranscodeJobResponseBody
}

type GetLiveTranscodeJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLiveTranscodeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLiveTranscodeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLiveTranscodeJobResponse) GoString() string {
	return s.String()
}

func (s *GetLiveTranscodeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLiveTranscodeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLiveTranscodeJobResponse) GetBody() *GetLiveTranscodeJobResponseBody {
	return s.Body
}

func (s *GetLiveTranscodeJobResponse) SetHeaders(v map[string]*string) *GetLiveTranscodeJobResponse {
	s.Headers = v
	return s
}

func (s *GetLiveTranscodeJobResponse) SetStatusCode(v int32) *GetLiveTranscodeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveTranscodeJobResponse) SetBody(v *GetLiveTranscodeJobResponseBody) *GetLiveTranscodeJobResponse {
	s.Body = v
	return s
}

func (s *GetLiveTranscodeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
