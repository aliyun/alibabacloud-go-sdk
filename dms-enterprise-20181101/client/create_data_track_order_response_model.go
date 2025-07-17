// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataTrackOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataTrackOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataTrackOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataTrackOrderResponseBody) *CreateDataTrackOrderResponse
	GetBody() *CreateDataTrackOrderResponseBody
}

type CreateDataTrackOrderResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataTrackOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataTrackOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataTrackOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateDataTrackOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataTrackOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataTrackOrderResponse) GetBody() *CreateDataTrackOrderResponseBody {
	return s.Body
}

func (s *CreateDataTrackOrderResponse) SetHeaders(v map[string]*string) *CreateDataTrackOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateDataTrackOrderResponse) SetStatusCode(v int32) *CreateDataTrackOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataTrackOrderResponse) SetBody(v *CreateDataTrackOrderResponseBody) *CreateDataTrackOrderResponse {
	s.Body = v
	return s
}

func (s *CreateDataTrackOrderResponse) Validate() error {
	return dara.Validate(s)
}
