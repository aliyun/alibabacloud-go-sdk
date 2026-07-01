// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRCSSignatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRCSSignatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRCSSignatureResponse
	GetStatusCode() *int32
	SetBody(v *GetRCSSignatureResponseBody) *GetRCSSignatureResponse
	GetBody() *GetRCSSignatureResponseBody
}

type GetRCSSignatureResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRCSSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRCSSignatureResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRCSSignatureResponse) GoString() string {
	return s.String()
}

func (s *GetRCSSignatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRCSSignatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRCSSignatureResponse) GetBody() *GetRCSSignatureResponseBody {
	return s.Body
}

func (s *GetRCSSignatureResponse) SetHeaders(v map[string]*string) *GetRCSSignatureResponse {
	s.Headers = v
	return s
}

func (s *GetRCSSignatureResponse) SetStatusCode(v int32) *GetRCSSignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRCSSignatureResponse) SetBody(v *GetRCSSignatureResponseBody) *GetRCSSignatureResponse {
	s.Body = v
	return s
}

func (s *GetRCSSignatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
