// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneEngineConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloneEngineConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloneEngineConfigResponse
	GetStatusCode() *int32
	SetBody(v *CloneEngineConfigResponseBody) *CloneEngineConfigResponse
	GetBody() *CloneEngineConfigResponseBody
}

type CloneEngineConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneEngineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneEngineConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CloneEngineConfigResponse) GoString() string {
	return s.String()
}

func (s *CloneEngineConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloneEngineConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloneEngineConfigResponse) GetBody() *CloneEngineConfigResponseBody {
	return s.Body
}

func (s *CloneEngineConfigResponse) SetHeaders(v map[string]*string) *CloneEngineConfigResponse {
	s.Headers = v
	return s
}

func (s *CloneEngineConfigResponse) SetStatusCode(v int32) *CloneEngineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneEngineConfigResponse) SetBody(v *CloneEngineConfigResponseBody) *CloneEngineConfigResponse {
	s.Body = v
	return s
}

func (s *CloneEngineConfigResponse) Validate() error {
	return dara.Validate(s)
}
