// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolarClawSkillRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPolarClawSkillRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPolarClawSkillRegionResponse
	GetStatusCode() *int32
	SetBody(v *GetPolarClawSkillRegionResponseBody) *GetPolarClawSkillRegionResponse
	GetBody() *GetPolarClawSkillRegionResponseBody
}

type GetPolarClawSkillRegionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPolarClawSkillRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPolarClawSkillRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPolarClawSkillRegionResponse) GoString() string {
	return s.String()
}

func (s *GetPolarClawSkillRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPolarClawSkillRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPolarClawSkillRegionResponse) GetBody() *GetPolarClawSkillRegionResponseBody {
	return s.Body
}

func (s *GetPolarClawSkillRegionResponse) SetHeaders(v map[string]*string) *GetPolarClawSkillRegionResponse {
	s.Headers = v
	return s
}

func (s *GetPolarClawSkillRegionResponse) SetStatusCode(v int32) *GetPolarClawSkillRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPolarClawSkillRegionResponse) SetBody(v *GetPolarClawSkillRegionResponseBody) *GetPolarClawSkillRegionResponse {
	s.Body = v
	return s
}

func (s *GetPolarClawSkillRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
