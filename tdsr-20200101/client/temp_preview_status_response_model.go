// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTempPreviewStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TempPreviewStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TempPreviewStatusResponse
	GetStatusCode() *int32
	SetBody(v *TempPreviewStatusResponseBody) *TempPreviewStatusResponse
	GetBody() *TempPreviewStatusResponseBody
}

type TempPreviewStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TempPreviewStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TempPreviewStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s TempPreviewStatusResponse) GoString() string {
	return s.String()
}

func (s *TempPreviewStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TempPreviewStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TempPreviewStatusResponse) GetBody() *TempPreviewStatusResponseBody {
	return s.Body
}

func (s *TempPreviewStatusResponse) SetHeaders(v map[string]*string) *TempPreviewStatusResponse {
	s.Headers = v
	return s
}

func (s *TempPreviewStatusResponse) SetStatusCode(v int32) *TempPreviewStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *TempPreviewStatusResponse) SetBody(v *TempPreviewStatusResponseBody) *TempPreviewStatusResponse {
	s.Body = v
	return s
}

func (s *TempPreviewStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
