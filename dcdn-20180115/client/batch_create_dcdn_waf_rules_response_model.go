// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateDcdnWafRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCreateDcdnWafRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCreateDcdnWafRulesResponse
	GetStatusCode() *int32
	SetBody(v *BatchCreateDcdnWafRulesResponseBody) *BatchCreateDcdnWafRulesResponse
	GetBody() *BatchCreateDcdnWafRulesResponseBody
}

type BatchCreateDcdnWafRulesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateDcdnWafRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateDcdnWafRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateDcdnWafRulesResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateDcdnWafRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCreateDcdnWafRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCreateDcdnWafRulesResponse) GetBody() *BatchCreateDcdnWafRulesResponseBody {
	return s.Body
}

func (s *BatchCreateDcdnWafRulesResponse) SetHeaders(v map[string]*string) *BatchCreateDcdnWafRulesResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateDcdnWafRulesResponse) SetStatusCode(v int32) *BatchCreateDcdnWafRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateDcdnWafRulesResponse) SetBody(v *BatchCreateDcdnWafRulesResponseBody) *BatchCreateDcdnWafRulesResponse {
	s.Body = v
	return s
}

func (s *BatchCreateDcdnWafRulesResponse) Validate() error {
	return dara.Validate(s)
}
