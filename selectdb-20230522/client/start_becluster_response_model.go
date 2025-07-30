// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBEClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartBEClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartBEClusterResponse
	GetStatusCode() *int32
	SetBody(v *StartBEClusterResponseBody) *StartBEClusterResponse
	GetBody() *StartBEClusterResponseBody
}

type StartBEClusterResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartBEClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartBEClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s StartBEClusterResponse) GoString() string {
	return s.String()
}

func (s *StartBEClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartBEClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartBEClusterResponse) GetBody() *StartBEClusterResponseBody {
	return s.Body
}

func (s *StartBEClusterResponse) SetHeaders(v map[string]*string) *StartBEClusterResponse {
	s.Headers = v
	return s
}

func (s *StartBEClusterResponse) SetStatusCode(v int32) *StartBEClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *StartBEClusterResponse) SetBody(v *StartBEClusterResponseBody) *StartBEClusterResponse {
	s.Body = v
	return s
}

func (s *StartBEClusterResponse) Validate() error {
	return dara.Validate(s)
}
