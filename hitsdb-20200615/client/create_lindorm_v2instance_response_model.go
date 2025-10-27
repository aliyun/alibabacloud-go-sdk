// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLindormV2InstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLindormV2InstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLindormV2InstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateLindormV2InstanceResponseBody) *CreateLindormV2InstanceResponse
	GetBody() *CreateLindormV2InstanceResponseBody
}

type CreateLindormV2InstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLindormV2InstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLindormV2InstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLindormV2InstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateLindormV2InstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLindormV2InstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLindormV2InstanceResponse) GetBody() *CreateLindormV2InstanceResponseBody {
	return s.Body
}

func (s *CreateLindormV2InstanceResponse) SetHeaders(v map[string]*string) *CreateLindormV2InstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateLindormV2InstanceResponse) SetStatusCode(v int32) *CreateLindormV2InstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLindormV2InstanceResponse) SetBody(v *CreateLindormV2InstanceResponseBody) *CreateLindormV2InstanceResponse {
	s.Body = v
	return s
}

func (s *CreateLindormV2InstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
