// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabaseConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDatabaseConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDatabaseConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDatabaseConfigResponseBody) *ModifyDatabaseConfigResponse
	GetBody() *ModifyDatabaseConfigResponseBody
}

type ModifyDatabaseConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDatabaseConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDatabaseConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabaseConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDatabaseConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDatabaseConfigResponse) GetBody() *ModifyDatabaseConfigResponseBody {
	return s.Body
}

func (s *ModifyDatabaseConfigResponse) SetHeaders(v map[string]*string) *ModifyDatabaseConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatabaseConfigResponse) SetStatusCode(v int32) *ModifyDatabaseConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDatabaseConfigResponse) SetBody(v *ModifyDatabaseConfigResponseBody) *ModifyDatabaseConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyDatabaseConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
