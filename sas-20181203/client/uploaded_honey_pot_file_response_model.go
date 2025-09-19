// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadedHoneyPotFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadedHoneyPotFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadedHoneyPotFileResponse
	GetStatusCode() *int32
	SetBody(v *UploadedHoneyPotFileResponseBody) *UploadedHoneyPotFileResponse
	GetBody() *UploadedHoneyPotFileResponseBody
}

type UploadedHoneyPotFileResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadedHoneyPotFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadedHoneyPotFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadedHoneyPotFileResponse) GoString() string {
	return s.String()
}

func (s *UploadedHoneyPotFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadedHoneyPotFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadedHoneyPotFileResponse) GetBody() *UploadedHoneyPotFileResponseBody {
	return s.Body
}

func (s *UploadedHoneyPotFileResponse) SetHeaders(v map[string]*string) *UploadedHoneyPotFileResponse {
	s.Headers = v
	return s
}

func (s *UploadedHoneyPotFileResponse) SetStatusCode(v int32) *UploadedHoneyPotFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadedHoneyPotFileResponse) SetBody(v *UploadedHoneyPotFileResponseBody) *UploadedHoneyPotFileResponse {
	s.Body = v
	return s
}

func (s *UploadedHoneyPotFileResponse) Validate() error {
	return dara.Validate(s)
}
