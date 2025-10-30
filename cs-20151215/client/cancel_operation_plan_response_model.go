// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOperationPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelOperationPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelOperationPlanResponse
	GetStatusCode() *int32
	SetBody(v *CancelOperationPlanResponseBody) *CancelOperationPlanResponse
	GetBody() *CancelOperationPlanResponseBody
}

type CancelOperationPlanResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelOperationPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelOperationPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelOperationPlanResponse) GoString() string {
	return s.String()
}

func (s *CancelOperationPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelOperationPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelOperationPlanResponse) GetBody() *CancelOperationPlanResponseBody {
	return s.Body
}

func (s *CancelOperationPlanResponse) SetHeaders(v map[string]*string) *CancelOperationPlanResponse {
	s.Headers = v
	return s
}

func (s *CancelOperationPlanResponse) SetStatusCode(v int32) *CancelOperationPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelOperationPlanResponse) SetBody(v *CancelOperationPlanResponseBody) *CancelOperationPlanResponse {
	s.Body = v
	return s
}

func (s *CancelOperationPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
