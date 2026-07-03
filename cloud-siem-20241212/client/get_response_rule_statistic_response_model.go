// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResponseRuleStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResponseRuleStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResponseRuleStatisticResponse
	GetStatusCode() *int32
	SetBody(v *GetResponseRuleStatisticResponseBody) *GetResponseRuleStatisticResponse
	GetBody() *GetResponseRuleStatisticResponseBody
}

type GetResponseRuleStatisticResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResponseRuleStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResponseRuleStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResponseRuleStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetResponseRuleStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResponseRuleStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResponseRuleStatisticResponse) GetBody() *GetResponseRuleStatisticResponseBody {
	return s.Body
}

func (s *GetResponseRuleStatisticResponse) SetHeaders(v map[string]*string) *GetResponseRuleStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetResponseRuleStatisticResponse) SetStatusCode(v int32) *GetResponseRuleStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResponseRuleStatisticResponse) SetBody(v *GetResponseRuleStatisticResponseBody) *GetResponseRuleStatisticResponse {
	s.Body = v
	return s
}

func (s *GetResponseRuleStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
