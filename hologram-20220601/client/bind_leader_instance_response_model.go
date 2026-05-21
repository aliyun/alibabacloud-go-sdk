// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindLeaderInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindLeaderInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindLeaderInstanceResponse
	GetStatusCode() *int32
	SetBody(v *BindLeaderInstanceResponseBody) *BindLeaderInstanceResponse
	GetBody() *BindLeaderInstanceResponseBody
}

type BindLeaderInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindLeaderInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindLeaderInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s BindLeaderInstanceResponse) GoString() string {
	return s.String()
}

func (s *BindLeaderInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindLeaderInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindLeaderInstanceResponse) GetBody() *BindLeaderInstanceResponseBody {
	return s.Body
}

func (s *BindLeaderInstanceResponse) SetHeaders(v map[string]*string) *BindLeaderInstanceResponse {
	s.Headers = v
	return s
}

func (s *BindLeaderInstanceResponse) SetStatusCode(v int32) *BindLeaderInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *BindLeaderInstanceResponse) SetBody(v *BindLeaderInstanceResponseBody) *BindLeaderInstanceResponse {
	s.Body = v
	return s
}

func (s *BindLeaderInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
