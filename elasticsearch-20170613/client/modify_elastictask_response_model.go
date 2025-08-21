// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElastictaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyElastictaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyElastictaskResponse
	GetStatusCode() *int32
	SetBody(v *ModifyElastictaskResponseBody) *ModifyElastictaskResponse
	GetBody() *ModifyElastictaskResponseBody
}

type ModifyElastictaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyElastictaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyElastictaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyElastictaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyElastictaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyElastictaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyElastictaskResponse) GetBody() *ModifyElastictaskResponseBody {
	return s.Body
}

func (s *ModifyElastictaskResponse) SetHeaders(v map[string]*string) *ModifyElastictaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyElastictaskResponse) SetStatusCode(v int32) *ModifyElastictaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyElastictaskResponse) SetBody(v *ModifyElastictaskResponseBody) *ModifyElastictaskResponse {
	s.Body = v
	return s
}

func (s *ModifyElastictaskResponse) Validate() error {
	return dara.Validate(s)
}
