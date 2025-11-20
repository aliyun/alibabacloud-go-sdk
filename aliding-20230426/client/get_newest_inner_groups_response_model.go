// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNewestInnerGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNewestInnerGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNewestInnerGroupsResponse
	GetStatusCode() *int32
	SetBody(v *GetNewestInnerGroupsResponseBody) *GetNewestInnerGroupsResponse
	GetBody() *GetNewestInnerGroupsResponseBody
}

type GetNewestInnerGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNewestInnerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNewestInnerGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNewestInnerGroupsResponse) GoString() string {
	return s.String()
}

func (s *GetNewestInnerGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNewestInnerGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNewestInnerGroupsResponse) GetBody() *GetNewestInnerGroupsResponseBody {
	return s.Body
}

func (s *GetNewestInnerGroupsResponse) SetHeaders(v map[string]*string) *GetNewestInnerGroupsResponse {
	s.Headers = v
	return s
}

func (s *GetNewestInnerGroupsResponse) SetStatusCode(v int32) *GetNewestInnerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNewestInnerGroupsResponse) SetBody(v *GetNewestInnerGroupsResponseBody) *GetNewestInnerGroupsResponse {
	s.Body = v
	return s
}

func (s *GetNewestInnerGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
