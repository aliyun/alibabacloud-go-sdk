// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSumRecordFlowPopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSumRecordFlowPopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSumRecordFlowPopResponse
	GetStatusCode() *int32
	SetBody(v *AddSumRecordFlowPopResponseBody) *AddSumRecordFlowPopResponse
	GetBody() *AddSumRecordFlowPopResponseBody
}

type AddSumRecordFlowPopResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSumRecordFlowPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSumRecordFlowPopResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSumRecordFlowPopResponse) GoString() string {
	return s.String()
}

func (s *AddSumRecordFlowPopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSumRecordFlowPopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSumRecordFlowPopResponse) GetBody() *AddSumRecordFlowPopResponseBody {
	return s.Body
}

func (s *AddSumRecordFlowPopResponse) SetHeaders(v map[string]*string) *AddSumRecordFlowPopResponse {
	s.Headers = v
	return s
}

func (s *AddSumRecordFlowPopResponse) SetStatusCode(v int32) *AddSumRecordFlowPopResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSumRecordFlowPopResponse) SetBody(v *AddSumRecordFlowPopResponseBody) *AddSumRecordFlowPopResponse {
	s.Body = v
	return s
}

func (s *AddSumRecordFlowPopResponse) Validate() error {
	return dara.Validate(s)
}
