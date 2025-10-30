// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExternalDataServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyExternalDataServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyExternalDataServiceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyExternalDataServiceResponseBody) *ModifyExternalDataServiceResponse
	GetBody() *ModifyExternalDataServiceResponseBody
}

type ModifyExternalDataServiceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyExternalDataServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyExternalDataServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyExternalDataServiceResponse) GoString() string {
	return s.String()
}

func (s *ModifyExternalDataServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyExternalDataServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyExternalDataServiceResponse) GetBody() *ModifyExternalDataServiceResponseBody {
	return s.Body
}

func (s *ModifyExternalDataServiceResponse) SetHeaders(v map[string]*string) *ModifyExternalDataServiceResponse {
	s.Headers = v
	return s
}

func (s *ModifyExternalDataServiceResponse) SetStatusCode(v int32) *ModifyExternalDataServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyExternalDataServiceResponse) SetBody(v *ModifyExternalDataServiceResponseBody) *ModifyExternalDataServiceResponse {
	s.Body = v
	return s
}

func (s *ModifyExternalDataServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
