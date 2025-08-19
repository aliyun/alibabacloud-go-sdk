// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLayerVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLayerVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLayerVersionResponse
	GetStatusCode() *int32
}

type DeleteLayerVersionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteLayerVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLayerVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteLayerVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLayerVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLayerVersionResponse) SetHeaders(v map[string]*string) *DeleteLayerVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteLayerVersionResponse) SetStatusCode(v int32) *DeleteLayerVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLayerVersionResponse) Validate() error {
	return dara.Validate(s)
}
