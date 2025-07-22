// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPartitionsHeatmapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPartitionsHeatmapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPartitionsHeatmapResponse
	GetStatusCode() *int32
	SetBody(v *GetPartitionsHeatmapResponseBody) *GetPartitionsHeatmapResponse
	GetBody() *GetPartitionsHeatmapResponseBody
}

type GetPartitionsHeatmapResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPartitionsHeatmapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPartitionsHeatmapResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPartitionsHeatmapResponse) GoString() string {
	return s.String()
}

func (s *GetPartitionsHeatmapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPartitionsHeatmapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPartitionsHeatmapResponse) GetBody() *GetPartitionsHeatmapResponseBody {
	return s.Body
}

func (s *GetPartitionsHeatmapResponse) SetHeaders(v map[string]*string) *GetPartitionsHeatmapResponse {
	s.Headers = v
	return s
}

func (s *GetPartitionsHeatmapResponse) SetStatusCode(v int32) *GetPartitionsHeatmapResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPartitionsHeatmapResponse) SetBody(v *GetPartitionsHeatmapResponseBody) *GetPartitionsHeatmapResponse {
	s.Body = v
	return s
}

func (s *GetPartitionsHeatmapResponse) Validate() error {
	return dara.Validate(s)
}
