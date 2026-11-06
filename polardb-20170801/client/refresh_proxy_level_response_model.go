// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshProxyLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshProxyLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshProxyLevelResponse
	GetStatusCode() *int32
	SetBody(v *RefreshProxyLevelResponseBody) *RefreshProxyLevelResponse
	GetBody() *RefreshProxyLevelResponseBody
}

type RefreshProxyLevelResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshProxyLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshProxyLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshProxyLevelResponse) GoString() string {
	return s.String()
}

func (s *RefreshProxyLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshProxyLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshProxyLevelResponse) GetBody() *RefreshProxyLevelResponseBody {
	return s.Body
}

func (s *RefreshProxyLevelResponse) SetHeaders(v map[string]*string) *RefreshProxyLevelResponse {
	s.Headers = v
	return s
}

func (s *RefreshProxyLevelResponse) SetStatusCode(v int32) *RefreshProxyLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshProxyLevelResponse) SetBody(v *RefreshProxyLevelResponseBody) *RefreshProxyLevelResponse {
	s.Body = v
	return s
}

func (s *RefreshProxyLevelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
