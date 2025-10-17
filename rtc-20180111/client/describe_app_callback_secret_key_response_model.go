// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppCallbackSecretKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppCallbackSecretKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppCallbackSecretKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppCallbackSecretKeyResponseBody) *DescribeAppCallbackSecretKeyResponse
	GetBody() *DescribeAppCallbackSecretKeyResponseBody
}

type DescribeAppCallbackSecretKeyResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppCallbackSecretKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppCallbackSecretKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppCallbackSecretKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppCallbackSecretKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppCallbackSecretKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppCallbackSecretKeyResponse) GetBody() *DescribeAppCallbackSecretKeyResponseBody {
	return s.Body
}

func (s *DescribeAppCallbackSecretKeyResponse) SetHeaders(v map[string]*string) *DescribeAppCallbackSecretKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppCallbackSecretKeyResponse) SetStatusCode(v int32) *DescribeAppCallbackSecretKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppCallbackSecretKeyResponse) SetBody(v *DescribeAppCallbackSecretKeyResponseBody) *DescribeAppCallbackSecretKeyResponse {
	s.Body = v
	return s
}

func (s *DescribeAppCallbackSecretKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
