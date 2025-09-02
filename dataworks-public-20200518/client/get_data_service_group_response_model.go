// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceGroupResponseBody) *GetDataServiceGroupResponse
	GetBody() *GetDataServiceGroupResponseBody
}

type GetDataServiceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceGroupResponse) GetBody() *GetDataServiceGroupResponseBody {
	return s.Body
}

func (s *GetDataServiceGroupResponse) SetHeaders(v map[string]*string) *GetDataServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceGroupResponse) SetStatusCode(v int32) *GetDataServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceGroupResponse) SetBody(v *GetDataServiceGroupResponseBody) *GetDataServiceGroupResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceGroupResponse) Validate() error {
	return dara.Validate(s)
}
