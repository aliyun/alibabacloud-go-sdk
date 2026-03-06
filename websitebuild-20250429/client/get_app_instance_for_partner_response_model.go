// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInstanceForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppInstanceForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppInstanceForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *GetAppInstanceForPartnerResponseBody) *GetAppInstanceForPartnerResponse
	GetBody() *GetAppInstanceForPartnerResponseBody
}

type GetAppInstanceForPartnerResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppInstanceForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppInstanceForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponse) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppInstanceForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppInstanceForPartnerResponse) GetBody() *GetAppInstanceForPartnerResponseBody {
	return s.Body
}

func (s *GetAppInstanceForPartnerResponse) SetHeaders(v map[string]*string) *GetAppInstanceForPartnerResponse {
	s.Headers = v
	return s
}

func (s *GetAppInstanceForPartnerResponse) SetStatusCode(v int32) *GetAppInstanceForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppInstanceForPartnerResponse) SetBody(v *GetAppInstanceForPartnerResponseBody) *GetAppInstanceForPartnerResponse {
	s.Body = v
	return s
}

func (s *GetAppInstanceForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
