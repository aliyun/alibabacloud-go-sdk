// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceStatusStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceStatusStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceStatusStatisticResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceStatusStatisticResponseBody) *GetInstanceStatusStatisticResponse
	GetBody() *GetInstanceStatusStatisticResponseBody
}

type GetInstanceStatusStatisticResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceStatusStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceStatusStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceStatusStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceStatusStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceStatusStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceStatusStatisticResponse) GetBody() *GetInstanceStatusStatisticResponseBody {
	return s.Body
}

func (s *GetInstanceStatusStatisticResponse) SetHeaders(v map[string]*string) *GetInstanceStatusStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceStatusStatisticResponse) SetStatusCode(v int32) *GetInstanceStatusStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceStatusStatisticResponse) SetBody(v *GetInstanceStatusStatisticResponseBody) *GetInstanceStatusStatisticResponse {
	s.Body = v
	return s
}

func (s *GetInstanceStatusStatisticResponse) Validate() error {
	return dara.Validate(s)
}
