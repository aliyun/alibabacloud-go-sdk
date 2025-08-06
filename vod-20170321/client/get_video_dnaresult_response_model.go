// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoDNAResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoDNAResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoDNAResultResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoDNAResultResponseBody) *GetVideoDNAResultResponse
	GetBody() *GetVideoDNAResultResponseBody
}

type GetVideoDNAResultResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoDNAResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoDNAResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDNAResultResponse) GoString() string {
	return s.String()
}

func (s *GetVideoDNAResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoDNAResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoDNAResultResponse) GetBody() *GetVideoDNAResultResponseBody {
	return s.Body
}

func (s *GetVideoDNAResultResponse) SetHeaders(v map[string]*string) *GetVideoDNAResultResponse {
	s.Headers = v
	return s
}

func (s *GetVideoDNAResultResponse) SetStatusCode(v int32) *GetVideoDNAResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoDNAResultResponse) SetBody(v *GetVideoDNAResultResponseBody) *GetVideoDNAResultResponse {
	s.Body = v
	return s
}

func (s *GetVideoDNAResultResponse) Validate() error {
	return dara.Validate(s)
}
