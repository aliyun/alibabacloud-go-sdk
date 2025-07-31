// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAppsByGroupIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceAppsByGroupIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceAppsByGroupIdResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceAppsByGroupIdResponseBody) *GetDataServiceAppsByGroupIdResponse
	GetBody() *GetDataServiceAppsByGroupIdResponseBody
}

type GetDataServiceAppsByGroupIdResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceAppsByGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceAppsByGroupIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppsByGroupIdResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppsByGroupIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceAppsByGroupIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceAppsByGroupIdResponse) GetBody() *GetDataServiceAppsByGroupIdResponseBody {
	return s.Body
}

func (s *GetDataServiceAppsByGroupIdResponse) SetHeaders(v map[string]*string) *GetDataServiceAppsByGroupIdResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceAppsByGroupIdResponse) SetStatusCode(v int32) *GetDataServiceAppsByGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceAppsByGroupIdResponse) SetBody(v *GetDataServiceAppsByGroupIdResponseBody) *GetDataServiceAppsByGroupIdResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceAppsByGroupIdResponse) Validate() error {
	return dara.Validate(s)
}
