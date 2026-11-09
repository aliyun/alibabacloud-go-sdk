// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomAuthPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCustomAuthPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCustomAuthPermissionResponse
	GetStatusCode() *int32
	SetBody(v *QueryCustomAuthPermissionResponseBody) *QueryCustomAuthPermissionResponse
	GetBody() *QueryCustomAuthPermissionResponseBody
}

type QueryCustomAuthPermissionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCustomAuthPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCustomAuthPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomAuthPermissionResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomAuthPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCustomAuthPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCustomAuthPermissionResponse) GetBody() *QueryCustomAuthPermissionResponseBody {
	return s.Body
}

func (s *QueryCustomAuthPermissionResponse) SetHeaders(v map[string]*string) *QueryCustomAuthPermissionResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomAuthPermissionResponse) SetStatusCode(v int32) *QueryCustomAuthPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCustomAuthPermissionResponse) SetBody(v *QueryCustomAuthPermissionResponseBody) *QueryCustomAuthPermissionResponse {
	s.Body = v
	return s
}

func (s *QueryCustomAuthPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
