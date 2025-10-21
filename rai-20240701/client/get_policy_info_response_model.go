// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPolicyInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPolicyInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetPolicyInfoResponseBody) *GetPolicyInfoResponse
	GetBody() *GetPolicyInfoResponseBody
}

type GetPolicyInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPolicyInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPolicyInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPolicyInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPolicyInfoResponse) GetBody() *GetPolicyInfoResponseBody {
	return s.Body
}

func (s *GetPolicyInfoResponse) SetHeaders(v map[string]*string) *GetPolicyInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPolicyInfoResponse) SetStatusCode(v int32) *GetPolicyInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPolicyInfoResponse) SetBody(v *GetPolicyInfoResponseBody) *GetPolicyInfoResponse {
	s.Body = v
	return s
}

func (s *GetPolicyInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
