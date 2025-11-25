// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDomainResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDomainResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDomainResourceResponseBody) *DeleteDomainResourceResponse
	GetBody() *DeleteDomainResourceResponseBody
}

type DeleteDomainResourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDomainResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDomainResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDomainResourceResponse) GetBody() *DeleteDomainResourceResponseBody {
	return s.Body
}

func (s *DeleteDomainResourceResponse) SetHeaders(v map[string]*string) *DeleteDomainResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainResourceResponse) SetStatusCode(v int32) *DeleteDomainResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainResourceResponse) SetBody(v *DeleteDomainResourceResponseBody) *DeleteDomainResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteDomainResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
