// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomCallTaggingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCustomCallTaggingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCustomCallTaggingResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCustomCallTaggingResponseBody) *ModifyCustomCallTaggingResponse
	GetBody() *ModifyCustomCallTaggingResponseBody
}

type ModifyCustomCallTaggingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCustomCallTaggingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCustomCallTaggingResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomCallTaggingResponse) GoString() string {
	return s.String()
}

func (s *ModifyCustomCallTaggingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCustomCallTaggingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCustomCallTaggingResponse) GetBody() *ModifyCustomCallTaggingResponseBody {
	return s.Body
}

func (s *ModifyCustomCallTaggingResponse) SetHeaders(v map[string]*string) *ModifyCustomCallTaggingResponse {
	s.Headers = v
	return s
}

func (s *ModifyCustomCallTaggingResponse) SetStatusCode(v int32) *ModifyCustomCallTaggingResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCustomCallTaggingResponse) SetBody(v *ModifyCustomCallTaggingResponseBody) *ModifyCustomCallTaggingResponse {
	s.Body = v
	return s
}

func (s *ModifyCustomCallTaggingResponse) Validate() error {
	return dara.Validate(s)
}
