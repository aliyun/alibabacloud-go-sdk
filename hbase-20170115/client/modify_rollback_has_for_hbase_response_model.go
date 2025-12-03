// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRollbackHasForHbaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRollbackHasForHbaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRollbackHasForHbaseResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRollbackHasForHbaseResponseBody) *ModifyRollbackHasForHbaseResponse
	GetBody() *ModifyRollbackHasForHbaseResponseBody
}

type ModifyRollbackHasForHbaseResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRollbackHasForHbaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRollbackHasForHbaseResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRollbackHasForHbaseResponse) GoString() string {
	return s.String()
}

func (s *ModifyRollbackHasForHbaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRollbackHasForHbaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRollbackHasForHbaseResponse) GetBody() *ModifyRollbackHasForHbaseResponseBody {
	return s.Body
}

func (s *ModifyRollbackHasForHbaseResponse) SetHeaders(v map[string]*string) *ModifyRollbackHasForHbaseResponse {
	s.Headers = v
	return s
}

func (s *ModifyRollbackHasForHbaseResponse) SetStatusCode(v int32) *ModifyRollbackHasForHbaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRollbackHasForHbaseResponse) SetBody(v *ModifyRollbackHasForHbaseResponseBody) *ModifyRollbackHasForHbaseResponse {
	s.Body = v
	return s
}

func (s *ModifyRollbackHasForHbaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
