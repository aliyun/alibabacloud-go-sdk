// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLayerVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLayerVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLayerVersionResponse
	GetStatusCode() *int32
	SetBody(v *Layer) *GetLayerVersionResponse
	GetBody() *Layer
}

type GetLayerVersionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Layer             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLayerVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLayerVersionResponse) GoString() string {
	return s.String()
}

func (s *GetLayerVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLayerVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLayerVersionResponse) GetBody() *Layer {
	return s.Body
}

func (s *GetLayerVersionResponse) SetHeaders(v map[string]*string) *GetLayerVersionResponse {
	s.Headers = v
	return s
}

func (s *GetLayerVersionResponse) SetStatusCode(v int32) *GetLayerVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLayerVersionResponse) SetBody(v *Layer) *GetLayerVersionResponse {
	s.Body = v
	return s
}

func (s *GetLayerVersionResponse) Validate() error {
	return dara.Validate(s)
}
