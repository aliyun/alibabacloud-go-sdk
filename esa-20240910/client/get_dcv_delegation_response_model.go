// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDcvDelegationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDcvDelegationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDcvDelegationResponse
	GetStatusCode() *int32
	SetBody(v *GetDcvDelegationResponseBody) *GetDcvDelegationResponse
	GetBody() *GetDcvDelegationResponseBody
}

type GetDcvDelegationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDcvDelegationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDcvDelegationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDcvDelegationResponse) GoString() string {
	return s.String()
}

func (s *GetDcvDelegationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDcvDelegationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDcvDelegationResponse) GetBody() *GetDcvDelegationResponseBody {
	return s.Body
}

func (s *GetDcvDelegationResponse) SetHeaders(v map[string]*string) *GetDcvDelegationResponse {
	s.Headers = v
	return s
}

func (s *GetDcvDelegationResponse) SetStatusCode(v int32) *GetDcvDelegationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDcvDelegationResponse) SetBody(v *GetDcvDelegationResponseBody) *GetDcvDelegationResponse {
	s.Body = v
	return s
}

func (s *GetDcvDelegationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
