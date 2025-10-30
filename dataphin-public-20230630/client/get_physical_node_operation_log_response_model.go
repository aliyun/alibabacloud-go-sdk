// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalNodeOperationLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPhysicalNodeOperationLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPhysicalNodeOperationLogResponse
	GetStatusCode() *int32
	SetBody(v *GetPhysicalNodeOperationLogResponseBody) *GetPhysicalNodeOperationLogResponse
	GetBody() *GetPhysicalNodeOperationLogResponseBody
}

type GetPhysicalNodeOperationLogResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhysicalNodeOperationLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhysicalNodeOperationLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeOperationLogResponse) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeOperationLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPhysicalNodeOperationLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPhysicalNodeOperationLogResponse) GetBody() *GetPhysicalNodeOperationLogResponseBody {
	return s.Body
}

func (s *GetPhysicalNodeOperationLogResponse) SetHeaders(v map[string]*string) *GetPhysicalNodeOperationLogResponse {
	s.Headers = v
	return s
}

func (s *GetPhysicalNodeOperationLogResponse) SetStatusCode(v int32) *GetPhysicalNodeOperationLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponse) SetBody(v *GetPhysicalNodeOperationLogResponseBody) *GetPhysicalNodeOperationLogResponse {
	s.Body = v
	return s
}

func (s *GetPhysicalNodeOperationLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
