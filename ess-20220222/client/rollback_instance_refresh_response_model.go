// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackInstanceRefreshResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackInstanceRefreshResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackInstanceRefreshResponse
	GetStatusCode() *int32
	SetBody(v *RollbackInstanceRefreshResponseBody) *RollbackInstanceRefreshResponse
	GetBody() *RollbackInstanceRefreshResponseBody
}

type RollbackInstanceRefreshResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackInstanceRefreshResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackInstanceRefreshResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackInstanceRefreshResponse) GoString() string {
	return s.String()
}

func (s *RollbackInstanceRefreshResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackInstanceRefreshResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackInstanceRefreshResponse) GetBody() *RollbackInstanceRefreshResponseBody {
	return s.Body
}

func (s *RollbackInstanceRefreshResponse) SetHeaders(v map[string]*string) *RollbackInstanceRefreshResponse {
	s.Headers = v
	return s
}

func (s *RollbackInstanceRefreshResponse) SetStatusCode(v int32) *RollbackInstanceRefreshResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackInstanceRefreshResponse) SetBody(v *RollbackInstanceRefreshResponseBody) *RollbackInstanceRefreshResponse {
	s.Body = v
	return s
}

func (s *RollbackInstanceRefreshResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
