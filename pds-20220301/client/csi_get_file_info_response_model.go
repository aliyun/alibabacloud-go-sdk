// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCsiGetFileInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CsiGetFileInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CsiGetFileInfoResponse
	GetStatusCode() *int32
	SetBody(v *CsiGetFileInfoResponseBody) *CsiGetFileInfoResponse
	GetBody() *CsiGetFileInfoResponseBody
}

type CsiGetFileInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CsiGetFileInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CsiGetFileInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s CsiGetFileInfoResponse) GoString() string {
	return s.String()
}

func (s *CsiGetFileInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CsiGetFileInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CsiGetFileInfoResponse) GetBody() *CsiGetFileInfoResponseBody {
	return s.Body
}

func (s *CsiGetFileInfoResponse) SetHeaders(v map[string]*string) *CsiGetFileInfoResponse {
	s.Headers = v
	return s
}

func (s *CsiGetFileInfoResponse) SetStatusCode(v int32) *CsiGetFileInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CsiGetFileInfoResponse) SetBody(v *CsiGetFileInfoResponseBody) *CsiGetFileInfoResponse {
	s.Body = v
	return s
}

func (s *CsiGetFileInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
