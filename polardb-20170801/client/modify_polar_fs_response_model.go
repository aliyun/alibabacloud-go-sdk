// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolarFsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPolarFsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPolarFsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPolarFsResponseBody) *ModifyPolarFsResponse
	GetBody() *ModifyPolarFsResponseBody
}

type ModifyPolarFsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPolarFsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPolarFsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolarFsResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolarFsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPolarFsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPolarFsResponse) GetBody() *ModifyPolarFsResponseBody {
	return s.Body
}

func (s *ModifyPolarFsResponse) SetHeaders(v map[string]*string) *ModifyPolarFsResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolarFsResponse) SetStatusCode(v int32) *ModifyPolarFsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPolarFsResponse) SetBody(v *ModifyPolarFsResponseBody) *ModifyPolarFsResponse {
	s.Body = v
	return s
}

func (s *ModifyPolarFsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
