// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDasInstanceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDasInstanceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDasInstanceConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDasInstanceConfigResponseBody) *ModifyDasInstanceConfigResponse
	GetBody() *ModifyDasInstanceConfigResponseBody
}

type ModifyDasInstanceConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDasInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDasInstanceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDasInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDasInstanceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDasInstanceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDasInstanceConfigResponse) GetBody() *ModifyDasInstanceConfigResponseBody {
	return s.Body
}

func (s *ModifyDasInstanceConfigResponse) SetHeaders(v map[string]*string) *ModifyDasInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDasInstanceConfigResponse) SetStatusCode(v int32) *ModifyDasInstanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDasInstanceConfigResponse) SetBody(v *ModifyDasInstanceConfigResponseBody) *ModifyDasInstanceConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyDasInstanceConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
