// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabaseInstanceParameterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDatabaseInstanceParameterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDatabaseInstanceParameterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDatabaseInstanceParameterResponseBody) *ModifyDatabaseInstanceParameterResponse
	GetBody() *ModifyDatabaseInstanceParameterResponseBody
}

type ModifyDatabaseInstanceParameterResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDatabaseInstanceParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDatabaseInstanceParameterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabaseInstanceParameterResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseInstanceParameterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDatabaseInstanceParameterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDatabaseInstanceParameterResponse) GetBody() *ModifyDatabaseInstanceParameterResponseBody {
	return s.Body
}

func (s *ModifyDatabaseInstanceParameterResponse) SetHeaders(v map[string]*string) *ModifyDatabaseInstanceParameterResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatabaseInstanceParameterResponse) SetStatusCode(v int32) *ModifyDatabaseInstanceParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDatabaseInstanceParameterResponse) SetBody(v *ModifyDatabaseInstanceParameterResponseBody) *ModifyDatabaseInstanceParameterResponse {
	s.Body = v
	return s
}

func (s *ModifyDatabaseInstanceParameterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
