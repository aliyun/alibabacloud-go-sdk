// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUploadPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUploadPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateUploadPolicyResponseBody) *CreateUploadPolicyResponse
	GetBody() *CreateUploadPolicyResponseBody
}

type CreateUploadPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUploadPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUploadPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateUploadPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUploadPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUploadPolicyResponse) GetBody() *CreateUploadPolicyResponseBody {
	return s.Body
}

func (s *CreateUploadPolicyResponse) SetHeaders(v map[string]*string) *CreateUploadPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateUploadPolicyResponse) SetStatusCode(v int32) *CreateUploadPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUploadPolicyResponse) SetBody(v *CreateUploadPolicyResponseBody) *CreateUploadPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateUploadPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
