// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreAccessKeyFromRecycleBinResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestoreAccessKeyFromRecycleBinResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestoreAccessKeyFromRecycleBinResponse
	GetStatusCode() *int32
	SetBody(v *RestoreAccessKeyFromRecycleBinResponseBody) *RestoreAccessKeyFromRecycleBinResponse
	GetBody() *RestoreAccessKeyFromRecycleBinResponseBody
}

type RestoreAccessKeyFromRecycleBinResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreAccessKeyFromRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreAccessKeyFromRecycleBinResponse) String() string {
	return dara.Prettify(s)
}

func (s RestoreAccessKeyFromRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *RestoreAccessKeyFromRecycleBinResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestoreAccessKeyFromRecycleBinResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestoreAccessKeyFromRecycleBinResponse) GetBody() *RestoreAccessKeyFromRecycleBinResponseBody {
	return s.Body
}

func (s *RestoreAccessKeyFromRecycleBinResponse) SetHeaders(v map[string]*string) *RestoreAccessKeyFromRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *RestoreAccessKeyFromRecycleBinResponse) SetStatusCode(v int32) *RestoreAccessKeyFromRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreAccessKeyFromRecycleBinResponse) SetBody(v *RestoreAccessKeyFromRecycleBinResponseBody) *RestoreAccessKeyFromRecycleBinResponse {
	s.Body = v
	return s
}

func (s *RestoreAccessKeyFromRecycleBinResponse) Validate() error {
	return dara.Validate(s)
}
