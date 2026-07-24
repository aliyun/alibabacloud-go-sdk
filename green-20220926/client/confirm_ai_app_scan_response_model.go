// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmAiAppScanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfirmAiAppScanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfirmAiAppScanResponse
	GetStatusCode() *int32
	SetBody(v *ConfirmAiAppScanResponseBody) *ConfirmAiAppScanResponse
	GetBody() *ConfirmAiAppScanResponseBody
}

type ConfirmAiAppScanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmAiAppScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmAiAppScanResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfirmAiAppScanResponse) GoString() string {
	return s.String()
}

func (s *ConfirmAiAppScanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfirmAiAppScanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfirmAiAppScanResponse) GetBody() *ConfirmAiAppScanResponseBody {
	return s.Body
}

func (s *ConfirmAiAppScanResponse) SetHeaders(v map[string]*string) *ConfirmAiAppScanResponse {
	s.Headers = v
	return s
}

func (s *ConfirmAiAppScanResponse) SetStatusCode(v int32) *ConfirmAiAppScanResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmAiAppScanResponse) SetBody(v *ConfirmAiAppScanResponseBody) *ConfirmAiAppScanResponse {
	s.Body = v
	return s
}

func (s *ConfirmAiAppScanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
