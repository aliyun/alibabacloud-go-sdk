// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStructSyncOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStructSyncOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStructSyncOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateStructSyncOrderResponseBody) *CreateStructSyncOrderResponse
	GetBody() *CreateStructSyncOrderResponseBody
}

type CreateStructSyncOrderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStructSyncOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStructSyncOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStructSyncOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateStructSyncOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStructSyncOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStructSyncOrderResponse) GetBody() *CreateStructSyncOrderResponseBody {
	return s.Body
}

func (s *CreateStructSyncOrderResponse) SetHeaders(v map[string]*string) *CreateStructSyncOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateStructSyncOrderResponse) SetStatusCode(v int32) *CreateStructSyncOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStructSyncOrderResponse) SetBody(v *CreateStructSyncOrderResponseBody) *CreateStructSyncOrderResponse {
	s.Body = v
	return s
}

func (s *CreateStructSyncOrderResponse) Validate() error {
	return dara.Validate(s)
}
