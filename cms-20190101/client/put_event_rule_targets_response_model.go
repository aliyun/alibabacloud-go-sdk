// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEventRuleTargetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutEventRuleTargetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutEventRuleTargetsResponse
	GetStatusCode() *int32
	SetBody(v *PutEventRuleTargetsResponseBody) *PutEventRuleTargetsResponse
	GetBody() *PutEventRuleTargetsResponseBody
}

type PutEventRuleTargetsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutEventRuleTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutEventRuleTargetsResponse) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleTargetsResponse) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutEventRuleTargetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutEventRuleTargetsResponse) GetBody() *PutEventRuleTargetsResponseBody {
	return s.Body
}

func (s *PutEventRuleTargetsResponse) SetHeaders(v map[string]*string) *PutEventRuleTargetsResponse {
	s.Headers = v
	return s
}

func (s *PutEventRuleTargetsResponse) SetStatusCode(v int32) *PutEventRuleTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *PutEventRuleTargetsResponse) SetBody(v *PutEventRuleTargetsResponseBody) *PutEventRuleTargetsResponse {
	s.Body = v
	return s
}

func (s *PutEventRuleTargetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
