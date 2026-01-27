// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCDiskAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRCDiskAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRCDiskAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRCDiskAttributeResponseBody) *ModifyRCDiskAttributeResponse
	GetBody() *ModifyRCDiskAttributeResponseBody
}

type ModifyRCDiskAttributeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRCDiskAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRCDiskAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCDiskAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyRCDiskAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRCDiskAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRCDiskAttributeResponse) GetBody() *ModifyRCDiskAttributeResponseBody {
	return s.Body
}

func (s *ModifyRCDiskAttributeResponse) SetHeaders(v map[string]*string) *ModifyRCDiskAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyRCDiskAttributeResponse) SetStatusCode(v int32) *ModifyRCDiskAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRCDiskAttributeResponse) SetBody(v *ModifyRCDiskAttributeResponseBody) *ModifyRCDiskAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyRCDiskAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
