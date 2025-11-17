// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstancePropertiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstancePropertiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstancePropertiesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstancePropertiesResponseBody) *ModifyInstancePropertiesResponse
	GetBody() *ModifyInstancePropertiesResponseBody
}

type ModifyInstancePropertiesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstancePropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstancePropertiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstancePropertiesResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstancePropertiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstancePropertiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstancePropertiesResponse) GetBody() *ModifyInstancePropertiesResponseBody {
	return s.Body
}

func (s *ModifyInstancePropertiesResponse) SetHeaders(v map[string]*string) *ModifyInstancePropertiesResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstancePropertiesResponse) SetStatusCode(v int32) *ModifyInstancePropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstancePropertiesResponse) SetBody(v *ModifyInstancePropertiesResponseBody) *ModifyInstancePropertiesResponse {
	s.Body = v
	return s
}

func (s *ModifyInstancePropertiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
