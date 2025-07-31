// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAuthorizedAppsByGroupIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceAuthorizedAppsByGroupIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceAuthorizedAppsByGroupIdResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceAuthorizedAppsByGroupIdResponseBody) *GetDataServiceAuthorizedAppsByGroupIdResponse
	GetBody() *GetDataServiceAuthorizedAppsByGroupIdResponseBody
}

type GetDataServiceAuthorizedAppsByGroupIdResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceAuthorizedAppsByGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceAuthorizedAppsByGroupIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAuthorizedAppsByGroupIdResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponse) GetBody() *GetDataServiceAuthorizedAppsByGroupIdResponseBody {
	return s.Body
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponse) SetHeaders(v map[string]*string) *GetDataServiceAuthorizedAppsByGroupIdResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponse) SetStatusCode(v int32) *GetDataServiceAuthorizedAppsByGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponse) SetBody(v *GetDataServiceAuthorizedAppsByGroupIdResponseBody) *GetDataServiceAuthorizedAppsByGroupIdResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceAuthorizedAppsByGroupIdResponse) Validate() error {
	return dara.Validate(s)
}
