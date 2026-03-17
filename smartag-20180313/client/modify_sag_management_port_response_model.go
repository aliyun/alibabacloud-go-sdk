// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagManagementPortResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySagManagementPortResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySagManagementPortResponse
	GetStatusCode() *int32
	SetBody(v *ModifySagManagementPortResponseBody) *ModifySagManagementPortResponse
	GetBody() *ModifySagManagementPortResponseBody
}

type ModifySagManagementPortResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySagManagementPortResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySagManagementPortResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySagManagementPortResponse) GoString() string {
	return s.String()
}

func (s *ModifySagManagementPortResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySagManagementPortResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySagManagementPortResponse) GetBody() *ModifySagManagementPortResponseBody {
	return s.Body
}

func (s *ModifySagManagementPortResponse) SetHeaders(v map[string]*string) *ModifySagManagementPortResponse {
	s.Headers = v
	return s
}

func (s *ModifySagManagementPortResponse) SetStatusCode(v int32) *ModifySagManagementPortResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySagManagementPortResponse) SetBody(v *ModifySagManagementPortResponseBody) *ModifySagManagementPortResponse {
	s.Body = v
	return s
}

func (s *ModifySagManagementPortResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
