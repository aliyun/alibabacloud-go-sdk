// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMem0SecurityIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMem0SecurityIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMem0SecurityIpsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMem0SecurityIpsResponseBody) *ModifyMem0SecurityIpsResponse
	GetBody() *ModifyMem0SecurityIpsResponseBody
}

type ModifyMem0SecurityIpsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMem0SecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMem0SecurityIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMem0SecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *ModifyMem0SecurityIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMem0SecurityIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMem0SecurityIpsResponse) GetBody() *ModifyMem0SecurityIpsResponseBody {
	return s.Body
}

func (s *ModifyMem0SecurityIpsResponse) SetHeaders(v map[string]*string) *ModifyMem0SecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *ModifyMem0SecurityIpsResponse) SetStatusCode(v int32) *ModifyMem0SecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMem0SecurityIpsResponse) SetBody(v *ModifyMem0SecurityIpsResponseBody) *ModifyMem0SecurityIpsResponse {
	s.Body = v
	return s
}

func (s *ModifyMem0SecurityIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
