// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DebugModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DebugModelResponse
	GetStatusCode() *int32
	SetBody(v *DebugModelResponseBody) *DebugModelResponse
	GetBody() *DebugModelResponseBody
}

type DebugModelResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DebugModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DebugModelResponse) String() string {
	return dara.Prettify(s)
}

func (s DebugModelResponse) GoString() string {
	return s.String()
}

func (s *DebugModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DebugModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DebugModelResponse) GetBody() *DebugModelResponseBody {
	return s.Body
}

func (s *DebugModelResponse) SetHeaders(v map[string]*string) *DebugModelResponse {
	s.Headers = v
	return s
}

func (s *DebugModelResponse) SetStatusCode(v int32) *DebugModelResponse {
	s.StatusCode = &v
	return s
}

func (s *DebugModelResponse) SetBody(v *DebugModelResponseBody) *DebugModelResponse {
	s.Body = v
	return s
}

func (s *DebugModelResponse) Validate() error {
	return dara.Validate(s)
}
