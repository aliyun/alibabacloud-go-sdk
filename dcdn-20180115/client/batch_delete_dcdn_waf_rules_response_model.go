// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDcdnWafRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteDcdnWafRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteDcdnWafRulesResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteDcdnWafRulesResponseBody) *BatchDeleteDcdnWafRulesResponse
	GetBody() *BatchDeleteDcdnWafRulesResponseBody
}

type BatchDeleteDcdnWafRulesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteDcdnWafRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteDcdnWafRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDcdnWafRulesResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteDcdnWafRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteDcdnWafRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteDcdnWafRulesResponse) GetBody() *BatchDeleteDcdnWafRulesResponseBody {
	return s.Body
}

func (s *BatchDeleteDcdnWafRulesResponse) SetHeaders(v map[string]*string) *BatchDeleteDcdnWafRulesResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteDcdnWafRulesResponse) SetStatusCode(v int32) *BatchDeleteDcdnWafRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteDcdnWafRulesResponse) SetBody(v *BatchDeleteDcdnWafRulesResponseBody) *BatchDeleteDcdnWafRulesResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteDcdnWafRulesResponse) Validate() error {
	return dara.Validate(s)
}
