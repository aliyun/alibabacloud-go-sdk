// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDataLevelPermissionRuleConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDataLevelPermissionRuleConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDataLevelPermissionRuleConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetDataLevelPermissionRuleConfigResponseBody) *SetDataLevelPermissionRuleConfigResponse
	GetBody() *SetDataLevelPermissionRuleConfigResponseBody
}

type SetDataLevelPermissionRuleConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDataLevelPermissionRuleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDataLevelPermissionRuleConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDataLevelPermissionRuleConfigResponse) GoString() string {
	return s.String()
}

func (s *SetDataLevelPermissionRuleConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDataLevelPermissionRuleConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDataLevelPermissionRuleConfigResponse) GetBody() *SetDataLevelPermissionRuleConfigResponseBody {
	return s.Body
}

func (s *SetDataLevelPermissionRuleConfigResponse) SetHeaders(v map[string]*string) *SetDataLevelPermissionRuleConfigResponse {
	s.Headers = v
	return s
}

func (s *SetDataLevelPermissionRuleConfigResponse) SetStatusCode(v int32) *SetDataLevelPermissionRuleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDataLevelPermissionRuleConfigResponse) SetBody(v *SetDataLevelPermissionRuleConfigResponseBody) *SetDataLevelPermissionRuleConfigResponse {
	s.Body = v
	return s
}

func (s *SetDataLevelPermissionRuleConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
