// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLakeFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataLakeFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataLakeFunctionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataLakeFunctionResponseBody) *DeleteDataLakeFunctionResponse
	GetBody() *DeleteDataLakeFunctionResponseBody
}

type DeleteDataLakeFunctionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataLakeFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataLakeFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLakeFunctionResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataLakeFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataLakeFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataLakeFunctionResponse) GetBody() *DeleteDataLakeFunctionResponseBody {
	return s.Body
}

func (s *DeleteDataLakeFunctionResponse) SetHeaders(v map[string]*string) *DeleteDataLakeFunctionResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataLakeFunctionResponse) SetStatusCode(v int32) *DeleteDataLakeFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataLakeFunctionResponse) SetBody(v *DeleteDataLakeFunctionResponseBody) *DeleteDataLakeFunctionResponse {
	s.Body = v
	return s
}

func (s *DeleteDataLakeFunctionResponse) Validate() error {
	return dara.Validate(s)
}
