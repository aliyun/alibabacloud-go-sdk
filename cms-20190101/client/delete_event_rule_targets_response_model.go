// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventRuleTargetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEventRuleTargetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEventRuleTargetsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEventRuleTargetsResponseBody) *DeleteEventRuleTargetsResponse
	GetBody() *DeleteEventRuleTargetsResponseBody
}

type DeleteEventRuleTargetsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventRuleTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventRuleTargetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventRuleTargetsResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventRuleTargetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEventRuleTargetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEventRuleTargetsResponse) GetBody() *DeleteEventRuleTargetsResponseBody {
	return s.Body
}

func (s *DeleteEventRuleTargetsResponse) SetHeaders(v map[string]*string) *DeleteEventRuleTargetsResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventRuleTargetsResponse) SetStatusCode(v int32) *DeleteEventRuleTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventRuleTargetsResponse) SetBody(v *DeleteEventRuleTargetsResponseBody) *DeleteEventRuleTargetsResponse {
	s.Body = v
	return s
}

func (s *DeleteEventRuleTargetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
