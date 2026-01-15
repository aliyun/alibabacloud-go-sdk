// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModuleBasicInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModuleBasicInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModuleBasicInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateModuleBasicInfoResponseBody) *UpdateModuleBasicInfoResponse
	GetBody() *UpdateModuleBasicInfoResponseBody
}

type UpdateModuleBasicInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModuleBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModuleBasicInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModuleBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateModuleBasicInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModuleBasicInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModuleBasicInfoResponse) GetBody() *UpdateModuleBasicInfoResponseBody {
	return s.Body
}

func (s *UpdateModuleBasicInfoResponse) SetHeaders(v map[string]*string) *UpdateModuleBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateModuleBasicInfoResponse) SetStatusCode(v int32) *UpdateModuleBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModuleBasicInfoResponse) SetBody(v *UpdateModuleBasicInfoResponseBody) *UpdateModuleBasicInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateModuleBasicInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
