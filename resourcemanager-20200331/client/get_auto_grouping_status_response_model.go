// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoGroupingStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAutoGroupingStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAutoGroupingStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetAutoGroupingStatusResponseBody) *GetAutoGroupingStatusResponse
	GetBody() *GetAutoGroupingStatusResponseBody
}

type GetAutoGroupingStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutoGroupingStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutoGroupingStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAutoGroupingStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAutoGroupingStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAutoGroupingStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAutoGroupingStatusResponse) GetBody() *GetAutoGroupingStatusResponseBody {
	return s.Body
}

func (s *GetAutoGroupingStatusResponse) SetHeaders(v map[string]*string) *GetAutoGroupingStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAutoGroupingStatusResponse) SetStatusCode(v int32) *GetAutoGroupingStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoGroupingStatusResponse) SetBody(v *GetAutoGroupingStatusResponseBody) *GetAutoGroupingStatusResponse {
	s.Body = v
	return s
}

func (s *GetAutoGroupingStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
