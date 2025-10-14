// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSiteResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSiteResponseBody) *DeleteSiteResponse
	GetBody() *DeleteSiteResponseBody
}

type DeleteSiteResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSiteResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSiteResponse) GoString() string {
	return s.String()
}

func (s *DeleteSiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSiteResponse) GetBody() *DeleteSiteResponseBody {
	return s.Body
}

func (s *DeleteSiteResponse) SetHeaders(v map[string]*string) *DeleteSiteResponse {
	s.Headers = v
	return s
}

func (s *DeleteSiteResponse) SetStatusCode(v int32) *DeleteSiteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSiteResponse) SetBody(v *DeleteSiteResponseBody) *DeleteSiteResponse {
	s.Body = v
	return s
}

func (s *DeleteSiteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
