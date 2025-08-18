// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateWafRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchUpdateWafRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchUpdateWafRulesResponse
	GetStatusCode() *int32
	SetBody(v *BatchUpdateWafRulesResponseBody) *BatchUpdateWafRulesResponse
	GetBody() *BatchUpdateWafRulesResponseBody
}

type BatchUpdateWafRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUpdateWafRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUpdateWafRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateWafRulesResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateWafRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchUpdateWafRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchUpdateWafRulesResponse) GetBody() *BatchUpdateWafRulesResponseBody {
	return s.Body
}

func (s *BatchUpdateWafRulesResponse) SetHeaders(v map[string]*string) *BatchUpdateWafRulesResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateWafRulesResponse) SetStatusCode(v int32) *BatchUpdateWafRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateWafRulesResponse) SetBody(v *BatchUpdateWafRulesResponseBody) *BatchUpdateWafRulesResponse {
	s.Body = v
	return s
}

func (s *BatchUpdateWafRulesResponse) Validate() error {
	return dara.Validate(s)
}
