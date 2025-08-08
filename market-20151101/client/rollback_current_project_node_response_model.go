// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackCurrentProjectNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackCurrentProjectNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackCurrentProjectNodeResponse
	GetStatusCode() *int32
	SetBody(v *RollbackCurrentProjectNodeResponseBody) *RollbackCurrentProjectNodeResponse
	GetBody() *RollbackCurrentProjectNodeResponseBody
}

type RollbackCurrentProjectNodeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackCurrentProjectNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackCurrentProjectNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackCurrentProjectNodeResponse) GoString() string {
	return s.String()
}

func (s *RollbackCurrentProjectNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackCurrentProjectNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackCurrentProjectNodeResponse) GetBody() *RollbackCurrentProjectNodeResponseBody {
	return s.Body
}

func (s *RollbackCurrentProjectNodeResponse) SetHeaders(v map[string]*string) *RollbackCurrentProjectNodeResponse {
	s.Headers = v
	return s
}

func (s *RollbackCurrentProjectNodeResponse) SetStatusCode(v int32) *RollbackCurrentProjectNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackCurrentProjectNodeResponse) SetBody(v *RollbackCurrentProjectNodeResponseBody) *RollbackCurrentProjectNodeResponse {
	s.Body = v
	return s
}

func (s *RollbackCurrentProjectNodeResponse) Validate() error {
	return dara.Validate(s)
}
