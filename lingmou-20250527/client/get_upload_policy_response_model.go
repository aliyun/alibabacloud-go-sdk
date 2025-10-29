// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUploadPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUploadPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetUploadPolicyResponseBody) *GetUploadPolicyResponse
	GetBody() *GetUploadPolicyResponseBody
}

type GetUploadPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUploadPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUploadPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUploadPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetUploadPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUploadPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUploadPolicyResponse) GetBody() *GetUploadPolicyResponseBody {
	return s.Body
}

func (s *GetUploadPolicyResponse) SetHeaders(v map[string]*string) *GetUploadPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetUploadPolicyResponse) SetStatusCode(v int32) *GetUploadPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadPolicyResponse) SetBody(v *GetUploadPolicyResponseBody) *GetUploadPolicyResponse {
	s.Body = v
	return s
}

func (s *GetUploadPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
