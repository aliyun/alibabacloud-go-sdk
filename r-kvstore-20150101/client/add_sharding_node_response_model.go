// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddShardingNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddShardingNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddShardingNodeResponse
	GetStatusCode() *int32
	SetBody(v *AddShardingNodeResponseBody) *AddShardingNodeResponse
	GetBody() *AddShardingNodeResponseBody
}

type AddShardingNodeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddShardingNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddShardingNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s AddShardingNodeResponse) GoString() string {
	return s.String()
}

func (s *AddShardingNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddShardingNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddShardingNodeResponse) GetBody() *AddShardingNodeResponseBody {
	return s.Body
}

func (s *AddShardingNodeResponse) SetHeaders(v map[string]*string) *AddShardingNodeResponse {
	s.Headers = v
	return s
}

func (s *AddShardingNodeResponse) SetStatusCode(v int32) *AddShardingNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *AddShardingNodeResponse) SetBody(v *AddShardingNodeResponseBody) *AddShardingNodeResponse {
	s.Body = v
	return s
}

func (s *AddShardingNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
