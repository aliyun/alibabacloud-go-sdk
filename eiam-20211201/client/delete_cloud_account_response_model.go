// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudAccountResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudAccountResponseBody) *DeleteCloudAccountResponse
	GetBody() *DeleteCloudAccountResponseBody
}

type DeleteCloudAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudAccountResponse) GetBody() *DeleteCloudAccountResponseBody {
	return s.Body
}

func (s *DeleteCloudAccountResponse) SetHeaders(v map[string]*string) *DeleteCloudAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudAccountResponse) SetStatusCode(v int32) *DeleteCloudAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudAccountResponse) SetBody(v *DeleteCloudAccountResponseBody) *DeleteCloudAccountResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
