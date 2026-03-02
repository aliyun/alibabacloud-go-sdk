// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMqGroupMsgsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMqGroupMsgsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMqGroupMsgsResponse
	GetStatusCode() *int32
	SetBody(v *MqMsgPageResult) *ListMqGroupMsgsResponse
	GetBody() *MqMsgPageResult
}

type ListMqGroupMsgsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MqMsgPageResult   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMqGroupMsgsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMqGroupMsgsResponse) GoString() string {
	return s.String()
}

func (s *ListMqGroupMsgsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMqGroupMsgsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMqGroupMsgsResponse) GetBody() *MqMsgPageResult {
	return s.Body
}

func (s *ListMqGroupMsgsResponse) SetHeaders(v map[string]*string) *ListMqGroupMsgsResponse {
	s.Headers = v
	return s
}

func (s *ListMqGroupMsgsResponse) SetStatusCode(v int32) *ListMqGroupMsgsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMqGroupMsgsResponse) SetBody(v *MqMsgPageResult) *ListMqGroupMsgsResponse {
	s.Body = v
	return s
}

func (s *ListMqGroupMsgsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
