// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSyntheticTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSyntheticTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSyntheticTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSyntheticTaskResponseBody) *DeleteSyntheticTaskResponse
	GetBody() *DeleteSyntheticTaskResponseBody
}

type DeleteSyntheticTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSyntheticTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSyntheticTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSyntheticTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteSyntheticTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSyntheticTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSyntheticTaskResponse) GetBody() *DeleteSyntheticTaskResponseBody {
	return s.Body
}

func (s *DeleteSyntheticTaskResponse) SetHeaders(v map[string]*string) *DeleteSyntheticTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteSyntheticTaskResponse) SetStatusCode(v int32) *DeleteSyntheticTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSyntheticTaskResponse) SetBody(v *DeleteSyntheticTaskResponseBody) *DeleteSyntheticTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteSyntheticTaskResponse) Validate() error {
	return dara.Validate(s)
}
