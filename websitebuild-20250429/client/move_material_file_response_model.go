// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveMaterialFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveMaterialFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveMaterialFileResponse
	GetStatusCode() *int32
	SetBody(v *MoveMaterialFileResponseBody) *MoveMaterialFileResponse
	GetBody() *MoveMaterialFileResponseBody
}

type MoveMaterialFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveMaterialFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveMaterialFileResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveMaterialFileResponse) GoString() string {
	return s.String()
}

func (s *MoveMaterialFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveMaterialFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveMaterialFileResponse) GetBody() *MoveMaterialFileResponseBody {
	return s.Body
}

func (s *MoveMaterialFileResponse) SetHeaders(v map[string]*string) *MoveMaterialFileResponse {
	s.Headers = v
	return s
}

func (s *MoveMaterialFileResponse) SetStatusCode(v int32) *MoveMaterialFileResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveMaterialFileResponse) SetBody(v *MoveMaterialFileResponseBody) *MoveMaterialFileResponse {
	s.Body = v
	return s
}

func (s *MoveMaterialFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
