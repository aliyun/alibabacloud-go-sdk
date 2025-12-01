// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataLimitResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataLimitResponseBody) *DeleteDataLimitResponse
	GetBody() *DeleteDataLimitResponseBody
}

type DeleteDataLimitResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLimitResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataLimitResponse) GetBody() *DeleteDataLimitResponseBody {
	return s.Body
}

func (s *DeleteDataLimitResponse) SetHeaders(v map[string]*string) *DeleteDataLimitResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataLimitResponse) SetStatusCode(v int32) *DeleteDataLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataLimitResponse) SetBody(v *DeleteDataLimitResponseBody) *DeleteDataLimitResponse {
	s.Body = v
	return s
}

func (s *DeleteDataLimitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
