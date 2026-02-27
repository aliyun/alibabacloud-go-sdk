// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataServiceAppGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataServiceAppGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataServiceAppGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataServiceAppGroupResponseBody) *UpdateDataServiceAppGroupResponse
	GetBody() *UpdateDataServiceAppGroupResponseBody
}

type UpdateDataServiceAppGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataServiceAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataServiceAppGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataServiceAppGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataServiceAppGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataServiceAppGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataServiceAppGroupResponse) GetBody() *UpdateDataServiceAppGroupResponseBody {
	return s.Body
}

func (s *UpdateDataServiceAppGroupResponse) SetHeaders(v map[string]*string) *UpdateDataServiceAppGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataServiceAppGroupResponse) SetStatusCode(v int32) *UpdateDataServiceAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataServiceAppGroupResponse) SetBody(v *UpdateDataServiceAppGroupResponseBody) *UpdateDataServiceAppGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateDataServiceAppGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
