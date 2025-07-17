// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWorkflowExtraInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetWorkflowExtraInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetWorkflowExtraInfoResponse
	GetStatusCode() *int32
	SetBody(v *SetWorkflowExtraInfoResponseBody) *SetWorkflowExtraInfoResponse
	GetBody() *SetWorkflowExtraInfoResponseBody
}

type SetWorkflowExtraInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetWorkflowExtraInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetWorkflowExtraInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s SetWorkflowExtraInfoResponse) GoString() string {
	return s.String()
}

func (s *SetWorkflowExtraInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetWorkflowExtraInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetWorkflowExtraInfoResponse) GetBody() *SetWorkflowExtraInfoResponseBody {
	return s.Body
}

func (s *SetWorkflowExtraInfoResponse) SetHeaders(v map[string]*string) *SetWorkflowExtraInfoResponse {
	s.Headers = v
	return s
}

func (s *SetWorkflowExtraInfoResponse) SetStatusCode(v int32) *SetWorkflowExtraInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SetWorkflowExtraInfoResponse) SetBody(v *SetWorkflowExtraInfoResponseBody) *SetWorkflowExtraInfoResponse {
	s.Body = v
	return s
}

func (s *SetWorkflowExtraInfoResponse) Validate() error {
	return dara.Validate(s)
}
