// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHealthPercentageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHealthPercentageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHealthPercentageResponse
	GetStatusCode() *int32
	SetBody(v *GetHealthPercentageResponseBody) *GetHealthPercentageResponse
	GetBody() *GetHealthPercentageResponseBody
}

type GetHealthPercentageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHealthPercentageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHealthPercentageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHealthPercentageResponse) GoString() string {
	return s.String()
}

func (s *GetHealthPercentageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHealthPercentageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHealthPercentageResponse) GetBody() *GetHealthPercentageResponseBody {
	return s.Body
}

func (s *GetHealthPercentageResponse) SetHeaders(v map[string]*string) *GetHealthPercentageResponse {
	s.Headers = v
	return s
}

func (s *GetHealthPercentageResponse) SetStatusCode(v int32) *GetHealthPercentageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHealthPercentageResponse) SetBody(v *GetHealthPercentageResponseBody) *GetHealthPercentageResponse {
	s.Body = v
	return s
}

func (s *GetHealthPercentageResponse) Validate() error {
	return dara.Validate(s)
}
