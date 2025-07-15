// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDistributeImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DistributeImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DistributeImageResponse
	GetStatusCode() *int32
	SetBody(v *DistributeImageResponseBody) *DistributeImageResponse
	GetBody() *DistributeImageResponseBody
}

type DistributeImageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DistributeImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DistributeImageResponse) String() string {
	return dara.Prettify(s)
}

func (s DistributeImageResponse) GoString() string {
	return s.String()
}

func (s *DistributeImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DistributeImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DistributeImageResponse) GetBody() *DistributeImageResponseBody {
	return s.Body
}

func (s *DistributeImageResponse) SetHeaders(v map[string]*string) *DistributeImageResponse {
	s.Headers = v
	return s
}

func (s *DistributeImageResponse) SetStatusCode(v int32) *DistributeImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DistributeImageResponse) SetBody(v *DistributeImageResponseBody) *DistributeImageResponse {
	s.Body = v
	return s
}

func (s *DistributeImageResponse) Validate() error {
	return dara.Validate(s)
}
