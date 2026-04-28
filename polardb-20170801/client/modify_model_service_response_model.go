// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyModelServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyModelServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyModelServiceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyModelServiceResponseBody) *ModifyModelServiceResponse
	GetBody() *ModifyModelServiceResponseBody
}

type ModifyModelServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyModelServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyModelServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelServiceResponse) GoString() string {
	return s.String()
}

func (s *ModifyModelServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyModelServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyModelServiceResponse) GetBody() *ModifyModelServiceResponseBody {
	return s.Body
}

func (s *ModifyModelServiceResponse) SetHeaders(v map[string]*string) *ModifyModelServiceResponse {
	s.Headers = v
	return s
}

func (s *ModifyModelServiceResponse) SetStatusCode(v int32) *ModifyModelServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyModelServiceResponse) SetBody(v *ModifyModelServiceResponseBody) *ModifyModelServiceResponse {
	s.Body = v
	return s
}

func (s *ModifyModelServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
