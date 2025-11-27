// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCDiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRCDiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRCDiskResponse
	GetStatusCode() *int32
	SetBody(v *CreateRCDiskResponseBody) *CreateRCDiskResponse
	GetBody() *CreateRCDiskResponseBody
}

type CreateRCDiskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRCDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRCDiskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRCDiskResponse) GoString() string {
	return s.String()
}

func (s *CreateRCDiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRCDiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRCDiskResponse) GetBody() *CreateRCDiskResponseBody {
	return s.Body
}

func (s *CreateRCDiskResponse) SetHeaders(v map[string]*string) *CreateRCDiskResponse {
	s.Headers = v
	return s
}

func (s *CreateRCDiskResponse) SetStatusCode(v int32) *CreateRCDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRCDiskResponse) SetBody(v *CreateRCDiskResponseBody) *CreateRCDiskResponse {
	s.Body = v
	return s
}

func (s *CreateRCDiskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
