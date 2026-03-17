// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQosCarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyQosCarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyQosCarResponse
	GetStatusCode() *int32
	SetBody(v *ModifyQosCarResponseBody) *ModifyQosCarResponse
	GetBody() *ModifyQosCarResponseBody
}

type ModifyQosCarResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyQosCarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyQosCarResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyQosCarResponse) GoString() string {
	return s.String()
}

func (s *ModifyQosCarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyQosCarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyQosCarResponse) GetBody() *ModifyQosCarResponseBody {
	return s.Body
}

func (s *ModifyQosCarResponse) SetHeaders(v map[string]*string) *ModifyQosCarResponse {
	s.Headers = v
	return s
}

func (s *ModifyQosCarResponse) SetStatusCode(v int32) *ModifyQosCarResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyQosCarResponse) SetBody(v *ModifyQosCarResponseBody) *ModifyQosCarResponse {
	s.Body = v
	return s
}

func (s *ModifyQosCarResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
