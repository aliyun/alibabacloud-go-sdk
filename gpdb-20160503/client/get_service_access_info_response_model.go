// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceAccessInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceAccessInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceAccessInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceAccessInfoResponseBody) *GetServiceAccessInfoResponse
	GetBody() *GetServiceAccessInfoResponseBody
}

type GetServiceAccessInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceAccessInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceAccessInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceAccessInfoResponse) GoString() string {
	return s.String()
}

func (s *GetServiceAccessInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceAccessInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceAccessInfoResponse) GetBody() *GetServiceAccessInfoResponseBody {
	return s.Body
}

func (s *GetServiceAccessInfoResponse) SetHeaders(v map[string]*string) *GetServiceAccessInfoResponse {
	s.Headers = v
	return s
}

func (s *GetServiceAccessInfoResponse) SetStatusCode(v int32) *GetServiceAccessInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceAccessInfoResponse) SetBody(v *GetServiceAccessInfoResponseBody) *GetServiceAccessInfoResponse {
	s.Body = v
	return s
}

func (s *GetServiceAccessInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
