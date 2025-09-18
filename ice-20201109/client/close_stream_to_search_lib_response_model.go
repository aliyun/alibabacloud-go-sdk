// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseStreamToSearchLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseStreamToSearchLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseStreamToSearchLibResponse
	GetStatusCode() *int32
	SetBody(v *CloseStreamToSearchLibResponseBody) *CloseStreamToSearchLibResponse
	GetBody() *CloseStreamToSearchLibResponseBody
}

type CloseStreamToSearchLibResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseStreamToSearchLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseStreamToSearchLibResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseStreamToSearchLibResponse) GoString() string {
	return s.String()
}

func (s *CloseStreamToSearchLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseStreamToSearchLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseStreamToSearchLibResponse) GetBody() *CloseStreamToSearchLibResponseBody {
	return s.Body
}

func (s *CloseStreamToSearchLibResponse) SetHeaders(v map[string]*string) *CloseStreamToSearchLibResponse {
	s.Headers = v
	return s
}

func (s *CloseStreamToSearchLibResponse) SetStatusCode(v int32) *CloseStreamToSearchLibResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseStreamToSearchLibResponse) SetBody(v *CloseStreamToSearchLibResponseBody) *CloseStreamToSearchLibResponse {
	s.Body = v
	return s
}

func (s *CloseStreamToSearchLibResponse) Validate() error {
	return dara.Validate(s)
}
