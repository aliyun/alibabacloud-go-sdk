// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineFlowVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OfflineFlowVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OfflineFlowVersionResponse
	GetStatusCode() *int32
	SetBody(v *OfflineFlowVersionResponseBody) *OfflineFlowVersionResponse
	GetBody() *OfflineFlowVersionResponseBody
}

type OfflineFlowVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineFlowVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineFlowVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s OfflineFlowVersionResponse) GoString() string {
	return s.String()
}

func (s *OfflineFlowVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OfflineFlowVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OfflineFlowVersionResponse) GetBody() *OfflineFlowVersionResponseBody {
	return s.Body
}

func (s *OfflineFlowVersionResponse) SetHeaders(v map[string]*string) *OfflineFlowVersionResponse {
	s.Headers = v
	return s
}

func (s *OfflineFlowVersionResponse) SetStatusCode(v int32) *OfflineFlowVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineFlowVersionResponse) SetBody(v *OfflineFlowVersionResponseBody) *OfflineFlowVersionResponse {
	s.Body = v
	return s
}

func (s *OfflineFlowVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
