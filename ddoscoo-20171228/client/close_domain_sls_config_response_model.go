// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseDomainSlsConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseDomainSlsConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseDomainSlsConfigResponse
	GetStatusCode() *int32
	SetBody(v *CloseDomainSlsConfigResponseBody) *CloseDomainSlsConfigResponse
	GetBody() *CloseDomainSlsConfigResponseBody
}

type CloseDomainSlsConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseDomainSlsConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseDomainSlsConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseDomainSlsConfigResponse) GoString() string {
	return s.String()
}

func (s *CloseDomainSlsConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseDomainSlsConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseDomainSlsConfigResponse) GetBody() *CloseDomainSlsConfigResponseBody {
	return s.Body
}

func (s *CloseDomainSlsConfigResponse) SetHeaders(v map[string]*string) *CloseDomainSlsConfigResponse {
	s.Headers = v
	return s
}

func (s *CloseDomainSlsConfigResponse) SetStatusCode(v int32) *CloseDomainSlsConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseDomainSlsConfigResponse) SetBody(v *CloseDomainSlsConfigResponseBody) *CloseDomainSlsConfigResponse {
	s.Body = v
	return s
}

func (s *CloseDomainSlsConfigResponse) Validate() error {
	return dara.Validate(s)
}
