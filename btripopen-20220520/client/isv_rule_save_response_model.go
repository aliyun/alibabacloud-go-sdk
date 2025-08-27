// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsvRuleSaveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IsvRuleSaveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IsvRuleSaveResponse
	GetStatusCode() *int32
	SetBody(v *IsvRuleSaveResponseBody) *IsvRuleSaveResponse
	GetBody() *IsvRuleSaveResponseBody
}

type IsvRuleSaveResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IsvRuleSaveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IsvRuleSaveResponse) String() string {
	return dara.Prettify(s)
}

func (s IsvRuleSaveResponse) GoString() string {
	return s.String()
}

func (s *IsvRuleSaveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IsvRuleSaveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IsvRuleSaveResponse) GetBody() *IsvRuleSaveResponseBody {
	return s.Body
}

func (s *IsvRuleSaveResponse) SetHeaders(v map[string]*string) *IsvRuleSaveResponse {
	s.Headers = v
	return s
}

func (s *IsvRuleSaveResponse) SetStatusCode(v int32) *IsvRuleSaveResponse {
	s.StatusCode = &v
	return s
}

func (s *IsvRuleSaveResponse) SetBody(v *IsvRuleSaveResponseBody) *IsvRuleSaveResponse {
	s.Body = v
	return s
}

func (s *IsvRuleSaveResponse) Validate() error {
	return dara.Validate(s)
}
