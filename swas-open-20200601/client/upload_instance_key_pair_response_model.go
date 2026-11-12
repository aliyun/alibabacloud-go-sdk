// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadInstanceKeyPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadInstanceKeyPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadInstanceKeyPairResponse
	GetStatusCode() *int32
	SetBody(v *UploadInstanceKeyPairResponseBody) *UploadInstanceKeyPairResponse
	GetBody() *UploadInstanceKeyPairResponseBody
}

type UploadInstanceKeyPairResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadInstanceKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadInstanceKeyPairResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadInstanceKeyPairResponse) GoString() string {
	return s.String()
}

func (s *UploadInstanceKeyPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadInstanceKeyPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadInstanceKeyPairResponse) GetBody() *UploadInstanceKeyPairResponseBody {
	return s.Body
}

func (s *UploadInstanceKeyPairResponse) SetHeaders(v map[string]*string) *UploadInstanceKeyPairResponse {
	s.Headers = v
	return s
}

func (s *UploadInstanceKeyPairResponse) SetStatusCode(v int32) *UploadInstanceKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadInstanceKeyPairResponse) SetBody(v *UploadInstanceKeyPairResponseBody) *UploadInstanceKeyPairResponse {
	s.Body = v
	return s
}

func (s *UploadInstanceKeyPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
