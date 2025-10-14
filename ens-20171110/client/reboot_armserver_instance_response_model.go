// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootARMServerInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebootARMServerInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebootARMServerInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RebootARMServerInstanceResponseBody) *RebootARMServerInstanceResponse
	GetBody() *RebootARMServerInstanceResponseBody
}

type RebootARMServerInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootARMServerInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootARMServerInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootARMServerInstanceResponse) GoString() string {
	return s.String()
}

func (s *RebootARMServerInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebootARMServerInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebootARMServerInstanceResponse) GetBody() *RebootARMServerInstanceResponseBody {
	return s.Body
}

func (s *RebootARMServerInstanceResponse) SetHeaders(v map[string]*string) *RebootARMServerInstanceResponse {
	s.Headers = v
	return s
}

func (s *RebootARMServerInstanceResponse) SetStatusCode(v int32) *RebootARMServerInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootARMServerInstanceResponse) SetBody(v *RebootARMServerInstanceResponseBody) *RebootARMServerInstanceResponse {
	s.Body = v
	return s
}

func (s *RebootARMServerInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
