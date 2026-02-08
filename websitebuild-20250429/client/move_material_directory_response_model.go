// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveMaterialDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveMaterialDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveMaterialDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *MoveMaterialDirectoryResponseBody) *MoveMaterialDirectoryResponse
	GetBody() *MoveMaterialDirectoryResponseBody
}

type MoveMaterialDirectoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveMaterialDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveMaterialDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveMaterialDirectoryResponse) GoString() string {
	return s.String()
}

func (s *MoveMaterialDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveMaterialDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveMaterialDirectoryResponse) GetBody() *MoveMaterialDirectoryResponseBody {
	return s.Body
}

func (s *MoveMaterialDirectoryResponse) SetHeaders(v map[string]*string) *MoveMaterialDirectoryResponse {
	s.Headers = v
	return s
}

func (s *MoveMaterialDirectoryResponse) SetStatusCode(v int32) *MoveMaterialDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveMaterialDirectoryResponse) SetBody(v *MoveMaterialDirectoryResponseBody) *MoveMaterialDirectoryResponse {
	s.Body = v
	return s
}

func (s *MoveMaterialDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
