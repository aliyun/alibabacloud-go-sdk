// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnterpriseCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEnterpriseCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEnterpriseCodeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEnterpriseCodeResponseBody) *DeleteEnterpriseCodeResponse
	GetBody() *DeleteEnterpriseCodeResponseBody
}

type DeleteEnterpriseCodeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEnterpriseCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEnterpriseCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnterpriseCodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEnterpriseCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEnterpriseCodeResponse) GetBody() *DeleteEnterpriseCodeResponseBody {
	return s.Body
}

func (s *DeleteEnterpriseCodeResponse) SetHeaders(v map[string]*string) *DeleteEnterpriseCodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnterpriseCodeResponse) SetStatusCode(v int32) *DeleteEnterpriseCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEnterpriseCodeResponse) SetBody(v *DeleteEnterpriseCodeResponseBody) *DeleteEnterpriseCodeResponse {
	s.Body = v
	return s
}

func (s *DeleteEnterpriseCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
