// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualificationUploadPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualificationUploadPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualificationUploadPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetQualificationUploadPolicyResponseBody) *GetQualificationUploadPolicyResponse
	GetBody() *GetQualificationUploadPolicyResponseBody
}

type GetQualificationUploadPolicyResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualificationUploadPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualificationUploadPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualificationUploadPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetQualificationUploadPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualificationUploadPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualificationUploadPolicyResponse) GetBody() *GetQualificationUploadPolicyResponseBody {
	return s.Body
}

func (s *GetQualificationUploadPolicyResponse) SetHeaders(v map[string]*string) *GetQualificationUploadPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetQualificationUploadPolicyResponse) SetStatusCode(v int32) *GetQualificationUploadPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualificationUploadPolicyResponse) SetBody(v *GetQualificationUploadPolicyResponseBody) *GetQualificationUploadPolicyResponse {
	s.Body = v
	return s
}

func (s *GetQualificationUploadPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
