// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyDrivesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMyDrivesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMyDrivesResponse
	GetStatusCode() *int32
	SetBody(v *ListMyDrivesResponseBody) *ListMyDrivesResponse
	GetBody() *ListMyDrivesResponseBody
}

type ListMyDrivesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMyDrivesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMyDrivesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMyDrivesResponse) GoString() string {
	return s.String()
}

func (s *ListMyDrivesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMyDrivesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMyDrivesResponse) GetBody() *ListMyDrivesResponseBody {
	return s.Body
}

func (s *ListMyDrivesResponse) SetHeaders(v map[string]*string) *ListMyDrivesResponse {
	s.Headers = v
	return s
}

func (s *ListMyDrivesResponse) SetStatusCode(v int32) *ListMyDrivesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMyDrivesResponse) SetBody(v *ListMyDrivesResponseBody) *ListMyDrivesResponse {
	s.Body = v
	return s
}

func (s *ListMyDrivesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
