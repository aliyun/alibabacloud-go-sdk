// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInstanceEntitlementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppInstanceEntitlementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppInstanceEntitlementResponse
	GetStatusCode() *int32
	SetBody(v *GetAppInstanceEntitlementResponseBody) *GetAppInstanceEntitlementResponse
	GetBody() *GetAppInstanceEntitlementResponseBody
}

type GetAppInstanceEntitlementResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppInstanceEntitlementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppInstanceEntitlementResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceEntitlementResponse) GoString() string {
	return s.String()
}

func (s *GetAppInstanceEntitlementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppInstanceEntitlementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppInstanceEntitlementResponse) GetBody() *GetAppInstanceEntitlementResponseBody {
	return s.Body
}

func (s *GetAppInstanceEntitlementResponse) SetHeaders(v map[string]*string) *GetAppInstanceEntitlementResponse {
	s.Headers = v
	return s
}

func (s *GetAppInstanceEntitlementResponse) SetStatusCode(v int32) *GetAppInstanceEntitlementResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppInstanceEntitlementResponse) SetBody(v *GetAppInstanceEntitlementResponseBody) *GetAppInstanceEntitlementResponse {
	s.Body = v
	return s
}

func (s *GetAppInstanceEntitlementResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
