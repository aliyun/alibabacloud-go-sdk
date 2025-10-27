// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenSensitiveFileScanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenSensitiveFileScanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenSensitiveFileScanResponse
	GetStatusCode() *int32
	SetBody(v *OpenSensitiveFileScanResponseBody) *OpenSensitiveFileScanResponse
	GetBody() *OpenSensitiveFileScanResponseBody
}

type OpenSensitiveFileScanResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenSensitiveFileScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenSensitiveFileScanResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenSensitiveFileScanResponse) GoString() string {
	return s.String()
}

func (s *OpenSensitiveFileScanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenSensitiveFileScanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenSensitiveFileScanResponse) GetBody() *OpenSensitiveFileScanResponseBody {
	return s.Body
}

func (s *OpenSensitiveFileScanResponse) SetHeaders(v map[string]*string) *OpenSensitiveFileScanResponse {
	s.Headers = v
	return s
}

func (s *OpenSensitiveFileScanResponse) SetStatusCode(v int32) *OpenSensitiveFileScanResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenSensitiveFileScanResponse) SetBody(v *OpenSensitiveFileScanResponseBody) *OpenSensitiveFileScanResponse {
	s.Body = v
	return s
}

func (s *OpenSensitiveFileScanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
