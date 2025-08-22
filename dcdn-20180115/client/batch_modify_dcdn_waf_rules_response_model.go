// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchModifyDcdnWafRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchModifyDcdnWafRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchModifyDcdnWafRulesResponse
	GetStatusCode() *int32
	SetBody(v *BatchModifyDcdnWafRulesResponseBody) *BatchModifyDcdnWafRulesResponse
	GetBody() *BatchModifyDcdnWafRulesResponseBody
}

type BatchModifyDcdnWafRulesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchModifyDcdnWafRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchModifyDcdnWafRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchModifyDcdnWafRulesResponse) GoString() string {
	return s.String()
}

func (s *BatchModifyDcdnWafRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchModifyDcdnWafRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchModifyDcdnWafRulesResponse) GetBody() *BatchModifyDcdnWafRulesResponseBody {
	return s.Body
}

func (s *BatchModifyDcdnWafRulesResponse) SetHeaders(v map[string]*string) *BatchModifyDcdnWafRulesResponse {
	s.Headers = v
	return s
}

func (s *BatchModifyDcdnWafRulesResponse) SetStatusCode(v int32) *BatchModifyDcdnWafRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchModifyDcdnWafRulesResponse) SetBody(v *BatchModifyDcdnWafRulesResponseBody) *BatchModifyDcdnWafRulesResponse {
	s.Body = v
	return s
}

func (s *BatchModifyDcdnWafRulesResponse) Validate() error {
	return dara.Validate(s)
}
