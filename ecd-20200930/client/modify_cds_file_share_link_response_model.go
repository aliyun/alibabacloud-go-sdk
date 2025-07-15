// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdsFileShareLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCdsFileShareLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCdsFileShareLinkResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCdsFileShareLinkResponseBody) *ModifyCdsFileShareLinkResponse
	GetBody() *ModifyCdsFileShareLinkResponseBody
}

type ModifyCdsFileShareLinkResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCdsFileShareLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCdsFileShareLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdsFileShareLinkResponse) GoString() string {
	return s.String()
}

func (s *ModifyCdsFileShareLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCdsFileShareLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCdsFileShareLinkResponse) GetBody() *ModifyCdsFileShareLinkResponseBody {
	return s.Body
}

func (s *ModifyCdsFileShareLinkResponse) SetHeaders(v map[string]*string) *ModifyCdsFileShareLinkResponse {
	s.Headers = v
	return s
}

func (s *ModifyCdsFileShareLinkResponse) SetStatusCode(v int32) *ModifyCdsFileShareLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCdsFileShareLinkResponse) SetBody(v *ModifyCdsFileShareLinkResponseBody) *ModifyCdsFileShareLinkResponse {
	s.Body = v
	return s
}

func (s *ModifyCdsFileShareLinkResponse) Validate() error {
	return dara.Validate(s)
}
