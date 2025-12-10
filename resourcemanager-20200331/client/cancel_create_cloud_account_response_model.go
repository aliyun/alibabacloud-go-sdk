// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCreateCloudAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelCreateCloudAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelCreateCloudAccountResponse
	GetStatusCode() *int32
	SetBody(v *CancelCreateCloudAccountResponseBody) *CancelCreateCloudAccountResponse
	GetBody() *CancelCreateCloudAccountResponseBody
}

type CancelCreateCloudAccountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelCreateCloudAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelCreateCloudAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelCreateCloudAccountResponse) GoString() string {
	return s.String()
}

func (s *CancelCreateCloudAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelCreateCloudAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelCreateCloudAccountResponse) GetBody() *CancelCreateCloudAccountResponseBody {
	return s.Body
}

func (s *CancelCreateCloudAccountResponse) SetHeaders(v map[string]*string) *CancelCreateCloudAccountResponse {
	s.Headers = v
	return s
}

func (s *CancelCreateCloudAccountResponse) SetStatusCode(v int32) *CancelCreateCloudAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelCreateCloudAccountResponse) SetBody(v *CancelCreateCloudAccountResponseBody) *CancelCreateCloudAccountResponse {
	s.Body = v
	return s
}

func (s *CancelCreateCloudAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
