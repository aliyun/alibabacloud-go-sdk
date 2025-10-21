// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGenerateResourcePlanResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGenerateResourcePlanResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGenerateResourcePlanResultResponse
	GetStatusCode() *int32
	SetBody(v *GetGenerateResourcePlanResultResponseBody) *GetGenerateResourcePlanResultResponse
	GetBody() *GetGenerateResourcePlanResultResponseBody
}

type GetGenerateResourcePlanResultResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGenerateResourcePlanResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGenerateResourcePlanResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGenerateResourcePlanResultResponse) GoString() string {
	return s.String()
}

func (s *GetGenerateResourcePlanResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGenerateResourcePlanResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGenerateResourcePlanResultResponse) GetBody() *GetGenerateResourcePlanResultResponseBody {
	return s.Body
}

func (s *GetGenerateResourcePlanResultResponse) SetHeaders(v map[string]*string) *GetGenerateResourcePlanResultResponse {
	s.Headers = v
	return s
}

func (s *GetGenerateResourcePlanResultResponse) SetStatusCode(v int32) *GetGenerateResourcePlanResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGenerateResourcePlanResultResponse) SetBody(v *GetGenerateResourcePlanResultResponseBody) *GetGenerateResourcePlanResultResponse {
	s.Body = v
	return s
}

func (s *GetGenerateResourcePlanResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
