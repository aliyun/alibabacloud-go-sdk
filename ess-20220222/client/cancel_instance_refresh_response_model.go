// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelInstanceRefreshResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelInstanceRefreshResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelInstanceRefreshResponse
	GetStatusCode() *int32
	SetBody(v *CancelInstanceRefreshResponseBody) *CancelInstanceRefreshResponse
	GetBody() *CancelInstanceRefreshResponseBody
}

type CancelInstanceRefreshResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelInstanceRefreshResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelInstanceRefreshResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelInstanceRefreshResponse) GoString() string {
	return s.String()
}

func (s *CancelInstanceRefreshResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelInstanceRefreshResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelInstanceRefreshResponse) GetBody() *CancelInstanceRefreshResponseBody {
	return s.Body
}

func (s *CancelInstanceRefreshResponse) SetHeaders(v map[string]*string) *CancelInstanceRefreshResponse {
	s.Headers = v
	return s
}

func (s *CancelInstanceRefreshResponse) SetStatusCode(v int32) *CancelInstanceRefreshResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelInstanceRefreshResponse) SetBody(v *CancelInstanceRefreshResponseBody) *CancelInstanceRefreshResponse {
	s.Body = v
	return s
}

func (s *CancelInstanceRefreshResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
