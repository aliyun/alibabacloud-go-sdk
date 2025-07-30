// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTensorboardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTensorboardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTensorboardResponse
	GetStatusCode() *int32
	SetBody(v *Tensorboard) *GetTensorboardResponse
	GetBody() *Tensorboard
}

type GetTensorboardResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Tensorboard       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTensorboardResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTensorboardResponse) GoString() string {
	return s.String()
}

func (s *GetTensorboardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTensorboardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTensorboardResponse) GetBody() *Tensorboard {
	return s.Body
}

func (s *GetTensorboardResponse) SetHeaders(v map[string]*string) *GetTensorboardResponse {
	s.Headers = v
	return s
}

func (s *GetTensorboardResponse) SetStatusCode(v int32) *GetTensorboardResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTensorboardResponse) SetBody(v *Tensorboard) *GetTensorboardResponse {
	s.Body = v
	return s
}

func (s *GetTensorboardResponse) Validate() error {
	return dara.Validate(s)
}
