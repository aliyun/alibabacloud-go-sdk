// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSyntheticMonitorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSyntheticMonitorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSyntheticMonitorsResponse
	GetStatusCode() *int32
	SetBody(v *GetSyntheticMonitorsResponseBody) *GetSyntheticMonitorsResponse
	GetBody() *GetSyntheticMonitorsResponseBody
}

type GetSyntheticMonitorsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSyntheticMonitorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSyntheticMonitorsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticMonitorsResponse) GoString() string {
	return s.String()
}

func (s *GetSyntheticMonitorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSyntheticMonitorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSyntheticMonitorsResponse) GetBody() *GetSyntheticMonitorsResponseBody {
	return s.Body
}

func (s *GetSyntheticMonitorsResponse) SetHeaders(v map[string]*string) *GetSyntheticMonitorsResponse {
	s.Headers = v
	return s
}

func (s *GetSyntheticMonitorsResponse) SetStatusCode(v int32) *GetSyntheticMonitorsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSyntheticMonitorsResponse) SetBody(v *GetSyntheticMonitorsResponseBody) *GetSyntheticMonitorsResponse {
	s.Body = v
	return s
}

func (s *GetSyntheticMonitorsResponse) Validate() error {
	return dara.Validate(s)
}
