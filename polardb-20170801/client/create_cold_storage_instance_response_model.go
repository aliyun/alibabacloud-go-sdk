// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateColdStorageInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateColdStorageInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateColdStorageInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateColdStorageInstanceResponseBody) *CreateColdStorageInstanceResponse
	GetBody() *CreateColdStorageInstanceResponseBody
}

type CreateColdStorageInstanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateColdStorageInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateColdStorageInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateColdStorageInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateColdStorageInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateColdStorageInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateColdStorageInstanceResponse) GetBody() *CreateColdStorageInstanceResponseBody {
	return s.Body
}

func (s *CreateColdStorageInstanceResponse) SetHeaders(v map[string]*string) *CreateColdStorageInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateColdStorageInstanceResponse) SetStatusCode(v int32) *CreateColdStorageInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateColdStorageInstanceResponse) SetBody(v *CreateColdStorageInstanceResponseBody) *CreateColdStorageInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateColdStorageInstanceResponse) Validate() error {
	return dara.Validate(s)
}
