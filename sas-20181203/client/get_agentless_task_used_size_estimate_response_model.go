// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentlessTaskUsedSizeEstimateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentlessTaskUsedSizeEstimateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentlessTaskUsedSizeEstimateResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentlessTaskUsedSizeEstimateResponseBody) *GetAgentlessTaskUsedSizeEstimateResponse
	GetBody() *GetAgentlessTaskUsedSizeEstimateResponseBody
}

type GetAgentlessTaskUsedSizeEstimateResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentlessTaskUsedSizeEstimateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentlessTaskUsedSizeEstimateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentlessTaskUsedSizeEstimateResponse) GoString() string {
	return s.String()
}

func (s *GetAgentlessTaskUsedSizeEstimateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentlessTaskUsedSizeEstimateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentlessTaskUsedSizeEstimateResponse) GetBody() *GetAgentlessTaskUsedSizeEstimateResponseBody {
	return s.Body
}

func (s *GetAgentlessTaskUsedSizeEstimateResponse) SetHeaders(v map[string]*string) *GetAgentlessTaskUsedSizeEstimateResponse {
	s.Headers = v
	return s
}

func (s *GetAgentlessTaskUsedSizeEstimateResponse) SetStatusCode(v int32) *GetAgentlessTaskUsedSizeEstimateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentlessTaskUsedSizeEstimateResponse) SetBody(v *GetAgentlessTaskUsedSizeEstimateResponseBody) *GetAgentlessTaskUsedSizeEstimateResponse {
	s.Body = v
	return s
}

func (s *GetAgentlessTaskUsedSizeEstimateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
