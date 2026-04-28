// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetShareLinkTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetShareLinkTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetShareLinkTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetShareLinkTokenResponseBody) *GetShareLinkTokenResponse
	GetBody() *GetShareLinkTokenResponseBody
}

type GetShareLinkTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetShareLinkTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetShareLinkTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetShareLinkTokenResponse) GoString() string {
	return s.String()
}

func (s *GetShareLinkTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetShareLinkTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetShareLinkTokenResponse) GetBody() *GetShareLinkTokenResponseBody {
	return s.Body
}

func (s *GetShareLinkTokenResponse) SetHeaders(v map[string]*string) *GetShareLinkTokenResponse {
	s.Headers = v
	return s
}

func (s *GetShareLinkTokenResponse) SetStatusCode(v int32) *GetShareLinkTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetShareLinkTokenResponse) SetBody(v *GetShareLinkTokenResponseBody) *GetShareLinkTokenResponse {
	s.Body = v
	return s
}

func (s *GetShareLinkTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
