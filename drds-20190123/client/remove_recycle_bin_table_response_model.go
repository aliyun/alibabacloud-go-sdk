// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRecycleBinTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveRecycleBinTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveRecycleBinTableResponse
	GetStatusCode() *int32
	SetBody(v *RemoveRecycleBinTableResponseBody) *RemoveRecycleBinTableResponse
	GetBody() *RemoveRecycleBinTableResponseBody
}

type RemoveRecycleBinTableResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveRecycleBinTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveRecycleBinTableResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveRecycleBinTableResponse) GoString() string {
	return s.String()
}

func (s *RemoveRecycleBinTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveRecycleBinTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveRecycleBinTableResponse) GetBody() *RemoveRecycleBinTableResponseBody {
	return s.Body
}

func (s *RemoveRecycleBinTableResponse) SetHeaders(v map[string]*string) *RemoveRecycleBinTableResponse {
	s.Headers = v
	return s
}

func (s *RemoveRecycleBinTableResponse) SetStatusCode(v int32) *RemoveRecycleBinTableResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveRecycleBinTableResponse) SetBody(v *RemoveRecycleBinTableResponseBody) *RemoveRecycleBinTableResponse {
	s.Body = v
	return s
}

func (s *RemoveRecycleBinTableResponse) Validate() error {
	return dara.Validate(s)
}
