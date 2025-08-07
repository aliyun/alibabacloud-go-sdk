// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformDBClusterPayTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransformDBClusterPayTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransformDBClusterPayTypeResponse
	GetStatusCode() *int32
	SetBody(v *TransformDBClusterPayTypeResponseBody) *TransformDBClusterPayTypeResponse
	GetBody() *TransformDBClusterPayTypeResponseBody
}

type TransformDBClusterPayTypeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransformDBClusterPayTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransformDBClusterPayTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s TransformDBClusterPayTypeResponse) GoString() string {
	return s.String()
}

func (s *TransformDBClusterPayTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransformDBClusterPayTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransformDBClusterPayTypeResponse) GetBody() *TransformDBClusterPayTypeResponseBody {
	return s.Body
}

func (s *TransformDBClusterPayTypeResponse) SetHeaders(v map[string]*string) *TransformDBClusterPayTypeResponse {
	s.Headers = v
	return s
}

func (s *TransformDBClusterPayTypeResponse) SetStatusCode(v int32) *TransformDBClusterPayTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *TransformDBClusterPayTypeResponse) SetBody(v *TransformDBClusterPayTypeResponseBody) *TransformDBClusterPayTypeResponse {
	s.Body = v
	return s
}

func (s *TransformDBClusterPayTypeResponse) Validate() error {
	return dara.Validate(s)
}
