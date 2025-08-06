// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKMSServiceKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKMSServiceKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKMSServiceKeyResponse
	GetStatusCode() *int32
	SetBody(v *GetKMSServiceKeyResponseBody) *GetKMSServiceKeyResponse
	GetBody() *GetKMSServiceKeyResponseBody
}

type GetKMSServiceKeyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKMSServiceKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKMSServiceKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKMSServiceKeyResponse) GoString() string {
	return s.String()
}

func (s *GetKMSServiceKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKMSServiceKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKMSServiceKeyResponse) GetBody() *GetKMSServiceKeyResponseBody {
	return s.Body
}

func (s *GetKMSServiceKeyResponse) SetHeaders(v map[string]*string) *GetKMSServiceKeyResponse {
	s.Headers = v
	return s
}

func (s *GetKMSServiceKeyResponse) SetStatusCode(v int32) *GetKMSServiceKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKMSServiceKeyResponse) SetBody(v *GetKMSServiceKeyResponseBody) *GetKMSServiceKeyResponse {
	s.Body = v
	return s
}

func (s *GetKMSServiceKeyResponse) Validate() error {
	return dara.Validate(s)
}
