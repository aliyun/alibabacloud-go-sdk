// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClickhouseEngineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClickhouseEngineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClickhouseEngineResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClickhouseEngineResponseBody) *ModifyClickhouseEngineResponse
	GetBody() *ModifyClickhouseEngineResponseBody
}

type ModifyClickhouseEngineResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClickhouseEngineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClickhouseEngineResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClickhouseEngineResponse) GoString() string {
	return s.String()
}

func (s *ModifyClickhouseEngineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClickhouseEngineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClickhouseEngineResponse) GetBody() *ModifyClickhouseEngineResponseBody {
	return s.Body
}

func (s *ModifyClickhouseEngineResponse) SetHeaders(v map[string]*string) *ModifyClickhouseEngineResponse {
	s.Headers = v
	return s
}

func (s *ModifyClickhouseEngineResponse) SetStatusCode(v int32) *ModifyClickhouseEngineResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClickhouseEngineResponse) SetBody(v *ModifyClickhouseEngineResponseBody) *ModifyClickhouseEngineResponse {
	s.Body = v
	return s
}

func (s *ModifyClickhouseEngineResponse) Validate() error {
	return dara.Validate(s)
}
