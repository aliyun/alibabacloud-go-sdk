// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLogMetaStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLogMetaStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLogMetaStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLogMetaStatusResponseBody) *ModifyLogMetaStatusResponse
	GetBody() *ModifyLogMetaStatusResponseBody
}

type ModifyLogMetaStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLogMetaStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLogMetaStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLogMetaStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyLogMetaStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLogMetaStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLogMetaStatusResponse) GetBody() *ModifyLogMetaStatusResponseBody {
	return s.Body
}

func (s *ModifyLogMetaStatusResponse) SetHeaders(v map[string]*string) *ModifyLogMetaStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyLogMetaStatusResponse) SetStatusCode(v int32) *ModifyLogMetaStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLogMetaStatusResponse) SetBody(v *ModifyLogMetaStatusResponseBody) *ModifyLogMetaStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyLogMetaStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
