// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBizUnitInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBizUnitInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBizUnitInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetBizUnitInfoResponseBody) *GetBizUnitInfoResponse
	GetBody() *GetBizUnitInfoResponseBody
}

type GetBizUnitInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBizUnitInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBizUnitInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBizUnitInfoResponse) GoString() string {
	return s.String()
}

func (s *GetBizUnitInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBizUnitInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBizUnitInfoResponse) GetBody() *GetBizUnitInfoResponseBody {
	return s.Body
}

func (s *GetBizUnitInfoResponse) SetHeaders(v map[string]*string) *GetBizUnitInfoResponse {
	s.Headers = v
	return s
}

func (s *GetBizUnitInfoResponse) SetStatusCode(v int32) *GetBizUnitInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBizUnitInfoResponse) SetBody(v *GetBizUnitInfoResponseBody) *GetBizUnitInfoResponse {
	s.Body = v
	return s
}

func (s *GetBizUnitInfoResponse) Validate() error {
	return dara.Validate(s)
}
