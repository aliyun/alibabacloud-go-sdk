// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScanResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScanResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScanResultResponse
	GetStatusCode() *int32
	SetBody(v *GetScanResultResponseBody) *GetScanResultResponse
	GetBody() *GetScanResultResponseBody
}

type GetScanResultResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScanResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScanResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScanResultResponse) GoString() string {
	return s.String()
}

func (s *GetScanResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScanResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScanResultResponse) GetBody() *GetScanResultResponseBody {
	return s.Body
}

func (s *GetScanResultResponse) SetHeaders(v map[string]*string) *GetScanResultResponse {
	s.Headers = v
	return s
}

func (s *GetScanResultResponse) SetStatusCode(v int32) *GetScanResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScanResultResponse) SetBody(v *GetScanResultResponseBody) *GetScanResultResponse {
	s.Body = v
	return s
}

func (s *GetScanResultResponse) Validate() error {
	return dara.Validate(s)
}
