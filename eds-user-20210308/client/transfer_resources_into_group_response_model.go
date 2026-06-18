// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferResourcesIntoGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferResourcesIntoGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferResourcesIntoGroupResponse
	GetStatusCode() *int32
	SetBody(v *TransferResourcesIntoGroupResponseBody) *TransferResourcesIntoGroupResponse
	GetBody() *TransferResourcesIntoGroupResponseBody
}

type TransferResourcesIntoGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferResourcesIntoGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferResourcesIntoGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferResourcesIntoGroupResponse) GoString() string {
	return s.String()
}

func (s *TransferResourcesIntoGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferResourcesIntoGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferResourcesIntoGroupResponse) GetBody() *TransferResourcesIntoGroupResponseBody {
	return s.Body
}

func (s *TransferResourcesIntoGroupResponse) SetHeaders(v map[string]*string) *TransferResourcesIntoGroupResponse {
	s.Headers = v
	return s
}

func (s *TransferResourcesIntoGroupResponse) SetStatusCode(v int32) *TransferResourcesIntoGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferResourcesIntoGroupResponse) SetBody(v *TransferResourcesIntoGroupResponseBody) *TransferResourcesIntoGroupResponse {
	s.Body = v
	return s
}

func (s *TransferResourcesIntoGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
