// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnBindLeaderInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnBindLeaderInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnBindLeaderInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UnBindLeaderInstanceResponseBody) *UnBindLeaderInstanceResponse
	GetBody() *UnBindLeaderInstanceResponseBody
}

type UnBindLeaderInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnBindLeaderInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnBindLeaderInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UnBindLeaderInstanceResponse) GoString() string {
	return s.String()
}

func (s *UnBindLeaderInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnBindLeaderInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnBindLeaderInstanceResponse) GetBody() *UnBindLeaderInstanceResponseBody {
	return s.Body
}

func (s *UnBindLeaderInstanceResponse) SetHeaders(v map[string]*string) *UnBindLeaderInstanceResponse {
	s.Headers = v
	return s
}

func (s *UnBindLeaderInstanceResponse) SetStatusCode(v int32) *UnBindLeaderInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnBindLeaderInstanceResponse) SetBody(v *UnBindLeaderInstanceResponseBody) *UnBindLeaderInstanceResponse {
	s.Body = v
	return s
}

func (s *UnBindLeaderInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
