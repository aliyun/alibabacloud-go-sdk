// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterPayTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterPayTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterPayTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterPayTypeResponseBody) *ModifyDBClusterPayTypeResponse
	GetBody() *ModifyDBClusterPayTypeResponseBody
}

type ModifyDBClusterPayTypeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterPayTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterPayTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterPayTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterPayTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterPayTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterPayTypeResponse) GetBody() *ModifyDBClusterPayTypeResponseBody {
	return s.Body
}

func (s *ModifyDBClusterPayTypeResponse) SetHeaders(v map[string]*string) *ModifyDBClusterPayTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterPayTypeResponse) SetStatusCode(v int32) *ModifyDBClusterPayTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterPayTypeResponse) SetBody(v *ModifyDBClusterPayTypeResponseBody) *ModifyDBClusterPayTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterPayTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
