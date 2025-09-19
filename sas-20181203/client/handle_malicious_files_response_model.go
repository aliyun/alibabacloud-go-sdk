// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleMaliciousFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HandleMaliciousFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HandleMaliciousFilesResponse
	GetStatusCode() *int32
	SetBody(v *HandleMaliciousFilesResponseBody) *HandleMaliciousFilesResponse
	GetBody() *HandleMaliciousFilesResponseBody
}

type HandleMaliciousFilesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HandleMaliciousFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HandleMaliciousFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s HandleMaliciousFilesResponse) GoString() string {
	return s.String()
}

func (s *HandleMaliciousFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HandleMaliciousFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HandleMaliciousFilesResponse) GetBody() *HandleMaliciousFilesResponseBody {
	return s.Body
}

func (s *HandleMaliciousFilesResponse) SetHeaders(v map[string]*string) *HandleMaliciousFilesResponse {
	s.Headers = v
	return s
}

func (s *HandleMaliciousFilesResponse) SetStatusCode(v int32) *HandleMaliciousFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *HandleMaliciousFilesResponse) SetBody(v *HandleMaliciousFilesResponseBody) *HandleMaliciousFilesResponse {
	s.Body = v
	return s
}

func (s *HandleMaliciousFilesResponse) Validate() error {
	return dara.Validate(s)
}
