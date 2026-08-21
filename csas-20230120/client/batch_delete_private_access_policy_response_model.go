// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeletePrivateAccessPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeletePrivateAccessPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeletePrivateAccessPolicyResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeletePrivateAccessPolicyResponseBody) *BatchDeletePrivateAccessPolicyResponse
	GetBody() *BatchDeletePrivateAccessPolicyResponseBody
}

type BatchDeletePrivateAccessPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeletePrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeletePrivateAccessPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeletePrivateAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *BatchDeletePrivateAccessPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeletePrivateAccessPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeletePrivateAccessPolicyResponse) GetBody() *BatchDeletePrivateAccessPolicyResponseBody {
	return s.Body
}

func (s *BatchDeletePrivateAccessPolicyResponse) SetHeaders(v map[string]*string) *BatchDeletePrivateAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *BatchDeletePrivateAccessPolicyResponse) SetStatusCode(v int32) *BatchDeletePrivateAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeletePrivateAccessPolicyResponse) SetBody(v *BatchDeletePrivateAccessPolicyResponseBody) *BatchDeletePrivateAccessPolicyResponse {
	s.Body = v
	return s
}

func (s *BatchDeletePrivateAccessPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
