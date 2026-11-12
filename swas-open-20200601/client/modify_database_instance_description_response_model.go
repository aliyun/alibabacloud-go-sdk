// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabaseInstanceDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDatabaseInstanceDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDatabaseInstanceDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDatabaseInstanceDescriptionResponseBody) *ModifyDatabaseInstanceDescriptionResponse
	GetBody() *ModifyDatabaseInstanceDescriptionResponseBody
}

type ModifyDatabaseInstanceDescriptionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDatabaseInstanceDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDatabaseInstanceDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabaseInstanceDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseInstanceDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDatabaseInstanceDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDatabaseInstanceDescriptionResponse) GetBody() *ModifyDatabaseInstanceDescriptionResponseBody {
	return s.Body
}

func (s *ModifyDatabaseInstanceDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDatabaseInstanceDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatabaseInstanceDescriptionResponse) SetStatusCode(v int32) *ModifyDatabaseInstanceDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDatabaseInstanceDescriptionResponse) SetBody(v *ModifyDatabaseInstanceDescriptionResponseBody) *ModifyDatabaseInstanceDescriptionResponse {
	s.Body = v
	return s
}

func (s *ModifyDatabaseInstanceDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
