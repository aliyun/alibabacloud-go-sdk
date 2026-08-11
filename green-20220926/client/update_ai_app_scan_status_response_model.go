// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiAppScanStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAiAppScanStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAiAppScanStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAiAppScanStatusResponseBody) *UpdateAiAppScanStatusResponse
	GetBody() *UpdateAiAppScanStatusResponseBody
}

type UpdateAiAppScanStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAiAppScanStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAiAppScanStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiAppScanStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateAiAppScanStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAiAppScanStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAiAppScanStatusResponse) GetBody() *UpdateAiAppScanStatusResponseBody {
	return s.Body
}

func (s *UpdateAiAppScanStatusResponse) SetHeaders(v map[string]*string) *UpdateAiAppScanStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateAiAppScanStatusResponse) SetStatusCode(v int32) *UpdateAiAppScanStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAiAppScanStatusResponse) SetBody(v *UpdateAiAppScanStatusResponseBody) *UpdateAiAppScanStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateAiAppScanStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
