// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMiningTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMiningTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMiningTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetMiningTaskResultResponseBody) *GetMiningTaskResultResponse
	GetBody() *GetMiningTaskResultResponseBody
}

type GetMiningTaskResultResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMiningTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMiningTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMiningTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetMiningTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMiningTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMiningTaskResultResponse) GetBody() *GetMiningTaskResultResponseBody {
	return s.Body
}

func (s *GetMiningTaskResultResponse) SetHeaders(v map[string]*string) *GetMiningTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetMiningTaskResultResponse) SetStatusCode(v int32) *GetMiningTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMiningTaskResultResponse) SetBody(v *GetMiningTaskResultResponseBody) *GetMiningTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetMiningTaskResultResponse) Validate() error {
	return dara.Validate(s)
}
