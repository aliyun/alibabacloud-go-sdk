// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRumUploadFileUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRumUploadFileUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRumUploadFileUrlResponse
	GetStatusCode() *int32
	SetBody(v *CreateRumUploadFileUrlResponseBody) *CreateRumUploadFileUrlResponse
	GetBody() *CreateRumUploadFileUrlResponseBody
}

type CreateRumUploadFileUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRumUploadFileUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRumUploadFileUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRumUploadFileUrlResponse) GoString() string {
	return s.String()
}

func (s *CreateRumUploadFileUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRumUploadFileUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRumUploadFileUrlResponse) GetBody() *CreateRumUploadFileUrlResponseBody {
	return s.Body
}

func (s *CreateRumUploadFileUrlResponse) SetHeaders(v map[string]*string) *CreateRumUploadFileUrlResponse {
	s.Headers = v
	return s
}

func (s *CreateRumUploadFileUrlResponse) SetStatusCode(v int32) *CreateRumUploadFileUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRumUploadFileUrlResponse) SetBody(v *CreateRumUploadFileUrlResponseBody) *CreateRumUploadFileUrlResponse {
	s.Body = v
	return s
}

func (s *CreateRumUploadFileUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
