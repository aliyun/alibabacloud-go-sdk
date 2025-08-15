// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGlobalEventsStorageRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGlobalEventsStorageRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGlobalEventsStorageRegionResponse
	GetStatusCode() *int32
	SetBody(v *GetGlobalEventsStorageRegionResponseBody) *GetGlobalEventsStorageRegionResponse
	GetBody() *GetGlobalEventsStorageRegionResponseBody
}

type GetGlobalEventsStorageRegionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGlobalEventsStorageRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGlobalEventsStorageRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGlobalEventsStorageRegionResponse) GoString() string {
	return s.String()
}

func (s *GetGlobalEventsStorageRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGlobalEventsStorageRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGlobalEventsStorageRegionResponse) GetBody() *GetGlobalEventsStorageRegionResponseBody {
	return s.Body
}

func (s *GetGlobalEventsStorageRegionResponse) SetHeaders(v map[string]*string) *GetGlobalEventsStorageRegionResponse {
	s.Headers = v
	return s
}

func (s *GetGlobalEventsStorageRegionResponse) SetStatusCode(v int32) *GetGlobalEventsStorageRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGlobalEventsStorageRegionResponse) SetBody(v *GetGlobalEventsStorageRegionResponseBody) *GetGlobalEventsStorageRegionResponse {
	s.Body = v
	return s
}

func (s *GetGlobalEventsStorageRegionResponse) Validate() error {
	return dara.Validate(s)
}
