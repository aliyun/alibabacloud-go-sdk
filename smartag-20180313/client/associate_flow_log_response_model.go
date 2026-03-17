// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateFlowLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateFlowLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateFlowLogResponse
	GetStatusCode() *int32
	SetBody(v *AssociateFlowLogResponseBody) *AssociateFlowLogResponse
	GetBody() *AssociateFlowLogResponseBody
}

type AssociateFlowLogResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateFlowLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateFlowLogResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateFlowLogResponse) GoString() string {
	return s.String()
}

func (s *AssociateFlowLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateFlowLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateFlowLogResponse) GetBody() *AssociateFlowLogResponseBody {
	return s.Body
}

func (s *AssociateFlowLogResponse) SetHeaders(v map[string]*string) *AssociateFlowLogResponse {
	s.Headers = v
	return s
}

func (s *AssociateFlowLogResponse) SetStatusCode(v int32) *AssociateFlowLogResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateFlowLogResponse) SetBody(v *AssociateFlowLogResponseBody) *AssociateFlowLogResponse {
	s.Body = v
	return s
}

func (s *AssociateFlowLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
