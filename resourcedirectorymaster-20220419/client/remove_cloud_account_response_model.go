// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveCloudAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveCloudAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveCloudAccountResponse
	GetStatusCode() *int32
	SetBody(v *RemoveCloudAccountResponseBody) *RemoveCloudAccountResponse
	GetBody() *RemoveCloudAccountResponseBody
}

type RemoveCloudAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveCloudAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveCloudAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveCloudAccountResponse) GoString() string {
	return s.String()
}

func (s *RemoveCloudAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveCloudAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveCloudAccountResponse) GetBody() *RemoveCloudAccountResponseBody {
	return s.Body
}

func (s *RemoveCloudAccountResponse) SetHeaders(v map[string]*string) *RemoveCloudAccountResponse {
	s.Headers = v
	return s
}

func (s *RemoveCloudAccountResponse) SetStatusCode(v int32) *RemoveCloudAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveCloudAccountResponse) SetBody(v *RemoveCloudAccountResponseBody) *RemoveCloudAccountResponse {
	s.Body = v
	return s
}

func (s *RemoveCloudAccountResponse) Validate() error {
	return dara.Validate(s)
}
