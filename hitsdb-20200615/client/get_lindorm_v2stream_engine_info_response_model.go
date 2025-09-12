// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2StreamEngineInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLindormV2StreamEngineInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLindormV2StreamEngineInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetLindormV2StreamEngineInfoResponseBody) *GetLindormV2StreamEngineInfoResponse
	GetBody() *GetLindormV2StreamEngineInfoResponseBody
}

type GetLindormV2StreamEngineInfoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLindormV2StreamEngineInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLindormV2StreamEngineInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2StreamEngineInfoResponse) GoString() string {
	return s.String()
}

func (s *GetLindormV2StreamEngineInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLindormV2StreamEngineInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLindormV2StreamEngineInfoResponse) GetBody() *GetLindormV2StreamEngineInfoResponseBody {
	return s.Body
}

func (s *GetLindormV2StreamEngineInfoResponse) SetHeaders(v map[string]*string) *GetLindormV2StreamEngineInfoResponse {
	s.Headers = v
	return s
}

func (s *GetLindormV2StreamEngineInfoResponse) SetStatusCode(v int32) *GetLindormV2StreamEngineInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLindormV2StreamEngineInfoResponse) SetBody(v *GetLindormV2StreamEngineInfoResponseBody) *GetLindormV2StreamEngineInfoResponse {
	s.Body = v
	return s
}

func (s *GetLindormV2StreamEngineInfoResponse) Validate() error {
	return dara.Validate(s)
}
