// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportInterveneFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportInterveneFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportInterveneFileResponse
	GetStatusCode() *int32
	SetBody(v *ImportInterveneFileResponseBody) *ImportInterveneFileResponse
	GetBody() *ImportInterveneFileResponseBody
}

type ImportInterveneFileResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportInterveneFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportInterveneFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportInterveneFileResponse) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportInterveneFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportInterveneFileResponse) GetBody() *ImportInterveneFileResponseBody {
	return s.Body
}

func (s *ImportInterveneFileResponse) SetHeaders(v map[string]*string) *ImportInterveneFileResponse {
	s.Headers = v
	return s
}

func (s *ImportInterveneFileResponse) SetStatusCode(v int32) *ImportInterveneFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportInterveneFileResponse) SetBody(v *ImportInterveneFileResponseBody) *ImportInterveneFileResponse {
	s.Body = v
	return s
}

func (s *ImportInterveneFileResponse) Validate() error {
	return dara.Validate(s)
}
