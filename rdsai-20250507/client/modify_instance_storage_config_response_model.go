// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceStorageConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceStorageConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceStorageConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceStorageConfigResponseBody) *ModifyInstanceStorageConfigResponse
	GetBody() *ModifyInstanceStorageConfigResponseBody
}

type ModifyInstanceStorageConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceStorageConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceStorageConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceStorageConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceStorageConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceStorageConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceStorageConfigResponse) GetBody() *ModifyInstanceStorageConfigResponseBody {
	return s.Body
}

func (s *ModifyInstanceStorageConfigResponse) SetHeaders(v map[string]*string) *ModifyInstanceStorageConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceStorageConfigResponse) SetStatusCode(v int32) *ModifyInstanceStorageConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceStorageConfigResponse) SetBody(v *ModifyInstanceStorageConfigResponseBody) *ModifyInstanceStorageConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceStorageConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
