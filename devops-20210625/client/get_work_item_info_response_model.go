// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkItemInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkItemInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkItemInfoResponseBody) *GetWorkItemInfoResponse
	GetBody() *GetWorkItemInfoResponseBody
}

type GetWorkItemInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkItemInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkItemInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemInfoResponse) GoString() string {
	return s.String()
}

func (s *GetWorkItemInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkItemInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkItemInfoResponse) GetBody() *GetWorkItemInfoResponseBody {
	return s.Body
}

func (s *GetWorkItemInfoResponse) SetHeaders(v map[string]*string) *GetWorkItemInfoResponse {
	s.Headers = v
	return s
}

func (s *GetWorkItemInfoResponse) SetStatusCode(v int32) *GetWorkItemInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkItemInfoResponse) SetBody(v *GetWorkItemInfoResponseBody) *GetWorkItemInfoResponse {
	s.Body = v
	return s
}

func (s *GetWorkItemInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
