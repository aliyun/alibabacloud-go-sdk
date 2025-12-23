// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudAccessResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudAccessResponseBody) *ListCloudAccessResponse
	GetBody() *ListCloudAccessResponseBody
}

type ListCloudAccessResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccessResponse) GoString() string {
	return s.String()
}

func (s *ListCloudAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudAccessResponse) GetBody() *ListCloudAccessResponseBody {
	return s.Body
}

func (s *ListCloudAccessResponse) SetHeaders(v map[string]*string) *ListCloudAccessResponse {
	s.Headers = v
	return s
}

func (s *ListCloudAccessResponse) SetStatusCode(v int32) *ListCloudAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudAccessResponse) SetBody(v *ListCloudAccessResponseBody) *ListCloudAccessResponse {
	s.Body = v
	return s
}

func (s *ListCloudAccessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
