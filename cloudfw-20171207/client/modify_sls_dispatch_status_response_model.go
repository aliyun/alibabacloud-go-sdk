// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySlsDispatchStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySlsDispatchStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySlsDispatchStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifySlsDispatchStatusResponseBody) *ModifySlsDispatchStatusResponse
	GetBody() *ModifySlsDispatchStatusResponseBody
}

type ModifySlsDispatchStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySlsDispatchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySlsDispatchStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySlsDispatchStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifySlsDispatchStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySlsDispatchStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySlsDispatchStatusResponse) GetBody() *ModifySlsDispatchStatusResponseBody {
	return s.Body
}

func (s *ModifySlsDispatchStatusResponse) SetHeaders(v map[string]*string) *ModifySlsDispatchStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifySlsDispatchStatusResponse) SetStatusCode(v int32) *ModifySlsDispatchStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySlsDispatchStatusResponse) SetBody(v *ModifySlsDispatchStatusResponseBody) *ModifySlsDispatchStatusResponse {
	s.Body = v
	return s
}

func (s *ModifySlsDispatchStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
