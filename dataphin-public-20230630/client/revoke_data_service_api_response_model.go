// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeDataServiceApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeDataServiceApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeDataServiceApiResponse
	GetStatusCode() *int32
	SetBody(v *RevokeDataServiceApiResponseBody) *RevokeDataServiceApiResponse
	GetBody() *RevokeDataServiceApiResponseBody
}

type RevokeDataServiceApiResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeDataServiceApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeDataServiceApiResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeDataServiceApiResponse) GoString() string {
	return s.String()
}

func (s *RevokeDataServiceApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeDataServiceApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeDataServiceApiResponse) GetBody() *RevokeDataServiceApiResponseBody {
	return s.Body
}

func (s *RevokeDataServiceApiResponse) SetHeaders(v map[string]*string) *RevokeDataServiceApiResponse {
	s.Headers = v
	return s
}

func (s *RevokeDataServiceApiResponse) SetStatusCode(v int32) *RevokeDataServiceApiResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeDataServiceApiResponse) SetBody(v *RevokeDataServiceApiResponseBody) *RevokeDataServiceApiResponse {
	s.Body = v
	return s
}

func (s *RevokeDataServiceApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
