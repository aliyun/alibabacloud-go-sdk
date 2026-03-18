// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPolicyContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPolicyContentResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPolicyContentResponseBody) *ModifyPolicyContentResponse
	GetBody() *ModifyPolicyContentResponseBody
}

type ModifyPolicyContentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPolicyContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPolicyContentResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyContentResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPolicyContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPolicyContentResponse) GetBody() *ModifyPolicyContentResponseBody {
	return s.Body
}

func (s *ModifyPolicyContentResponse) SetHeaders(v map[string]*string) *ModifyPolicyContentResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolicyContentResponse) SetStatusCode(v int32) *ModifyPolicyContentResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPolicyContentResponse) SetBody(v *ModifyPolicyContentResponseBody) *ModifyPolicyContentResponse {
	s.Body = v
	return s
}

func (s *ModifyPolicyContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
