// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterVipResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterVipResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterVipResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterVipResponseBody) *ModifyDBClusterVipResponse
	GetBody() *ModifyDBClusterVipResponseBody
}

type ModifyDBClusterVipResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterVipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterVipResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterVipResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterVipResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterVipResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterVipResponse) GetBody() *ModifyDBClusterVipResponseBody {
	return s.Body
}

func (s *ModifyDBClusterVipResponse) SetHeaders(v map[string]*string) *ModifyDBClusterVipResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterVipResponse) SetStatusCode(v int32) *ModifyDBClusterVipResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterVipResponse) SetBody(v *ModifyDBClusterVipResponseBody) *ModifyDBClusterVipResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterVipResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
