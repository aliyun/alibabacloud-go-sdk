// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteCurrentNSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSiteCurrentNSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSiteCurrentNSResponse
	GetStatusCode() *int32
	SetBody(v *GetSiteCurrentNSResponseBody) *GetSiteCurrentNSResponse
	GetBody() *GetSiteCurrentNSResponseBody
}

type GetSiteCurrentNSResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSiteCurrentNSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSiteCurrentNSResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSiteCurrentNSResponse) GoString() string {
	return s.String()
}

func (s *GetSiteCurrentNSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSiteCurrentNSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSiteCurrentNSResponse) GetBody() *GetSiteCurrentNSResponseBody {
	return s.Body
}

func (s *GetSiteCurrentNSResponse) SetHeaders(v map[string]*string) *GetSiteCurrentNSResponse {
	s.Headers = v
	return s
}

func (s *GetSiteCurrentNSResponse) SetStatusCode(v int32) *GetSiteCurrentNSResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSiteCurrentNSResponse) SetBody(v *GetSiteCurrentNSResponseBody) *GetSiteCurrentNSResponse {
	s.Body = v
	return s
}

func (s *GetSiteCurrentNSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
