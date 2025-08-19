// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLayerVersionByArnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLayerVersionByArnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLayerVersionByArnResponse
	GetStatusCode() *int32
	SetBody(v *Layer) *GetLayerVersionByArnResponse
	GetBody() *Layer
}

type GetLayerVersionByArnResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Layer             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLayerVersionByArnResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLayerVersionByArnResponse) GoString() string {
	return s.String()
}

func (s *GetLayerVersionByArnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLayerVersionByArnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLayerVersionByArnResponse) GetBody() *Layer {
	return s.Body
}

func (s *GetLayerVersionByArnResponse) SetHeaders(v map[string]*string) *GetLayerVersionByArnResponse {
	s.Headers = v
	return s
}

func (s *GetLayerVersionByArnResponse) SetStatusCode(v int32) *GetLayerVersionByArnResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLayerVersionByArnResponse) SetBody(v *Layer) *GetLayerVersionByArnResponse {
	s.Body = v
	return s
}

func (s *GetLayerVersionByArnResponse) Validate() error {
	return dara.Validate(s)
}
