// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSiteResponse
	GetStatusCode() *int32
	SetBody(v *GetSiteResponseBody) *GetSiteResponse
	GetBody() *GetSiteResponseBody
}

type GetSiteResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSiteResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSiteResponse) GoString() string {
	return s.String()
}

func (s *GetSiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSiteResponse) GetBody() *GetSiteResponseBody {
	return s.Body
}

func (s *GetSiteResponse) SetHeaders(v map[string]*string) *GetSiteResponse {
	s.Headers = v
	return s
}

func (s *GetSiteResponse) SetStatusCode(v int32) *GetSiteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSiteResponse) SetBody(v *GetSiteResponseBody) *GetSiteResponse {
	s.Body = v
	return s
}

func (s *GetSiteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
