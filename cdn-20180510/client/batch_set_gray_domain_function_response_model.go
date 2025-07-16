// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetGrayDomainFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchSetGrayDomainFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchSetGrayDomainFunctionResponse
	GetStatusCode() *int32
	SetBody(v *BatchSetGrayDomainFunctionResponseBody) *BatchSetGrayDomainFunctionResponse
	GetBody() *BatchSetGrayDomainFunctionResponseBody
}

type BatchSetGrayDomainFunctionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSetGrayDomainFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSetGrayDomainFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchSetGrayDomainFunctionResponse) GoString() string {
	return s.String()
}

func (s *BatchSetGrayDomainFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchSetGrayDomainFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchSetGrayDomainFunctionResponse) GetBody() *BatchSetGrayDomainFunctionResponseBody {
	return s.Body
}

func (s *BatchSetGrayDomainFunctionResponse) SetHeaders(v map[string]*string) *BatchSetGrayDomainFunctionResponse {
	s.Headers = v
	return s
}

func (s *BatchSetGrayDomainFunctionResponse) SetStatusCode(v int32) *BatchSetGrayDomainFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSetGrayDomainFunctionResponse) SetBody(v *BatchSetGrayDomainFunctionResponseBody) *BatchSetGrayDomainFunctionResponse {
	s.Body = v
	return s
}

func (s *BatchSetGrayDomainFunctionResponse) Validate() error {
	return dara.Validate(s)
}
