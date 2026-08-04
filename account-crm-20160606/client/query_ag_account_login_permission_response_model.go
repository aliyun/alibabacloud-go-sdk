// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAgAccountLoginPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAgAccountLoginPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAgAccountLoginPermissionResponse
	GetStatusCode() *int32
	SetBody(v *QueryAgAccountLoginPermissionResponseBody) *QueryAgAccountLoginPermissionResponse
	GetBody() *QueryAgAccountLoginPermissionResponseBody
}

type QueryAgAccountLoginPermissionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAgAccountLoginPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAgAccountLoginPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAgAccountLoginPermissionResponse) GoString() string {
	return s.String()
}

func (s *QueryAgAccountLoginPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAgAccountLoginPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAgAccountLoginPermissionResponse) GetBody() *QueryAgAccountLoginPermissionResponseBody {
	return s.Body
}

func (s *QueryAgAccountLoginPermissionResponse) SetHeaders(v map[string]*string) *QueryAgAccountLoginPermissionResponse {
	s.Headers = v
	return s
}

func (s *QueryAgAccountLoginPermissionResponse) SetStatusCode(v int32) *QueryAgAccountLoginPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAgAccountLoginPermissionResponse) SetBody(v *QueryAgAccountLoginPermissionResponseBody) *QueryAgAccountLoginPermissionResponse {
	s.Body = v
	return s
}

func (s *QueryAgAccountLoginPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
