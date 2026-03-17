// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClientUserDNSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClientUserDNSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClientUserDNSResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClientUserDNSResponseBody) *ModifyClientUserDNSResponse
	GetBody() *ModifyClientUserDNSResponseBody
}

type ModifyClientUserDNSResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClientUserDNSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClientUserDNSResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClientUserDNSResponse) GoString() string {
	return s.String()
}

func (s *ModifyClientUserDNSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClientUserDNSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClientUserDNSResponse) GetBody() *ModifyClientUserDNSResponseBody {
	return s.Body
}

func (s *ModifyClientUserDNSResponse) SetHeaders(v map[string]*string) *ModifyClientUserDNSResponse {
	s.Headers = v
	return s
}

func (s *ModifyClientUserDNSResponse) SetStatusCode(v int32) *ModifyClientUserDNSResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClientUserDNSResponse) SetBody(v *ModifyClientUserDNSResponseBody) *ModifyClientUserDNSResponse {
	s.Body = v
	return s
}

func (s *ModifyClientUserDNSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
