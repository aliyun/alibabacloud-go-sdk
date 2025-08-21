// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreUserFromRecycleBinResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestoreUserFromRecycleBinResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestoreUserFromRecycleBinResponse
	GetStatusCode() *int32
	SetBody(v *RestoreUserFromRecycleBinResponseBody) *RestoreUserFromRecycleBinResponse
	GetBody() *RestoreUserFromRecycleBinResponseBody
}

type RestoreUserFromRecycleBinResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreUserFromRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreUserFromRecycleBinResponse) String() string {
	return dara.Prettify(s)
}

func (s RestoreUserFromRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *RestoreUserFromRecycleBinResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestoreUserFromRecycleBinResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestoreUserFromRecycleBinResponse) GetBody() *RestoreUserFromRecycleBinResponseBody {
	return s.Body
}

func (s *RestoreUserFromRecycleBinResponse) SetHeaders(v map[string]*string) *RestoreUserFromRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *RestoreUserFromRecycleBinResponse) SetStatusCode(v int32) *RestoreUserFromRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreUserFromRecycleBinResponse) SetBody(v *RestoreUserFromRecycleBinResponseBody) *RestoreUserFromRecycleBinResponse {
	s.Body = v
	return s
}

func (s *RestoreUserFromRecycleBinResponse) Validate() error {
	return dara.Validate(s)
}
