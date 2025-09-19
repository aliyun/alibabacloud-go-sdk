// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthVersionStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuthVersionStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuthVersionStatisticResponse
	GetStatusCode() *int32
	SetBody(v *GetAuthVersionStatisticResponseBody) *GetAuthVersionStatisticResponse
	GetBody() *GetAuthVersionStatisticResponseBody
}

type GetAuthVersionStatisticResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuthVersionStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuthVersionStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuthVersionStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetAuthVersionStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuthVersionStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuthVersionStatisticResponse) GetBody() *GetAuthVersionStatisticResponseBody {
	return s.Body
}

func (s *GetAuthVersionStatisticResponse) SetHeaders(v map[string]*string) *GetAuthVersionStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetAuthVersionStatisticResponse) SetStatusCode(v int32) *GetAuthVersionStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthVersionStatisticResponse) SetBody(v *GetAuthVersionStatisticResponseBody) *GetAuthVersionStatisticResponse {
	s.Body = v
	return s
}

func (s *GetAuthVersionStatisticResponse) Validate() error {
	return dara.Validate(s)
}
