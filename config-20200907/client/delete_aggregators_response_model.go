// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggregatorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAggregatorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAggregatorsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAggregatorsResponseBody) *DeleteAggregatorsResponse
	GetBody() *DeleteAggregatorsResponseBody
}

type DeleteAggregatorsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAggregatorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAggregatorsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregatorsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAggregatorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAggregatorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAggregatorsResponse) GetBody() *DeleteAggregatorsResponseBody {
	return s.Body
}

func (s *DeleteAggregatorsResponse) SetHeaders(v map[string]*string) *DeleteAggregatorsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAggregatorsResponse) SetStatusCode(v int32) *DeleteAggregatorsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAggregatorsResponse) SetBody(v *DeleteAggregatorsResponseBody) *DeleteAggregatorsResponse {
	s.Body = v
	return s
}

func (s *DeleteAggregatorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
