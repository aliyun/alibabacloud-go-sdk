// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudPhoneNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCloudPhoneNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCloudPhoneNodeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCloudPhoneNodeResponseBody) *ModifyCloudPhoneNodeResponse
	GetBody() *ModifyCloudPhoneNodeResponseBody
}

type ModifyCloudPhoneNodeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCloudPhoneNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCloudPhoneNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudPhoneNodeResponse) GoString() string {
	return s.String()
}

func (s *ModifyCloudPhoneNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCloudPhoneNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCloudPhoneNodeResponse) GetBody() *ModifyCloudPhoneNodeResponseBody {
	return s.Body
}

func (s *ModifyCloudPhoneNodeResponse) SetHeaders(v map[string]*string) *ModifyCloudPhoneNodeResponse {
	s.Headers = v
	return s
}

func (s *ModifyCloudPhoneNodeResponse) SetStatusCode(v int32) *ModifyCloudPhoneNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCloudPhoneNodeResponse) SetBody(v *ModifyCloudPhoneNodeResponseBody) *ModifyCloudPhoneNodeResponse {
	s.Body = v
	return s
}

func (s *ModifyCloudPhoneNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
