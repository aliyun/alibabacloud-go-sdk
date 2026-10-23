// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDBTenantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgenticDBTenantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgenticDBTenantResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgenticDBTenantResponseBody) *DeleteAgenticDBTenantResponse
	GetBody() *DeleteAgenticDBTenantResponseBody
}

type DeleteAgenticDBTenantResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgenticDBTenantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgenticDBTenantResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDBTenantResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDBTenantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgenticDBTenantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgenticDBTenantResponse) GetBody() *DeleteAgenticDBTenantResponseBody {
	return s.Body
}

func (s *DeleteAgenticDBTenantResponse) SetHeaders(v map[string]*string) *DeleteAgenticDBTenantResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgenticDBTenantResponse) SetStatusCode(v int32) *DeleteAgenticDBTenantResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgenticDBTenantResponse) SetBody(v *DeleteAgenticDBTenantResponseBody) *DeleteAgenticDBTenantResponse {
	s.Body = v
	return s
}

func (s *DeleteAgenticDBTenantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
