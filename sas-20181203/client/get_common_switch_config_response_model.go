// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommonSwitchConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCommonSwitchConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCommonSwitchConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetCommonSwitchConfigResponseBody) *GetCommonSwitchConfigResponse
	GetBody() *GetCommonSwitchConfigResponseBody
}

type GetCommonSwitchConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCommonSwitchConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCommonSwitchConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCommonSwitchConfigResponse) GoString() string {
	return s.String()
}

func (s *GetCommonSwitchConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCommonSwitchConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCommonSwitchConfigResponse) GetBody() *GetCommonSwitchConfigResponseBody {
	return s.Body
}

func (s *GetCommonSwitchConfigResponse) SetHeaders(v map[string]*string) *GetCommonSwitchConfigResponse {
	s.Headers = v
	return s
}

func (s *GetCommonSwitchConfigResponse) SetStatusCode(v int32) *GetCommonSwitchConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCommonSwitchConfigResponse) SetBody(v *GetCommonSwitchConfigResponseBody) *GetCommonSwitchConfigResponse {
	s.Body = v
	return s
}

func (s *GetCommonSwitchConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
