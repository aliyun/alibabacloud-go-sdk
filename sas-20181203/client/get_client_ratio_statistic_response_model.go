// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientRatioStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClientRatioStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClientRatioStatisticResponse
	GetStatusCode() *int32
	SetBody(v *GetClientRatioStatisticResponseBody) *GetClientRatioStatisticResponse
	GetBody() *GetClientRatioStatisticResponseBody
}

type GetClientRatioStatisticResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClientRatioStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClientRatioStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClientRatioStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetClientRatioStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClientRatioStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClientRatioStatisticResponse) GetBody() *GetClientRatioStatisticResponseBody {
	return s.Body
}

func (s *GetClientRatioStatisticResponse) SetHeaders(v map[string]*string) *GetClientRatioStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetClientRatioStatisticResponse) SetStatusCode(v int32) *GetClientRatioStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClientRatioStatisticResponse) SetBody(v *GetClientRatioStatisticResponseBody) *GetClientRatioStatisticResponse {
	s.Body = v
	return s
}

func (s *GetClientRatioStatisticResponse) Validate() error {
	return dara.Validate(s)
}
