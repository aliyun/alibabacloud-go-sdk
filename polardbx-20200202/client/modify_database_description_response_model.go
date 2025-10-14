// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabaseDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDatabaseDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDatabaseDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDatabaseDescriptionResponseBody) *ModifyDatabaseDescriptionResponse
	GetBody() *ModifyDatabaseDescriptionResponseBody
}

type ModifyDatabaseDescriptionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDatabaseDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDatabaseDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabaseDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDatabaseDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDatabaseDescriptionResponse) GetBody() *ModifyDatabaseDescriptionResponseBody {
	return s.Body
}

func (s *ModifyDatabaseDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDatabaseDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatabaseDescriptionResponse) SetStatusCode(v int32) *ModifyDatabaseDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDatabaseDescriptionResponse) SetBody(v *ModifyDatabaseDescriptionResponseBody) *ModifyDatabaseDescriptionResponse {
	s.Body = v
	return s
}

func (s *ModifyDatabaseDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
