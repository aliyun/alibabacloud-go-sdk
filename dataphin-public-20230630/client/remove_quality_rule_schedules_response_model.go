// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveQualityRuleSchedulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveQualityRuleSchedulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveQualityRuleSchedulesResponse
	GetStatusCode() *int32
	SetBody(v *RemoveQualityRuleSchedulesResponseBody) *RemoveQualityRuleSchedulesResponse
	GetBody() *RemoveQualityRuleSchedulesResponseBody
}

type RemoveQualityRuleSchedulesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveQualityRuleSchedulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveQualityRuleSchedulesResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveQualityRuleSchedulesResponse) GoString() string {
	return s.String()
}

func (s *RemoveQualityRuleSchedulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveQualityRuleSchedulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveQualityRuleSchedulesResponse) GetBody() *RemoveQualityRuleSchedulesResponseBody {
	return s.Body
}

func (s *RemoveQualityRuleSchedulesResponse) SetHeaders(v map[string]*string) *RemoveQualityRuleSchedulesResponse {
	s.Headers = v
	return s
}

func (s *RemoveQualityRuleSchedulesResponse) SetStatusCode(v int32) *RemoveQualityRuleSchedulesResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveQualityRuleSchedulesResponse) SetBody(v *RemoveQualityRuleSchedulesResponseBody) *RemoveQualityRuleSchedulesResponse {
	s.Body = v
	return s
}

func (s *RemoveQualityRuleSchedulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
