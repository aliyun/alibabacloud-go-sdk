// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterTDEResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterTDEResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterTDEResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterTDEResponseBody) *ModifyDBClusterTDEResponse
	GetBody() *ModifyDBClusterTDEResponseBody
}

type ModifyDBClusterTDEResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterTDEResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterTDEResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterTDEResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterTDEResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterTDEResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterTDEResponse) GetBody() *ModifyDBClusterTDEResponseBody {
	return s.Body
}

func (s *ModifyDBClusterTDEResponse) SetHeaders(v map[string]*string) *ModifyDBClusterTDEResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterTDEResponse) SetStatusCode(v int32) *ModifyDBClusterTDEResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterTDEResponse) SetBody(v *ModifyDBClusterTDEResponseBody) *ModifyDBClusterTDEResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterTDEResponse) Validate() error {
	return dara.Validate(s)
}
