// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppLayoutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppLayoutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppLayoutResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppLayoutResponseBody) *ModifyAppLayoutResponse
	GetBody() *ModifyAppLayoutResponseBody
}

type ModifyAppLayoutResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppLayoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppLayoutResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppLayoutResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppLayoutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppLayoutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppLayoutResponse) GetBody() *ModifyAppLayoutResponseBody {
	return s.Body
}

func (s *ModifyAppLayoutResponse) SetHeaders(v map[string]*string) *ModifyAppLayoutResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppLayoutResponse) SetStatusCode(v int32) *ModifyAppLayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppLayoutResponse) SetBody(v *ModifyAppLayoutResponseBody) *ModifyAppLayoutResponse {
	s.Body = v
	return s
}

func (s *ModifyAppLayoutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
