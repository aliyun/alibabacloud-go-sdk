// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudAccountResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudAccountResponseBody) *CreateCloudAccountResponse
	GetBody() *CreateCloudAccountResponseBody
}

type CreateCloudAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudAccountResponse) GetBody() *CreateCloudAccountResponseBody {
	return s.Body
}

func (s *CreateCloudAccountResponse) SetHeaders(v map[string]*string) *CreateCloudAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudAccountResponse) SetStatusCode(v int32) *CreateCloudAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudAccountResponse) SetBody(v *CreateCloudAccountResponseBody) *CreateCloudAccountResponse {
	s.Body = v
	return s
}

func (s *CreateCloudAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
