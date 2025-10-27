// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneyPotUploadPolicyInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHoneyPotUploadPolicyInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHoneyPotUploadPolicyInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetHoneyPotUploadPolicyInfoResponseBody) *GetHoneyPotUploadPolicyInfoResponse
	GetBody() *GetHoneyPotUploadPolicyInfoResponseBody
}

type GetHoneyPotUploadPolicyInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHoneyPotUploadPolicyInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHoneyPotUploadPolicyInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHoneyPotUploadPolicyInfoResponse) GoString() string {
	return s.String()
}

func (s *GetHoneyPotUploadPolicyInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHoneyPotUploadPolicyInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHoneyPotUploadPolicyInfoResponse) GetBody() *GetHoneyPotUploadPolicyInfoResponseBody {
	return s.Body
}

func (s *GetHoneyPotUploadPolicyInfoResponse) SetHeaders(v map[string]*string) *GetHoneyPotUploadPolicyInfoResponse {
	s.Headers = v
	return s
}

func (s *GetHoneyPotUploadPolicyInfoResponse) SetStatusCode(v int32) *GetHoneyPotUploadPolicyInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHoneyPotUploadPolicyInfoResponse) SetBody(v *GetHoneyPotUploadPolicyInfoResponseBody) *GetHoneyPotUploadPolicyInfoResponse {
	s.Body = v
	return s
}

func (s *GetHoneyPotUploadPolicyInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
