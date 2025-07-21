// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserSuppressionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveUserSuppressionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveUserSuppressionResponse
	GetStatusCode() *int32
	SetBody(v *RemoveUserSuppressionResponseBody) *RemoveUserSuppressionResponse
	GetBody() *RemoveUserSuppressionResponseBody
}

type RemoveUserSuppressionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserSuppressionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUserSuppressionResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserSuppressionResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserSuppressionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveUserSuppressionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveUserSuppressionResponse) GetBody() *RemoveUserSuppressionResponseBody {
	return s.Body
}

func (s *RemoveUserSuppressionResponse) SetHeaders(v map[string]*string) *RemoveUserSuppressionResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserSuppressionResponse) SetStatusCode(v int32) *RemoveUserSuppressionResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserSuppressionResponse) SetBody(v *RemoveUserSuppressionResponseBody) *RemoveUserSuppressionResponse {
	s.Body = v
	return s
}

func (s *RemoveUserSuppressionResponse) Validate() error {
	return dara.Validate(s)
}
