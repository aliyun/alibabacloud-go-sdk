// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateWafRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCreateWafRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCreateWafRulesResponse
	GetStatusCode() *int32
	SetBody(v *BatchCreateWafRulesResponseBody) *BatchCreateWafRulesResponse
	GetBody() *BatchCreateWafRulesResponseBody
}

type BatchCreateWafRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateWafRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateWafRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateWafRulesResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateWafRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCreateWafRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCreateWafRulesResponse) GetBody() *BatchCreateWafRulesResponseBody {
	return s.Body
}

func (s *BatchCreateWafRulesResponse) SetHeaders(v map[string]*string) *BatchCreateWafRulesResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateWafRulesResponse) SetStatusCode(v int32) *BatchCreateWafRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateWafRulesResponse) SetBody(v *BatchCreateWafRulesResponseBody) *BatchCreateWafRulesResponse {
	s.Body = v
	return s
}

func (s *BatchCreateWafRulesResponse) Validate() error {
	return dara.Validate(s)
}
