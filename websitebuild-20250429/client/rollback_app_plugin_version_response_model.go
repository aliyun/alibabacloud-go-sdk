// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackAppPluginVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackAppPluginVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackAppPluginVersionResponse
	GetStatusCode() *int32
	SetBody(v *RollbackAppPluginVersionResponseBody) *RollbackAppPluginVersionResponse
	GetBody() *RollbackAppPluginVersionResponseBody
}

type RollbackAppPluginVersionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackAppPluginVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackAppPluginVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackAppPluginVersionResponse) GoString() string {
	return s.String()
}

func (s *RollbackAppPluginVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackAppPluginVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackAppPluginVersionResponse) GetBody() *RollbackAppPluginVersionResponseBody {
	return s.Body
}

func (s *RollbackAppPluginVersionResponse) SetHeaders(v map[string]*string) *RollbackAppPluginVersionResponse {
	s.Headers = v
	return s
}

func (s *RollbackAppPluginVersionResponse) SetStatusCode(v int32) *RollbackAppPluginVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackAppPluginVersionResponse) SetBody(v *RollbackAppPluginVersionResponseBody) *RollbackAppPluginVersionResponse {
	s.Body = v
	return s
}

func (s *RollbackAppPluginVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
