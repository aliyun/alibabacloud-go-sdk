// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnWafPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDcdnWafPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDcdnWafPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDcdnWafPolicyResponseBody) *DeleteDcdnWafPolicyResponse
	GetBody() *DeleteDcdnWafPolicyResponseBody
}

type DeleteDcdnWafPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDcdnWafPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDcdnWafPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnWafPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnWafPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDcdnWafPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDcdnWafPolicyResponse) GetBody() *DeleteDcdnWafPolicyResponseBody {
	return s.Body
}

func (s *DeleteDcdnWafPolicyResponse) SetHeaders(v map[string]*string) *DeleteDcdnWafPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnWafPolicyResponse) SetStatusCode(v int32) *DeleteDcdnWafPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDcdnWafPolicyResponse) SetBody(v *DeleteDcdnWafPolicyResponseBody) *DeleteDcdnWafPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteDcdnWafPolicyResponse) Validate() error {
	return dara.Validate(s)
}
