// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileTypeStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileTypeStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileTypeStatisticResponse
	GetStatusCode() *int32
	SetBody(v *GetFileTypeStatisticResponseBody) *GetFileTypeStatisticResponse
	GetBody() *GetFileTypeStatisticResponseBody
}

type GetFileTypeStatisticResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileTypeStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileTypeStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileTypeStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetFileTypeStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileTypeStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileTypeStatisticResponse) GetBody() *GetFileTypeStatisticResponseBody {
	return s.Body
}

func (s *GetFileTypeStatisticResponse) SetHeaders(v map[string]*string) *GetFileTypeStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetFileTypeStatisticResponse) SetStatusCode(v int32) *GetFileTypeStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileTypeStatisticResponse) SetBody(v *GetFileTypeStatisticResponseBody) *GetFileTypeStatisticResponse {
	s.Body = v
	return s
}

func (s *GetFileTypeStatisticResponse) Validate() error {
	return dara.Validate(s)
}
