// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeExtensionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeExtensionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeExtensionsResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeExtensionsResponseBody) *UpgradeExtensionsResponse
	GetBody() *UpgradeExtensionsResponseBody
}

type UpgradeExtensionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeExtensionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeExtensionsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeExtensionsResponse) GoString() string {
	return s.String()
}

func (s *UpgradeExtensionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeExtensionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeExtensionsResponse) GetBody() *UpgradeExtensionsResponseBody {
	return s.Body
}

func (s *UpgradeExtensionsResponse) SetHeaders(v map[string]*string) *UpgradeExtensionsResponse {
	s.Headers = v
	return s
}

func (s *UpgradeExtensionsResponse) SetStatusCode(v int32) *UpgradeExtensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeExtensionsResponse) SetBody(v *UpgradeExtensionsResponseBody) *UpgradeExtensionsResponse {
	s.Body = v
	return s
}

func (s *UpgradeExtensionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
