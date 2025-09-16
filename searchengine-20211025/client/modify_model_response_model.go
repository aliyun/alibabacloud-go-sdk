// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyModelResponse
	GetStatusCode() *int32
	SetBody(v *ModifyModelResponseBody) *ModifyModelResponse
	GetBody() *ModifyModelResponseBody
}

type ModifyModelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyModelResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelResponse) GoString() string {
	return s.String()
}

func (s *ModifyModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyModelResponse) GetBody() *ModifyModelResponseBody {
	return s.Body
}

func (s *ModifyModelResponse) SetHeaders(v map[string]*string) *ModifyModelResponse {
	s.Headers = v
	return s
}

func (s *ModifyModelResponse) SetStatusCode(v int32) *ModifyModelResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyModelResponse) SetBody(v *ModifyModelResponseBody) *ModifyModelResponse {
	s.Body = v
	return s
}

func (s *ModifyModelResponse) Validate() error {
	return dara.Validate(s)
}
