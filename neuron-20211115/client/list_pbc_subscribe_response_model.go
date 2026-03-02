// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPbcSubscribeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPbcSubscribeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPbcSubscribeResponse
	GetStatusCode() *int32
	SetBody(v *ListPbcSubscribeResponseBody) *ListPbcSubscribeResponse
	GetBody() *ListPbcSubscribeResponseBody
}

type ListPbcSubscribeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPbcSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPbcSubscribeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPbcSubscribeResponse) GoString() string {
	return s.String()
}

func (s *ListPbcSubscribeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPbcSubscribeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPbcSubscribeResponse) GetBody() *ListPbcSubscribeResponseBody {
	return s.Body
}

func (s *ListPbcSubscribeResponse) SetHeaders(v map[string]*string) *ListPbcSubscribeResponse {
	s.Headers = v
	return s
}

func (s *ListPbcSubscribeResponse) SetStatusCode(v int32) *ListPbcSubscribeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPbcSubscribeResponse) SetBody(v *ListPbcSubscribeResponseBody) *ListPbcSubscribeResponse {
	s.Body = v
	return s
}

func (s *ListPbcSubscribeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
