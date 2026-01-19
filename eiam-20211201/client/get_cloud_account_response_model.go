// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCloudAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCloudAccountResponse
	GetStatusCode() *int32
	SetBody(v *GetCloudAccountResponseBody) *GetCloudAccountResponse
	GetBody() *GetCloudAccountResponseBody
}

type GetCloudAccountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCloudAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCloudAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAccountResponse) GoString() string {
	return s.String()
}

func (s *GetCloudAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCloudAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCloudAccountResponse) GetBody() *GetCloudAccountResponseBody {
	return s.Body
}

func (s *GetCloudAccountResponse) SetHeaders(v map[string]*string) *GetCloudAccountResponse {
	s.Headers = v
	return s
}

func (s *GetCloudAccountResponse) SetStatusCode(v int32) *GetCloudAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCloudAccountResponse) SetBody(v *GetCloudAccountResponseBody) *GetCloudAccountResponse {
	s.Body = v
	return s
}

func (s *GetCloudAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
