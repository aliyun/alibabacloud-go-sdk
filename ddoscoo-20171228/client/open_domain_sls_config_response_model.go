// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDomainSlsConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenDomainSlsConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenDomainSlsConfigResponse
	GetStatusCode() *int32
	SetBody(v *OpenDomainSlsConfigResponseBody) *OpenDomainSlsConfigResponse
	GetBody() *OpenDomainSlsConfigResponseBody
}

type OpenDomainSlsConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenDomainSlsConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenDomainSlsConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenDomainSlsConfigResponse) GoString() string {
	return s.String()
}

func (s *OpenDomainSlsConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenDomainSlsConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenDomainSlsConfigResponse) GetBody() *OpenDomainSlsConfigResponseBody {
	return s.Body
}

func (s *OpenDomainSlsConfigResponse) SetHeaders(v map[string]*string) *OpenDomainSlsConfigResponse {
	s.Headers = v
	return s
}

func (s *OpenDomainSlsConfigResponse) SetStatusCode(v int32) *OpenDomainSlsConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenDomainSlsConfigResponse) SetBody(v *OpenDomainSlsConfigResponseBody) *OpenDomainSlsConfigResponse {
	s.Body = v
	return s
}

func (s *OpenDomainSlsConfigResponse) Validate() error {
	return dara.Validate(s)
}
