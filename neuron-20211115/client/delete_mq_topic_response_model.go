// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMqTopicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMqTopicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMqTopicResponse
	GetStatusCode() *int32
}

type DeleteMqTopicResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteMqTopicResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMqTopicResponse) GoString() string {
	return s.String()
}

func (s *DeleteMqTopicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMqTopicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMqTopicResponse) SetHeaders(v map[string]*string) *DeleteMqTopicResponse {
	s.Headers = v
	return s
}

func (s *DeleteMqTopicResponse) SetStatusCode(v int32) *DeleteMqTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMqTopicResponse) Validate() error {
	return dara.Validate(s)
}
