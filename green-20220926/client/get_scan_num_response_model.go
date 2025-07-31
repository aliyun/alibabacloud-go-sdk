// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScanNumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScanNumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScanNumResponse
	GetStatusCode() *int32
	SetBody(v *GetScanNumResponseBody) *GetScanNumResponse
	GetBody() *GetScanNumResponseBody
}

type GetScanNumResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScanNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScanNumResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScanNumResponse) GoString() string {
	return s.String()
}

func (s *GetScanNumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScanNumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScanNumResponse) GetBody() *GetScanNumResponseBody {
	return s.Body
}

func (s *GetScanNumResponse) SetHeaders(v map[string]*string) *GetScanNumResponse {
	s.Headers = v
	return s
}

func (s *GetScanNumResponse) SetStatusCode(v int32) *GetScanNumResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScanNumResponse) SetBody(v *GetScanNumResponseBody) *GetScanNumResponse {
	s.Body = v
	return s
}

func (s *GetScanNumResponse) Validate() error {
	return dara.Validate(s)
}
