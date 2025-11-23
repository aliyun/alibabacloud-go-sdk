// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReRunTaskFlowInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReRunTaskFlowInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReRunTaskFlowInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReRunTaskFlowInstanceResponseBody) *ReRunTaskFlowInstanceResponse
	GetBody() *ReRunTaskFlowInstanceResponseBody
}

type ReRunTaskFlowInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReRunTaskFlowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReRunTaskFlowInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReRunTaskFlowInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReRunTaskFlowInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReRunTaskFlowInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReRunTaskFlowInstanceResponse) GetBody() *ReRunTaskFlowInstanceResponseBody {
	return s.Body
}

func (s *ReRunTaskFlowInstanceResponse) SetHeaders(v map[string]*string) *ReRunTaskFlowInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReRunTaskFlowInstanceResponse) SetStatusCode(v int32) *ReRunTaskFlowInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReRunTaskFlowInstanceResponse) SetBody(v *ReRunTaskFlowInstanceResponseBody) *ReRunTaskFlowInstanceResponse {
	s.Body = v
	return s
}

func (s *ReRunTaskFlowInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
