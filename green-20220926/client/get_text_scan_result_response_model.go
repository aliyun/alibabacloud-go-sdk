// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextScanResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTextScanResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTextScanResultResponse
	GetStatusCode() *int32
	SetBody(v *GetTextScanResultResponseBody) *GetTextScanResultResponse
	GetBody() *GetTextScanResultResponseBody
}

type GetTextScanResultResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTextScanResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTextScanResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTextScanResultResponse) GoString() string {
	return s.String()
}

func (s *GetTextScanResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTextScanResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTextScanResultResponse) GetBody() *GetTextScanResultResponseBody {
	return s.Body
}

func (s *GetTextScanResultResponse) SetHeaders(v map[string]*string) *GetTextScanResultResponse {
	s.Headers = v
	return s
}

func (s *GetTextScanResultResponse) SetStatusCode(v int32) *GetTextScanResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTextScanResultResponse) SetBody(v *GetTextScanResultResponseBody) *GetTextScanResultResponse {
	s.Body = v
	return s
}

func (s *GetTextScanResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
