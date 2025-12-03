// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkitemTimeTypeListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkitemTimeTypeListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkitemTimeTypeListResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkitemTimeTypeListResponseBody) *GetWorkitemTimeTypeListResponse
	GetBody() *GetWorkitemTimeTypeListResponseBody
}

type GetWorkitemTimeTypeListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkitemTimeTypeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkitemTimeTypeListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkitemTimeTypeListResponse) GoString() string {
	return s.String()
}

func (s *GetWorkitemTimeTypeListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkitemTimeTypeListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkitemTimeTypeListResponse) GetBody() *GetWorkitemTimeTypeListResponseBody {
	return s.Body
}

func (s *GetWorkitemTimeTypeListResponse) SetHeaders(v map[string]*string) *GetWorkitemTimeTypeListResponse {
	s.Headers = v
	return s
}

func (s *GetWorkitemTimeTypeListResponse) SetStatusCode(v int32) *GetWorkitemTimeTypeListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponse) SetBody(v *GetWorkitemTimeTypeListResponseBody) *GetWorkitemTimeTypeListResponse {
	s.Body = v
	return s
}

func (s *GetWorkitemTimeTypeListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
