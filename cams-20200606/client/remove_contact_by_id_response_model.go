// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveContactByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveContactByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveContactByIdResponse
	GetStatusCode() *int32
	SetBody(v *RemoveContactByIdResponseBody) *RemoveContactByIdResponse
	GetBody() *RemoveContactByIdResponseBody
}

type RemoveContactByIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveContactByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveContactByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveContactByIdResponse) GoString() string {
	return s.String()
}

func (s *RemoveContactByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveContactByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveContactByIdResponse) GetBody() *RemoveContactByIdResponseBody {
	return s.Body
}

func (s *RemoveContactByIdResponse) SetHeaders(v map[string]*string) *RemoveContactByIdResponse {
	s.Headers = v
	return s
}

func (s *RemoveContactByIdResponse) SetStatusCode(v int32) *RemoveContactByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveContactByIdResponse) SetBody(v *RemoveContactByIdResponseBody) *RemoveContactByIdResponse {
	s.Body = v
	return s
}

func (s *RemoveContactByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
