// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetShareLinkByAnonymousResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetShareLinkByAnonymousResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetShareLinkByAnonymousResponse
	GetStatusCode() *int32
	SetBody(v *GetShareLinkByAnonymousResponseBody) *GetShareLinkByAnonymousResponse
	GetBody() *GetShareLinkByAnonymousResponseBody
}

type GetShareLinkByAnonymousResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetShareLinkByAnonymousResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetShareLinkByAnonymousResponse) String() string {
	return dara.Prettify(s)
}

func (s GetShareLinkByAnonymousResponse) GoString() string {
	return s.String()
}

func (s *GetShareLinkByAnonymousResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetShareLinkByAnonymousResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetShareLinkByAnonymousResponse) GetBody() *GetShareLinkByAnonymousResponseBody {
	return s.Body
}

func (s *GetShareLinkByAnonymousResponse) SetHeaders(v map[string]*string) *GetShareLinkByAnonymousResponse {
	s.Headers = v
	return s
}

func (s *GetShareLinkByAnonymousResponse) SetStatusCode(v int32) *GetShareLinkByAnonymousResponse {
	s.StatusCode = &v
	return s
}

func (s *GetShareLinkByAnonymousResponse) SetBody(v *GetShareLinkByAnonymousResponseBody) *GetShareLinkByAnonymousResponse {
	s.Body = v
	return s
}

func (s *GetShareLinkByAnonymousResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
