// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteCustomLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSiteCustomLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSiteCustomLogResponse
	GetStatusCode() *int32
	SetBody(v *GetSiteCustomLogResponseBody) *GetSiteCustomLogResponse
	GetBody() *GetSiteCustomLogResponseBody
}

type GetSiteCustomLogResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSiteCustomLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSiteCustomLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSiteCustomLogResponse) GoString() string {
	return s.String()
}

func (s *GetSiteCustomLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSiteCustomLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSiteCustomLogResponse) GetBody() *GetSiteCustomLogResponseBody {
	return s.Body
}

func (s *GetSiteCustomLogResponse) SetHeaders(v map[string]*string) *GetSiteCustomLogResponse {
	s.Headers = v
	return s
}

func (s *GetSiteCustomLogResponse) SetStatusCode(v int32) *GetSiteCustomLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSiteCustomLogResponse) SetBody(v *GetSiteCustomLogResponseBody) *GetSiteCustomLogResponse {
	s.Body = v
	return s
}

func (s *GetSiteCustomLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
