// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScanResponse
	GetStatusCode() *int32
	SetBody(v *CreateScanResponseBody) *CreateScanResponse
	GetBody() *CreateScanResponseBody
}

type CreateScanResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScanResponse) GoString() string {
	return s.String()
}

func (s *CreateScanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScanResponse) GetBody() *CreateScanResponseBody {
	return s.Body
}

func (s *CreateScanResponse) SetHeaders(v map[string]*string) *CreateScanResponse {
	s.Headers = v
	return s
}

func (s *CreateScanResponse) SetStatusCode(v int32) *CreateScanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScanResponse) SetBody(v *CreateScanResponseBody) *CreateScanResponse {
	s.Body = v
	return s
}

func (s *CreateScanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
