// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceNetworkSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRCInstanceNetworkSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRCInstanceNetworkSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRCInstanceNetworkSpecResponseBody) *ModifyRCInstanceNetworkSpecResponse
	GetBody() *ModifyRCInstanceNetworkSpecResponseBody
}

type ModifyRCInstanceNetworkSpecResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRCInstanceNetworkSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRCInstanceNetworkSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceNetworkSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceNetworkSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRCInstanceNetworkSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRCInstanceNetworkSpecResponse) GetBody() *ModifyRCInstanceNetworkSpecResponseBody {
	return s.Body
}

func (s *ModifyRCInstanceNetworkSpecResponse) SetHeaders(v map[string]*string) *ModifyRCInstanceNetworkSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyRCInstanceNetworkSpecResponse) SetStatusCode(v int32) *ModifyRCInstanceNetworkSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRCInstanceNetworkSpecResponse) SetBody(v *ModifyRCInstanceNetworkSpecResponseBody) *ModifyRCInstanceNetworkSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyRCInstanceNetworkSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
