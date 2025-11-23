// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineTaskFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OfflineTaskFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OfflineTaskFlowResponse
	GetStatusCode() *int32
	SetBody(v *OfflineTaskFlowResponseBody) *OfflineTaskFlowResponse
	GetBody() *OfflineTaskFlowResponseBody
}

type OfflineTaskFlowResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineTaskFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineTaskFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s OfflineTaskFlowResponse) GoString() string {
	return s.String()
}

func (s *OfflineTaskFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OfflineTaskFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OfflineTaskFlowResponse) GetBody() *OfflineTaskFlowResponseBody {
	return s.Body
}

func (s *OfflineTaskFlowResponse) SetHeaders(v map[string]*string) *OfflineTaskFlowResponse {
	s.Headers = v
	return s
}

func (s *OfflineTaskFlowResponse) SetStatusCode(v int32) *OfflineTaskFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineTaskFlowResponse) SetBody(v *OfflineTaskFlowResponseBody) *OfflineTaskFlowResponse {
	s.Body = v
	return s
}

func (s *OfflineTaskFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
