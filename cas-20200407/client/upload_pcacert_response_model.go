// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadPCACertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadPCACertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadPCACertResponse
	GetStatusCode() *int32
	SetBody(v *UploadPCACertResponseBody) *UploadPCACertResponse
	GetBody() *UploadPCACertResponseBody
}

type UploadPCACertResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadPCACertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadPCACertResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadPCACertResponse) GoString() string {
	return s.String()
}

func (s *UploadPCACertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadPCACertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadPCACertResponse) GetBody() *UploadPCACertResponseBody {
	return s.Body
}

func (s *UploadPCACertResponse) SetHeaders(v map[string]*string) *UploadPCACertResponse {
	s.Headers = v
	return s
}

func (s *UploadPCACertResponse) SetStatusCode(v int32) *UploadPCACertResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadPCACertResponse) SetBody(v *UploadPCACertResponseBody) *UploadPCACertResponse {
	s.Body = v
	return s
}

func (s *UploadPCACertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
