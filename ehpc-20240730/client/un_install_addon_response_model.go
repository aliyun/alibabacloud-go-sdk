// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnInstallAddonResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnInstallAddonResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnInstallAddonResponse
	GetStatusCode() *int32
	SetBody(v *UnInstallAddonResponseBody) *UnInstallAddonResponse
	GetBody() *UnInstallAddonResponseBody
}

type UnInstallAddonResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnInstallAddonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnInstallAddonResponse) String() string {
	return dara.Prettify(s)
}

func (s UnInstallAddonResponse) GoString() string {
	return s.String()
}

func (s *UnInstallAddonResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnInstallAddonResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnInstallAddonResponse) GetBody() *UnInstallAddonResponseBody {
	return s.Body
}

func (s *UnInstallAddonResponse) SetHeaders(v map[string]*string) *UnInstallAddonResponse {
	s.Headers = v
	return s
}

func (s *UnInstallAddonResponse) SetStatusCode(v int32) *UnInstallAddonResponse {
	s.StatusCode = &v
	return s
}

func (s *UnInstallAddonResponse) SetBody(v *UnInstallAddonResponseBody) *UnInstallAddonResponse {
	s.Body = v
	return s
}

func (s *UnInstallAddonResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
