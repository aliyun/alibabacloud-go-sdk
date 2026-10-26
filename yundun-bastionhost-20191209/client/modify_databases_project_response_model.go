// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabasesProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDatabasesProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDatabasesProjectResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDatabasesProjectResponseBody) *ModifyDatabasesProjectResponse
	GetBody() *ModifyDatabasesProjectResponseBody
}

type ModifyDatabasesProjectResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDatabasesProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDatabasesProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabasesProjectResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatabasesProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDatabasesProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDatabasesProjectResponse) GetBody() *ModifyDatabasesProjectResponseBody {
	return s.Body
}

func (s *ModifyDatabasesProjectResponse) SetHeaders(v map[string]*string) *ModifyDatabasesProjectResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatabasesProjectResponse) SetStatusCode(v int32) *ModifyDatabasesProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDatabasesProjectResponse) SetBody(v *ModifyDatabasesProjectResponseBody) *ModifyDatabasesProjectResponse {
	s.Body = v
	return s
}

func (s *ModifyDatabasesProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
