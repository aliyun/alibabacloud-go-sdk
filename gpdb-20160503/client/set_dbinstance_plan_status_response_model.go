// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDBInstancePlanStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDBInstancePlanStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDBInstancePlanStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetDBInstancePlanStatusResponseBody) *SetDBInstancePlanStatusResponse
	GetBody() *SetDBInstancePlanStatusResponseBody
}

type SetDBInstancePlanStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDBInstancePlanStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDBInstancePlanStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDBInstancePlanStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDBInstancePlanStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDBInstancePlanStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDBInstancePlanStatusResponse) GetBody() *SetDBInstancePlanStatusResponseBody {
	return s.Body
}

func (s *SetDBInstancePlanStatusResponse) SetHeaders(v map[string]*string) *SetDBInstancePlanStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDBInstancePlanStatusResponse) SetStatusCode(v int32) *SetDBInstancePlanStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDBInstancePlanStatusResponse) SetBody(v *SetDBInstancePlanStatusResponseBody) *SetDBInstancePlanStatusResponse {
	s.Body = v
	return s
}

func (s *SetDBInstancePlanStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
