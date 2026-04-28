// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetShareLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetShareLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetShareLinkResponse
	GetStatusCode() *int32
	SetBody(v *ShareLink) *GetShareLinkResponse
	GetBody() *ShareLink
}

type GetShareLinkResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShareLink         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetShareLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s GetShareLinkResponse) GoString() string {
	return s.String()
}

func (s *GetShareLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetShareLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetShareLinkResponse) GetBody() *ShareLink {
	return s.Body
}

func (s *GetShareLinkResponse) SetHeaders(v map[string]*string) *GetShareLinkResponse {
	s.Headers = v
	return s
}

func (s *GetShareLinkResponse) SetStatusCode(v int32) *GetShareLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetShareLinkResponse) SetBody(v *ShareLink) *GetShareLinkResponse {
	s.Body = v
	return s
}

func (s *GetShareLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
