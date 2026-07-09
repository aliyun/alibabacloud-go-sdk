// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSensitiveScanResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSensitiveScanResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSensitiveScanResultResponse
	GetStatusCode() *int32
	SetBody(v *GetSensitiveScanResultResponseBody) *GetSensitiveScanResultResponse
	GetBody() *GetSensitiveScanResultResponseBody
}

type GetSensitiveScanResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSensitiveScanResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSensitiveScanResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSensitiveScanResultResponse) GoString() string {
	return s.String()
}

func (s *GetSensitiveScanResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSensitiveScanResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSensitiveScanResultResponse) GetBody() *GetSensitiveScanResultResponseBody {
	return s.Body
}

func (s *GetSensitiveScanResultResponse) SetHeaders(v map[string]*string) *GetSensitiveScanResultResponse {
	s.Headers = v
	return s
}

func (s *GetSensitiveScanResultResponse) SetStatusCode(v int32) *GetSensitiveScanResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSensitiveScanResultResponse) SetBody(v *GetSensitiveScanResultResponseBody) *GetSensitiveScanResultResponse {
	s.Body = v
	return s
}

func (s *GetSensitiveScanResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
