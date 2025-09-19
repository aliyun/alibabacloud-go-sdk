// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSoarStrategyTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSoarStrategyTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSoarStrategyTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSoarStrategyTaskResponseBody) *DeleteSoarStrategyTaskResponse
	GetBody() *DeleteSoarStrategyTaskResponseBody
}

type DeleteSoarStrategyTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSoarStrategyTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSoarStrategyTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSoarStrategyTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteSoarStrategyTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSoarStrategyTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSoarStrategyTaskResponse) GetBody() *DeleteSoarStrategyTaskResponseBody {
	return s.Body
}

func (s *DeleteSoarStrategyTaskResponse) SetHeaders(v map[string]*string) *DeleteSoarStrategyTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteSoarStrategyTaskResponse) SetStatusCode(v int32) *DeleteSoarStrategyTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSoarStrategyTaskResponse) SetBody(v *DeleteSoarStrategyTaskResponseBody) *DeleteSoarStrategyTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteSoarStrategyTaskResponse) Validate() error {
	return dara.Validate(s)
}
