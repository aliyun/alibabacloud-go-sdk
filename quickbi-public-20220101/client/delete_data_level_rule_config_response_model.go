// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLevelRuleConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataLevelRuleConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataLevelRuleConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataLevelRuleConfigResponseBody) *DeleteDataLevelRuleConfigResponse
	GetBody() *DeleteDataLevelRuleConfigResponseBody
}

type DeleteDataLevelRuleConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataLevelRuleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataLevelRuleConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLevelRuleConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataLevelRuleConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataLevelRuleConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataLevelRuleConfigResponse) GetBody() *DeleteDataLevelRuleConfigResponseBody {
	return s.Body
}

func (s *DeleteDataLevelRuleConfigResponse) SetHeaders(v map[string]*string) *DeleteDataLevelRuleConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataLevelRuleConfigResponse) SetStatusCode(v int32) *DeleteDataLevelRuleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataLevelRuleConfigResponse) SetBody(v *DeleteDataLevelRuleConfigResponseBody) *DeleteDataLevelRuleConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteDataLevelRuleConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
