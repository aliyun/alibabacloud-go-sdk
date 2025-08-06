// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKMSServiceKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKMSServiceKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKMSServiceKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateKMSServiceKeyResponseBody) *CreateKMSServiceKeyResponse
	GetBody() *CreateKMSServiceKeyResponseBody
}

type CreateKMSServiceKeyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKMSServiceKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKMSServiceKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKMSServiceKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateKMSServiceKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKMSServiceKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKMSServiceKeyResponse) GetBody() *CreateKMSServiceKeyResponseBody {
	return s.Body
}

func (s *CreateKMSServiceKeyResponse) SetHeaders(v map[string]*string) *CreateKMSServiceKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateKMSServiceKeyResponse) SetStatusCode(v int32) *CreateKMSServiceKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKMSServiceKeyResponse) SetBody(v *CreateKMSServiceKeyResponseBody) *CreateKMSServiceKeyResponse {
	s.Body = v
	return s
}

func (s *CreateKMSServiceKeyResponse) Validate() error {
	return dara.Validate(s)
}
