// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMqGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMqGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMqGroupResponse
	GetStatusCode() *int32
	SetBody(v *MqGroup) *CreateMqGroupResponse
	GetBody() *MqGroup
}

type CreateMqGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MqGroup           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMqGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMqGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMqGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMqGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMqGroupResponse) GetBody() *MqGroup {
	return s.Body
}

func (s *CreateMqGroupResponse) SetHeaders(v map[string]*string) *CreateMqGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateMqGroupResponse) SetStatusCode(v int32) *CreateMqGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMqGroupResponse) SetBody(v *MqGroup) *CreateMqGroupResponse {
	s.Body = v
	return s
}

func (s *CreateMqGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
