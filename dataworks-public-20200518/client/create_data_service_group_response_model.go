// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataServiceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataServiceGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataServiceGroupResponseBody) *CreateDataServiceGroupResponse
	GetBody() *CreateDataServiceGroupResponseBody
}

type CreateDataServiceGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataServiceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataServiceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDataServiceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataServiceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataServiceGroupResponse) GetBody() *CreateDataServiceGroupResponseBody {
	return s.Body
}

func (s *CreateDataServiceGroupResponse) SetHeaders(v map[string]*string) *CreateDataServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDataServiceGroupResponse) SetStatusCode(v int32) *CreateDataServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataServiceGroupResponse) SetBody(v *CreateDataServiceGroupResponseBody) *CreateDataServiceGroupResponse {
	s.Body = v
	return s
}

func (s *CreateDataServiceGroupResponse) Validate() error {
	return dara.Validate(s)
}
