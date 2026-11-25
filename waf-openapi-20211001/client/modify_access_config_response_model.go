// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccessConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAccessConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAccessConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAccessConfigResponseBody) *ModifyAccessConfigResponse
	GetBody() *ModifyAccessConfigResponseBody
}

type ModifyAccessConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccessConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccessConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccessConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccessConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAccessConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAccessConfigResponse) GetBody() *ModifyAccessConfigResponseBody {
	return s.Body
}

func (s *ModifyAccessConfigResponse) SetHeaders(v map[string]*string) *ModifyAccessConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccessConfigResponse) SetStatusCode(v int32) *ModifyAccessConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccessConfigResponse) SetBody(v *ModifyAccessConfigResponseBody) *ModifyAccessConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyAccessConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
