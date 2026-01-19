// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudAccountsResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudAccountsResponseBody) *ListCloudAccountsResponse
	GetBody() *ListCloudAccountsResponseBody
}

type ListCloudAccountsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListCloudAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudAccountsResponse) GetBody() *ListCloudAccountsResponseBody {
	return s.Body
}

func (s *ListCloudAccountsResponse) SetHeaders(v map[string]*string) *ListCloudAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListCloudAccountsResponse) SetStatusCode(v int32) *ListCloudAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudAccountsResponse) SetBody(v *ListCloudAccountsResponseBody) *ListCloudAccountsResponse {
	s.Body = v
	return s
}

func (s *ListCloudAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
