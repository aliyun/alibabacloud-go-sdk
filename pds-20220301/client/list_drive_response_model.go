// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDriveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDriveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDriveResponse
	GetStatusCode() *int32
	SetBody(v *ListDriveResponseBody) *ListDriveResponse
	GetBody() *ListDriveResponseBody
}

type ListDriveResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDriveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDriveResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDriveResponse) GoString() string {
	return s.String()
}

func (s *ListDriveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDriveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDriveResponse) GetBody() *ListDriveResponseBody {
	return s.Body
}

func (s *ListDriveResponse) SetHeaders(v map[string]*string) *ListDriveResponse {
	s.Headers = v
	return s
}

func (s *ListDriveResponse) SetStatusCode(v int32) *ListDriveResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDriveResponse) SetBody(v *ListDriveResponseBody) *ListDriveResponse {
	s.Body = v
	return s
}

func (s *ListDriveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
