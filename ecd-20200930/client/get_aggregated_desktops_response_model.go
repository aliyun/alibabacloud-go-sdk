// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregatedDesktopsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregatedDesktopsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregatedDesktopsResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregatedDesktopsResponseBody) *GetAggregatedDesktopsResponse
	GetBody() *GetAggregatedDesktopsResponseBody
}

type GetAggregatedDesktopsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregatedDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregatedDesktopsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregatedDesktopsResponse) GoString() string {
	return s.String()
}

func (s *GetAggregatedDesktopsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregatedDesktopsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregatedDesktopsResponse) GetBody() *GetAggregatedDesktopsResponseBody {
	return s.Body
}

func (s *GetAggregatedDesktopsResponse) SetHeaders(v map[string]*string) *GetAggregatedDesktopsResponse {
	s.Headers = v
	return s
}

func (s *GetAggregatedDesktopsResponse) SetStatusCode(v int32) *GetAggregatedDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregatedDesktopsResponse) SetBody(v *GetAggregatedDesktopsResponseBody) *GetAggregatedDesktopsResponse {
	s.Body = v
	return s
}

func (s *GetAggregatedDesktopsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
