// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStorageNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStorageNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetStorageNotifyConfigResponseBody) *GetStorageNotifyConfigResponse
	GetBody() *GetStorageNotifyConfigResponseBody
}

type GetStorageNotifyConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStorageNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStorageNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStorageNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *GetStorageNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStorageNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStorageNotifyConfigResponse) GetBody() *GetStorageNotifyConfigResponseBody {
	return s.Body
}

func (s *GetStorageNotifyConfigResponse) SetHeaders(v map[string]*string) *GetStorageNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *GetStorageNotifyConfigResponse) SetStatusCode(v int32) *GetStorageNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStorageNotifyConfigResponse) SetBody(v *GetStorageNotifyConfigResponseBody) *GetStorageNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *GetStorageNotifyConfigResponse) Validate() error {
	return dara.Validate(s)
}
