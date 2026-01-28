// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeVideoCastCrewListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeVideoCastCrewListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeVideoCastCrewListResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeVideoCastCrewListResponseBody) *RecognizeVideoCastCrewListResponse
	GetBody() *RecognizeVideoCastCrewListResponseBody
}

type RecognizeVideoCastCrewListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeVideoCastCrewListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeVideoCastCrewListResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeVideoCastCrewListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeVideoCastCrewListResponse) GetBody() *RecognizeVideoCastCrewListResponseBody {
	return s.Body
}

func (s *RecognizeVideoCastCrewListResponse) SetHeaders(v map[string]*string) *RecognizeVideoCastCrewListResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVideoCastCrewListResponse) SetStatusCode(v int32) *RecognizeVideoCastCrewListResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponse) SetBody(v *RecognizeVideoCastCrewListResponseBody) *RecognizeVideoCastCrewListResponse {
	s.Body = v
	return s
}

func (s *RecognizeVideoCastCrewListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
