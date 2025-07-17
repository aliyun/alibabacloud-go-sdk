// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceSharedRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataSourceSharedRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataSourceSharedRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataSourceSharedRuleResponseBody) *DeleteDataSourceSharedRuleResponse
	GetBody() *DeleteDataSourceSharedRuleResponseBody
}

type DeleteDataSourceSharedRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataSourceSharedRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataSourceSharedRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceSharedRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceSharedRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataSourceSharedRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataSourceSharedRuleResponse) GetBody() *DeleteDataSourceSharedRuleResponseBody {
	return s.Body
}

func (s *DeleteDataSourceSharedRuleResponse) SetHeaders(v map[string]*string) *DeleteDataSourceSharedRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSourceSharedRuleResponse) SetStatusCode(v int32) *DeleteDataSourceSharedRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataSourceSharedRuleResponse) SetBody(v *DeleteDataSourceSharedRuleResponseBody) *DeleteDataSourceSharedRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteDataSourceSharedRuleResponse) Validate() error {
	return dara.Validate(s)
}
