// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyObjectGroupOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyObjectGroupOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyObjectGroupOperationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyObjectGroupOperationResponseBody) *ModifyObjectGroupOperationResponse
	GetBody() *ModifyObjectGroupOperationResponseBody
}

type ModifyObjectGroupOperationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyObjectGroupOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyObjectGroupOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyObjectGroupOperationResponse) GoString() string {
	return s.String()
}

func (s *ModifyObjectGroupOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyObjectGroupOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyObjectGroupOperationResponse) GetBody() *ModifyObjectGroupOperationResponseBody {
	return s.Body
}

func (s *ModifyObjectGroupOperationResponse) SetHeaders(v map[string]*string) *ModifyObjectGroupOperationResponse {
	s.Headers = v
	return s
}

func (s *ModifyObjectGroupOperationResponse) SetStatusCode(v int32) *ModifyObjectGroupOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyObjectGroupOperationResponse) SetBody(v *ModifyObjectGroupOperationResponseBody) *ModifyObjectGroupOperationResponse {
	s.Body = v
	return s
}

func (s *ModifyObjectGroupOperationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
