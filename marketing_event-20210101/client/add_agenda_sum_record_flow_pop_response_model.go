// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAgendaSumRecordFlowPopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAgendaSumRecordFlowPopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAgendaSumRecordFlowPopResponse
	GetStatusCode() *int32
	SetBody(v *AddAgendaSumRecordFlowPopResponseBody) *AddAgendaSumRecordFlowPopResponse
	GetBody() *AddAgendaSumRecordFlowPopResponseBody
}

type AddAgendaSumRecordFlowPopResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAgendaSumRecordFlowPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAgendaSumRecordFlowPopResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAgendaSumRecordFlowPopResponse) GoString() string {
	return s.String()
}

func (s *AddAgendaSumRecordFlowPopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAgendaSumRecordFlowPopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAgendaSumRecordFlowPopResponse) GetBody() *AddAgendaSumRecordFlowPopResponseBody {
	return s.Body
}

func (s *AddAgendaSumRecordFlowPopResponse) SetHeaders(v map[string]*string) *AddAgendaSumRecordFlowPopResponse {
	s.Headers = v
	return s
}

func (s *AddAgendaSumRecordFlowPopResponse) SetStatusCode(v int32) *AddAgendaSumRecordFlowPopResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAgendaSumRecordFlowPopResponse) SetBody(v *AddAgendaSumRecordFlowPopResponseBody) *AddAgendaSumRecordFlowPopResponse {
	s.Body = v
	return s
}

func (s *AddAgendaSumRecordFlowPopResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
