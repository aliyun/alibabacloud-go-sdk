// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsDebugSampleLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPtsDebugSampleLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPtsDebugSampleLogsResponse
	GetStatusCode() *int32
	SetBody(v *GetPtsDebugSampleLogsResponseBody) *GetPtsDebugSampleLogsResponse
	GetBody() *GetPtsDebugSampleLogsResponseBody
}

type GetPtsDebugSampleLogsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPtsDebugSampleLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPtsDebugSampleLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPtsDebugSampleLogsResponse) GoString() string {
	return s.String()
}

func (s *GetPtsDebugSampleLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPtsDebugSampleLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPtsDebugSampleLogsResponse) GetBody() *GetPtsDebugSampleLogsResponseBody {
	return s.Body
}

func (s *GetPtsDebugSampleLogsResponse) SetHeaders(v map[string]*string) *GetPtsDebugSampleLogsResponse {
	s.Headers = v
	return s
}

func (s *GetPtsDebugSampleLogsResponse) SetStatusCode(v int32) *GetPtsDebugSampleLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponse) SetBody(v *GetPtsDebugSampleLogsResponseBody) *GetPtsDebugSampleLogsResponse {
	s.Body = v
	return s
}

func (s *GetPtsDebugSampleLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
