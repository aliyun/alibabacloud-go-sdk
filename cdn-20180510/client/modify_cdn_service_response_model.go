// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdnServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCdnServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCdnServiceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCdnServiceResponseBody) *ModifyCdnServiceResponse
	GetBody() *ModifyCdnServiceResponseBody
}

type ModifyCdnServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCdnServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCdnServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdnServiceResponse) GoString() string {
	return s.String()
}

func (s *ModifyCdnServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCdnServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCdnServiceResponse) GetBody() *ModifyCdnServiceResponseBody {
	return s.Body
}

func (s *ModifyCdnServiceResponse) SetHeaders(v map[string]*string) *ModifyCdnServiceResponse {
	s.Headers = v
	return s
}

func (s *ModifyCdnServiceResponse) SetStatusCode(v int32) *ModifyCdnServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCdnServiceResponse) SetBody(v *ModifyCdnServiceResponseBody) *ModifyCdnServiceResponse {
	s.Body = v
	return s
}

func (s *ModifyCdnServiceResponse) Validate() error {
	return dara.Validate(s)
}
