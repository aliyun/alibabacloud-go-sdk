// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateDNSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrivateDNSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrivateDNSResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrivateDNSResponseBody) *DeletePrivateDNSResponse
	GetBody() *DeletePrivateDNSResponseBody
}

type DeletePrivateDNSResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrivateDNSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrivateDNSResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateDNSResponse) GoString() string {
	return s.String()
}

func (s *DeletePrivateDNSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrivateDNSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrivateDNSResponse) GetBody() *DeletePrivateDNSResponseBody {
	return s.Body
}

func (s *DeletePrivateDNSResponse) SetHeaders(v map[string]*string) *DeletePrivateDNSResponse {
	s.Headers = v
	return s
}

func (s *DeletePrivateDNSResponse) SetStatusCode(v int32) *DeletePrivateDNSResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrivateDNSResponse) SetBody(v *DeletePrivateDNSResponseBody) *DeletePrivateDNSResponse {
	s.Body = v
	return s
}

func (s *DeletePrivateDNSResponse) Validate() error {
	return dara.Validate(s)
}
