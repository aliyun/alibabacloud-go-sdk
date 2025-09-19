// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDiscoverDatabaseTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDiscoverDatabaseTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDiscoverDatabaseTaskResponse
	GetStatusCode() *int32
	SetBody(v *StartDiscoverDatabaseTaskResponseBody) *StartDiscoverDatabaseTaskResponse
	GetBody() *StartDiscoverDatabaseTaskResponseBody
}

type StartDiscoverDatabaseTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDiscoverDatabaseTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDiscoverDatabaseTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDiscoverDatabaseTaskResponse) GoString() string {
	return s.String()
}

func (s *StartDiscoverDatabaseTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDiscoverDatabaseTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDiscoverDatabaseTaskResponse) GetBody() *StartDiscoverDatabaseTaskResponseBody {
	return s.Body
}

func (s *StartDiscoverDatabaseTaskResponse) SetHeaders(v map[string]*string) *StartDiscoverDatabaseTaskResponse {
	s.Headers = v
	return s
}

func (s *StartDiscoverDatabaseTaskResponse) SetStatusCode(v int32) *StartDiscoverDatabaseTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDiscoverDatabaseTaskResponse) SetBody(v *StartDiscoverDatabaseTaskResponseBody) *StartDiscoverDatabaseTaskResponse {
	s.Body = v
	return s
}

func (s *StartDiscoverDatabaseTaskResponse) Validate() error {
	return dara.Validate(s)
}
