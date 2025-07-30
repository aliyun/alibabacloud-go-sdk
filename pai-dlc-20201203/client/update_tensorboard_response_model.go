// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTensorboardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTensorboardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTensorboardResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTensorboardResponseBody) *UpdateTensorboardResponse
	GetBody() *UpdateTensorboardResponseBody
}

type UpdateTensorboardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTensorboardResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTensorboardResponse) GoString() string {
	return s.String()
}

func (s *UpdateTensorboardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTensorboardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTensorboardResponse) GetBody() *UpdateTensorboardResponseBody {
	return s.Body
}

func (s *UpdateTensorboardResponse) SetHeaders(v map[string]*string) *UpdateTensorboardResponse {
	s.Headers = v
	return s
}

func (s *UpdateTensorboardResponse) SetStatusCode(v int32) *UpdateTensorboardResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTensorboardResponse) SetBody(v *UpdateTensorboardResponseBody) *UpdateTensorboardResponse {
	s.Body = v
	return s
}

func (s *UpdateTensorboardResponse) Validate() error {
	return dara.Validate(s)
}
