// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachTaskSessionResourceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAICoachTaskSessionResourceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAICoachTaskSessionResourceUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetAICoachTaskSessionResourceUsageResponseBody) *GetAICoachTaskSessionResourceUsageResponse
	GetBody() *GetAICoachTaskSessionResourceUsageResponseBody
}

type GetAICoachTaskSessionResourceUsageResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAICoachTaskSessionResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAICoachTaskSessionResourceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachTaskSessionResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *GetAICoachTaskSessionResourceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAICoachTaskSessionResourceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAICoachTaskSessionResourceUsageResponse) GetBody() *GetAICoachTaskSessionResourceUsageResponseBody {
	return s.Body
}

func (s *GetAICoachTaskSessionResourceUsageResponse) SetHeaders(v map[string]*string) *GetAICoachTaskSessionResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *GetAICoachTaskSessionResourceUsageResponse) SetStatusCode(v int32) *GetAICoachTaskSessionResourceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAICoachTaskSessionResourceUsageResponse) SetBody(v *GetAICoachTaskSessionResourceUsageResponseBody) *GetAICoachTaskSessionResourceUsageResponse {
	s.Body = v
	return s
}

func (s *GetAICoachTaskSessionResourceUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
