// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGlobalEventsStorageRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGlobalEventsStorageRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGlobalEventsStorageRegionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGlobalEventsStorageRegionResponseBody) *UpdateGlobalEventsStorageRegionResponse
	GetBody() *UpdateGlobalEventsStorageRegionResponseBody
}

type UpdateGlobalEventsStorageRegionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGlobalEventsStorageRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGlobalEventsStorageRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGlobalEventsStorageRegionResponse) GoString() string {
	return s.String()
}

func (s *UpdateGlobalEventsStorageRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGlobalEventsStorageRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGlobalEventsStorageRegionResponse) GetBody() *UpdateGlobalEventsStorageRegionResponseBody {
	return s.Body
}

func (s *UpdateGlobalEventsStorageRegionResponse) SetHeaders(v map[string]*string) *UpdateGlobalEventsStorageRegionResponse {
	s.Headers = v
	return s
}

func (s *UpdateGlobalEventsStorageRegionResponse) SetStatusCode(v int32) *UpdateGlobalEventsStorageRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGlobalEventsStorageRegionResponse) SetBody(v *UpdateGlobalEventsStorageRegionResponseBody) *UpdateGlobalEventsStorageRegionResponse {
	s.Body = v
	return s
}

func (s *UpdateGlobalEventsStorageRegionResponse) Validate() error {
	return dara.Validate(s)
}
