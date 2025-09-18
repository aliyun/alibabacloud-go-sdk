// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddStreamTagToSearchLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddStreamTagToSearchLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddStreamTagToSearchLibResponse
	GetStatusCode() *int32
	SetBody(v *AddStreamTagToSearchLibResponseBody) *AddStreamTagToSearchLibResponse
	GetBody() *AddStreamTagToSearchLibResponseBody
}

type AddStreamTagToSearchLibResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddStreamTagToSearchLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddStreamTagToSearchLibResponse) String() string {
	return dara.Prettify(s)
}

func (s AddStreamTagToSearchLibResponse) GoString() string {
	return s.String()
}

func (s *AddStreamTagToSearchLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddStreamTagToSearchLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddStreamTagToSearchLibResponse) GetBody() *AddStreamTagToSearchLibResponseBody {
	return s.Body
}

func (s *AddStreamTagToSearchLibResponse) SetHeaders(v map[string]*string) *AddStreamTagToSearchLibResponse {
	s.Headers = v
	return s
}

func (s *AddStreamTagToSearchLibResponse) SetStatusCode(v int32) *AddStreamTagToSearchLibResponse {
	s.StatusCode = &v
	return s
}

func (s *AddStreamTagToSearchLibResponse) SetBody(v *AddStreamTagToSearchLibResponseBody) *AddStreamTagToSearchLibResponse {
	s.Body = v
	return s
}

func (s *AddStreamTagToSearchLibResponse) Validate() error {
	return dara.Validate(s)
}
