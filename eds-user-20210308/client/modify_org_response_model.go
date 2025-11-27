// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOrgResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyOrgResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyOrgResponse
	GetStatusCode() *int32
	SetBody(v *ModifyOrgResponseBody) *ModifyOrgResponse
	GetBody() *ModifyOrgResponseBody
}

type ModifyOrgResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOrgResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyOrgResponse) GoString() string {
	return s.String()
}

func (s *ModifyOrgResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyOrgResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyOrgResponse) GetBody() *ModifyOrgResponseBody {
	return s.Body
}

func (s *ModifyOrgResponse) SetHeaders(v map[string]*string) *ModifyOrgResponse {
	s.Headers = v
	return s
}

func (s *ModifyOrgResponse) SetStatusCode(v int32) *ModifyOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOrgResponse) SetBody(v *ModifyOrgResponseBody) *ModifyOrgResponse {
	s.Body = v
	return s
}

func (s *ModifyOrgResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
