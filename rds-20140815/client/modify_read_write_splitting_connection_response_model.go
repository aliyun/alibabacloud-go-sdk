// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReadWriteSplittingConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyReadWriteSplittingConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyReadWriteSplittingConnectionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyReadWriteSplittingConnectionResponseBody) *ModifyReadWriteSplittingConnectionResponse
	GetBody() *ModifyReadWriteSplittingConnectionResponseBody
}

type ModifyReadWriteSplittingConnectionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyReadWriteSplittingConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyReadWriteSplittingConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyReadWriteSplittingConnectionResponse) GoString() string {
	return s.String()
}

func (s *ModifyReadWriteSplittingConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyReadWriteSplittingConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyReadWriteSplittingConnectionResponse) GetBody() *ModifyReadWriteSplittingConnectionResponseBody {
	return s.Body
}

func (s *ModifyReadWriteSplittingConnectionResponse) SetHeaders(v map[string]*string) *ModifyReadWriteSplittingConnectionResponse {
	s.Headers = v
	return s
}

func (s *ModifyReadWriteSplittingConnectionResponse) SetStatusCode(v int32) *ModifyReadWriteSplittingConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionResponse) SetBody(v *ModifyReadWriteSplittingConnectionResponseBody) *ModifyReadWriteSplittingConnectionResponse {
	s.Body = v
	return s
}

func (s *ModifyReadWriteSplittingConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
