// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkItemResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkItemResponseBody) *GetWorkItemResponse
	GetBody() *GetWorkItemResponseBody
}

type GetWorkItemResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkItemResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemResponse) GoString() string {
	return s.String()
}

func (s *GetWorkItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkItemResponse) GetBody() *GetWorkItemResponseBody {
	return s.Body
}

func (s *GetWorkItemResponse) SetHeaders(v map[string]*string) *GetWorkItemResponse {
	s.Headers = v
	return s
}

func (s *GetWorkItemResponse) SetStatusCode(v int32) *GetWorkItemResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkItemResponse) SetBody(v *GetWorkItemResponseBody) *GetWorkItemResponse {
	s.Body = v
	return s
}

func (s *GetWorkItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
