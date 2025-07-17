// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataCorrectExecSQLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDataCorrectExecSQLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDataCorrectExecSQLResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDataCorrectExecSQLResponseBody) *ModifyDataCorrectExecSQLResponse
	GetBody() *ModifyDataCorrectExecSQLResponseBody
}

type ModifyDataCorrectExecSQLResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDataCorrectExecSQLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDataCorrectExecSQLResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataCorrectExecSQLResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataCorrectExecSQLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDataCorrectExecSQLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDataCorrectExecSQLResponse) GetBody() *ModifyDataCorrectExecSQLResponseBody {
	return s.Body
}

func (s *ModifyDataCorrectExecSQLResponse) SetHeaders(v map[string]*string) *ModifyDataCorrectExecSQLResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataCorrectExecSQLResponse) SetStatusCode(v int32) *ModifyDataCorrectExecSQLResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataCorrectExecSQLResponse) SetBody(v *ModifyDataCorrectExecSQLResponseBody) *ModifyDataCorrectExecSQLResponse {
	s.Body = v
	return s
}

func (s *ModifyDataCorrectExecSQLResponse) Validate() error {
	return dara.Validate(s)
}
