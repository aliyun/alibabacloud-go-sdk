// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnInstallClusterAddonsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnInstallClusterAddonsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnInstallClusterAddonsResponse
	GetStatusCode() *int32
	SetBody(v *UnInstallClusterAddonsResponseBody) *UnInstallClusterAddonsResponse
	GetBody() *UnInstallClusterAddonsResponseBody
}

type UnInstallClusterAddonsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnInstallClusterAddonsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnInstallClusterAddonsResponse) String() string {
	return dara.Prettify(s)
}

func (s UnInstallClusterAddonsResponse) GoString() string {
	return s.String()
}

func (s *UnInstallClusterAddonsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnInstallClusterAddonsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnInstallClusterAddonsResponse) GetBody() *UnInstallClusterAddonsResponseBody {
	return s.Body
}

func (s *UnInstallClusterAddonsResponse) SetHeaders(v map[string]*string) *UnInstallClusterAddonsResponse {
	s.Headers = v
	return s
}

func (s *UnInstallClusterAddonsResponse) SetStatusCode(v int32) *UnInstallClusterAddonsResponse {
	s.StatusCode = &v
	return s
}

func (s *UnInstallClusterAddonsResponse) SetBody(v *UnInstallClusterAddonsResponseBody) *UnInstallClusterAddonsResponse {
	s.Body = v
	return s
}

func (s *UnInstallClusterAddonsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
