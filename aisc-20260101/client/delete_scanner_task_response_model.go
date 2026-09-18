// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScannerTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScannerTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScannerTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteScannerTaskResponseBody) *DeleteScannerTaskResponse
	GetBody() *DeleteScannerTaskResponseBody
}

type DeleteScannerTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScannerTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScannerTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScannerTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteScannerTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScannerTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScannerTaskResponse) GetBody() *DeleteScannerTaskResponseBody {
	return s.Body
}

func (s *DeleteScannerTaskResponse) SetHeaders(v map[string]*string) *DeleteScannerTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteScannerTaskResponse) SetStatusCode(v int32) *DeleteScannerTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScannerTaskResponse) SetBody(v *DeleteScannerTaskResponseBody) *DeleteScannerTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteScannerTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
