// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *RefreshDataSourceResponseBody) *RefreshDataSourceResponse
	GetBody() *RefreshDataSourceResponseBody
}

type RefreshDataSourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshDataSourceResponse) GoString() string {
	return s.String()
}

func (s *RefreshDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshDataSourceResponse) GetBody() *RefreshDataSourceResponseBody {
	return s.Body
}

func (s *RefreshDataSourceResponse) SetHeaders(v map[string]*string) *RefreshDataSourceResponse {
	s.Headers = v
	return s
}

func (s *RefreshDataSourceResponse) SetStatusCode(v int32) *RefreshDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshDataSourceResponse) SetBody(v *RefreshDataSourceResponseBody) *RefreshDataSourceResponse {
	s.Body = v
	return s
}

func (s *RefreshDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
