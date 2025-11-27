// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPGHbaConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPGHbaConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPGHbaConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPGHbaConfigResponseBody) *ModifyPGHbaConfigResponse
	GetBody() *ModifyPGHbaConfigResponseBody
}

type ModifyPGHbaConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPGHbaConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPGHbaConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPGHbaConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyPGHbaConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPGHbaConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPGHbaConfigResponse) GetBody() *ModifyPGHbaConfigResponseBody {
	return s.Body
}

func (s *ModifyPGHbaConfigResponse) SetHeaders(v map[string]*string) *ModifyPGHbaConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyPGHbaConfigResponse) SetStatusCode(v int32) *ModifyPGHbaConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPGHbaConfigResponse) SetBody(v *ModifyPGHbaConfigResponseBody) *ModifyPGHbaConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyPGHbaConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
