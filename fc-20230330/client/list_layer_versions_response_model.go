// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLayerVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLayerVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLayerVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListLayerVersionOutput) *ListLayerVersionsResponse
	GetBody() *ListLayerVersionOutput
}

type ListLayerVersionsResponse struct {
	Headers    map[string]*string      `json:"headers" xml:"headers"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLayerVersionOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLayerVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLayerVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListLayerVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLayerVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLayerVersionsResponse) GetBody() *ListLayerVersionOutput {
	return s.Body
}

func (s *ListLayerVersionsResponse) SetHeaders(v map[string]*string) *ListLayerVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListLayerVersionsResponse) SetStatusCode(v int32) *ListLayerVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLayerVersionsResponse) SetBody(v *ListLayerVersionOutput) *ListLayerVersionsResponse {
	s.Body = v
	return s
}

func (s *ListLayerVersionsResponse) Validate() error {
	return dara.Validate(s)
}
