// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveTranscodeTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveTranscodeTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveTranscodeTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveTranscodeTemplatesResponseBody) *ListLiveTranscodeTemplatesResponse
	GetBody() *ListLiveTranscodeTemplatesResponseBody
}

type ListLiveTranscodeTemplatesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveTranscodeTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveTranscodeTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTranscodeTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListLiveTranscodeTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveTranscodeTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveTranscodeTemplatesResponse) GetBody() *ListLiveTranscodeTemplatesResponseBody {
	return s.Body
}

func (s *ListLiveTranscodeTemplatesResponse) SetHeaders(v map[string]*string) *ListLiveTranscodeTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListLiveTranscodeTemplatesResponse) SetStatusCode(v int32) *ListLiveTranscodeTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveTranscodeTemplatesResponse) SetBody(v *ListLiveTranscodeTemplatesResponseBody) *ListLiveTranscodeTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListLiveTranscodeTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
