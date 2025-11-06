// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAppEnhanceProcessInMsaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserAppEnhanceProcessInMsaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserAppEnhanceProcessInMsaResponse
	GetStatusCode() *int32
	SetBody(v *GetUserAppEnhanceProcessInMsaResponseBody) *GetUserAppEnhanceProcessInMsaResponse
	GetBody() *GetUserAppEnhanceProcessInMsaResponseBody
}

type GetUserAppEnhanceProcessInMsaResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserAppEnhanceProcessInMsaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserAppEnhanceProcessInMsaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserAppEnhanceProcessInMsaResponse) GoString() string {
	return s.String()
}

func (s *GetUserAppEnhanceProcessInMsaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserAppEnhanceProcessInMsaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserAppEnhanceProcessInMsaResponse) GetBody() *GetUserAppEnhanceProcessInMsaResponseBody {
	return s.Body
}

func (s *GetUserAppEnhanceProcessInMsaResponse) SetHeaders(v map[string]*string) *GetUserAppEnhanceProcessInMsaResponse {
	s.Headers = v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponse) SetStatusCode(v int32) *GetUserAppEnhanceProcessInMsaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponse) SetBody(v *GetUserAppEnhanceProcessInMsaResponseBody) *GetUserAppEnhanceProcessInMsaResponse {
	s.Body = v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
