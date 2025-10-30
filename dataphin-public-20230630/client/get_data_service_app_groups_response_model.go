// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAppGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceAppGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceAppGroupsResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceAppGroupsResponseBody) *GetDataServiceAppGroupsResponse
	GetBody() *GetDataServiceAppGroupsResponseBody
}

type GetDataServiceAppGroupsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceAppGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceAppGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppGroupsResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceAppGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceAppGroupsResponse) GetBody() *GetDataServiceAppGroupsResponseBody {
	return s.Body
}

func (s *GetDataServiceAppGroupsResponse) SetHeaders(v map[string]*string) *GetDataServiceAppGroupsResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceAppGroupsResponse) SetStatusCode(v int32) *GetDataServiceAppGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceAppGroupsResponse) SetBody(v *GetDataServiceAppGroupsResponseBody) *GetDataServiceAppGroupsResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceAppGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
