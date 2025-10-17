// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterArchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterArchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterArchResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterArchResponseBody) *ModifyDBClusterArchResponse
	GetBody() *ModifyDBClusterArchResponseBody
}

type ModifyDBClusterArchResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterArchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterArchResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterArchResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterArchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterArchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterArchResponse) GetBody() *ModifyDBClusterArchResponseBody {
	return s.Body
}

func (s *ModifyDBClusterArchResponse) SetHeaders(v map[string]*string) *ModifyDBClusterArchResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterArchResponse) SetStatusCode(v int32) *ModifyDBClusterArchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterArchResponse) SetBody(v *ModifyDBClusterArchResponseBody) *ModifyDBClusterArchResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterArchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
