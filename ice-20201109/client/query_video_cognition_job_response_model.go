// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVideoCognitionJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryVideoCognitionJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryVideoCognitionJobResponse
	GetStatusCode() *int32
	SetBody(v *QueryVideoCognitionJobResponseBody) *QueryVideoCognitionJobResponse
	GetBody() *QueryVideoCognitionJobResponseBody
}

type QueryVideoCognitionJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVideoCognitionJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVideoCognitionJobResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryVideoCognitionJobResponse) GoString() string {
	return s.String()
}

func (s *QueryVideoCognitionJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryVideoCognitionJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryVideoCognitionJobResponse) GetBody() *QueryVideoCognitionJobResponseBody {
	return s.Body
}

func (s *QueryVideoCognitionJobResponse) SetHeaders(v map[string]*string) *QueryVideoCognitionJobResponse {
	s.Headers = v
	return s
}

func (s *QueryVideoCognitionJobResponse) SetStatusCode(v int32) *QueryVideoCognitionJobResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVideoCognitionJobResponse) SetBody(v *QueryVideoCognitionJobResponseBody) *QueryVideoCognitionJobResponse {
	s.Body = v
	return s
}

func (s *QueryVideoCognitionJobResponse) Validate() error {
	return dara.Validate(s)
}
