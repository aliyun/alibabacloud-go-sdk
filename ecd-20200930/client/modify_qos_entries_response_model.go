// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQosEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyQosEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyQosEntriesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyQosEntriesResponseBody) *ModifyQosEntriesResponse
	GetBody() *ModifyQosEntriesResponseBody
}

type ModifyQosEntriesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyQosEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyQosEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyQosEntriesResponse) GoString() string {
	return s.String()
}

func (s *ModifyQosEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyQosEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyQosEntriesResponse) GetBody() *ModifyQosEntriesResponseBody {
	return s.Body
}

func (s *ModifyQosEntriesResponse) SetHeaders(v map[string]*string) *ModifyQosEntriesResponse {
	s.Headers = v
	return s
}

func (s *ModifyQosEntriesResponse) SetStatusCode(v int32) *ModifyQosEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyQosEntriesResponse) SetBody(v *ModifyQosEntriesResponseBody) *ModifyQosEntriesResponse {
	s.Body = v
	return s
}

func (s *ModifyQosEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
