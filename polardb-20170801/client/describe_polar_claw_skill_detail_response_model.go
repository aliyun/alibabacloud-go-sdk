// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawSkillDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarClawSkillDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarClawSkillDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarClawSkillDetailResponseBody) *DescribePolarClawSkillDetailResponse
	GetBody() *DescribePolarClawSkillDetailResponseBody
}

type DescribePolarClawSkillDetailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarClawSkillDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarClawSkillDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarClawSkillDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarClawSkillDetailResponse) GetBody() *DescribePolarClawSkillDetailResponseBody {
	return s.Body
}

func (s *DescribePolarClawSkillDetailResponse) SetHeaders(v map[string]*string) *DescribePolarClawSkillDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarClawSkillDetailResponse) SetStatusCode(v int32) *DescribePolarClawSkillDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponse) SetBody(v *DescribePolarClawSkillDetailResponseBody) *DescribePolarClawSkillDetailResponse {
	s.Body = v
	return s
}

func (s *DescribePolarClawSkillDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
