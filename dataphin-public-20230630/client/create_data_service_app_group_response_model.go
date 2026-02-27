// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceAppGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataServiceAppGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataServiceAppGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataServiceAppGroupResponseBody) *CreateDataServiceAppGroupResponse
	GetBody() *CreateDataServiceAppGroupResponseBody
}

type CreateDataServiceAppGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataServiceAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataServiceAppGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceAppGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDataServiceAppGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataServiceAppGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataServiceAppGroupResponse) GetBody() *CreateDataServiceAppGroupResponseBody {
	return s.Body
}

func (s *CreateDataServiceAppGroupResponse) SetHeaders(v map[string]*string) *CreateDataServiceAppGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDataServiceAppGroupResponse) SetStatusCode(v int32) *CreateDataServiceAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataServiceAppGroupResponse) SetBody(v *CreateDataServiceAppGroupResponseBody) *CreateDataServiceAppGroupResponse {
	s.Body = v
	return s
}

func (s *CreateDataServiceAppGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
