// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFpShotDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFpShotDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFpShotDBResponse
	GetStatusCode() *int32
	SetBody(v *ListFpShotDBResponseBody) *ListFpShotDBResponse
	GetBody() *ListFpShotDBResponseBody
}

type ListFpShotDBResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFpShotDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFpShotDBResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFpShotDBResponse) GoString() string {
	return s.String()
}

func (s *ListFpShotDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFpShotDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFpShotDBResponse) GetBody() *ListFpShotDBResponseBody {
	return s.Body
}

func (s *ListFpShotDBResponse) SetHeaders(v map[string]*string) *ListFpShotDBResponse {
	s.Headers = v
	return s
}

func (s *ListFpShotDBResponse) SetStatusCode(v int32) *ListFpShotDBResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFpShotDBResponse) SetBody(v *ListFpShotDBResponseBody) *ListFpShotDBResponse {
	s.Body = v
	return s
}

func (s *ListFpShotDBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
