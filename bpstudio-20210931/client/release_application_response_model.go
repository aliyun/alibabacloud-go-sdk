// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseApplicationResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseApplicationResponseBody) *ReleaseApplicationResponse
	GetBody() *ReleaseApplicationResponseBody
}

type ReleaseApplicationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseApplicationResponse) GoString() string {
	return s.String()
}

func (s *ReleaseApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseApplicationResponse) GetBody() *ReleaseApplicationResponseBody {
	return s.Body
}

func (s *ReleaseApplicationResponse) SetHeaders(v map[string]*string) *ReleaseApplicationResponse {
	s.Headers = v
	return s
}

func (s *ReleaseApplicationResponse) SetStatusCode(v int32) *ReleaseApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseApplicationResponse) SetBody(v *ReleaseApplicationResponseBody) *ReleaseApplicationResponse {
	s.Body = v
	return s
}

func (s *ReleaseApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
