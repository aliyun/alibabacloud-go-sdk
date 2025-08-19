// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLayerVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLayerVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLayerVersionResponse
	GetStatusCode() *int32
	SetBody(v *Layer) *CreateLayerVersionResponse
	GetBody() *Layer
}

type CreateLayerVersionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Layer             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLayerVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLayerVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateLayerVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLayerVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLayerVersionResponse) GetBody() *Layer {
	return s.Body
}

func (s *CreateLayerVersionResponse) SetHeaders(v map[string]*string) *CreateLayerVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateLayerVersionResponse) SetStatusCode(v int32) *CreateLayerVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLayerVersionResponse) SetBody(v *Layer) *CreateLayerVersionResponse {
	s.Body = v
	return s
}

func (s *CreateLayerVersionResponse) Validate() error {
	return dara.Validate(s)
}
