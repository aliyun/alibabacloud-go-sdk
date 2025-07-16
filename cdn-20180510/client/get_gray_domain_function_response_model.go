// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGrayDomainFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGrayDomainFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGrayDomainFunctionResponse
	GetStatusCode() *int32
	SetBody(v *GetGrayDomainFunctionResponseBody) *GetGrayDomainFunctionResponse
	GetBody() *GetGrayDomainFunctionResponseBody
}

type GetGrayDomainFunctionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGrayDomainFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGrayDomainFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGrayDomainFunctionResponse) GoString() string {
	return s.String()
}

func (s *GetGrayDomainFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGrayDomainFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGrayDomainFunctionResponse) GetBody() *GetGrayDomainFunctionResponseBody {
	return s.Body
}

func (s *GetGrayDomainFunctionResponse) SetHeaders(v map[string]*string) *GetGrayDomainFunctionResponse {
	s.Headers = v
	return s
}

func (s *GetGrayDomainFunctionResponse) SetStatusCode(v int32) *GetGrayDomainFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGrayDomainFunctionResponse) SetBody(v *GetGrayDomainFunctionResponseBody) *GetGrayDomainFunctionResponse {
	s.Body = v
	return s
}

func (s *GetGrayDomainFunctionResponse) Validate() error {
	return dara.Validate(s)
}
