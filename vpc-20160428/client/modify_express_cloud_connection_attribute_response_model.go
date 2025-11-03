// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressCloudConnectionAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyExpressCloudConnectionAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyExpressCloudConnectionAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyExpressCloudConnectionAttributeResponseBody) *ModifyExpressCloudConnectionAttributeResponse
	GetBody() *ModifyExpressCloudConnectionAttributeResponseBody
}

type ModifyExpressCloudConnectionAttributeResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyExpressCloudConnectionAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyExpressCloudConnectionAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressCloudConnectionAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyExpressCloudConnectionAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyExpressCloudConnectionAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyExpressCloudConnectionAttributeResponse) GetBody() *ModifyExpressCloudConnectionAttributeResponseBody {
	return s.Body
}

func (s *ModifyExpressCloudConnectionAttributeResponse) SetHeaders(v map[string]*string) *ModifyExpressCloudConnectionAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyExpressCloudConnectionAttributeResponse) SetStatusCode(v int32) *ModifyExpressCloudConnectionAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyExpressCloudConnectionAttributeResponse) SetBody(v *ModifyExpressCloudConnectionAttributeResponseBody) *ModifyExpressCloudConnectionAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyExpressCloudConnectionAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
