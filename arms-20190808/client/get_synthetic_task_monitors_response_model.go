// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSyntheticTaskMonitorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSyntheticTaskMonitorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSyntheticTaskMonitorsResponse
	GetStatusCode() *int32
	SetBody(v *GetSyntheticTaskMonitorsResponseBody) *GetSyntheticTaskMonitorsResponse
	GetBody() *GetSyntheticTaskMonitorsResponseBody
}

type GetSyntheticTaskMonitorsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSyntheticTaskMonitorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSyntheticTaskMonitorsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskMonitorsResponse) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskMonitorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSyntheticTaskMonitorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSyntheticTaskMonitorsResponse) GetBody() *GetSyntheticTaskMonitorsResponseBody {
	return s.Body
}

func (s *GetSyntheticTaskMonitorsResponse) SetHeaders(v map[string]*string) *GetSyntheticTaskMonitorsResponse {
	s.Headers = v
	return s
}

func (s *GetSyntheticTaskMonitorsResponse) SetStatusCode(v int32) *GetSyntheticTaskMonitorsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponse) SetBody(v *GetSyntheticTaskMonitorsResponseBody) *GetSyntheticTaskMonitorsResponse {
	s.Body = v
	return s
}

func (s *GetSyntheticTaskMonitorsResponse) Validate() error {
	return dara.Validate(s)
}
