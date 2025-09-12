// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeLindormV2StreamEngineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeLindormV2StreamEngineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeLindormV2StreamEngineResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeLindormV2StreamEngineResponseBody) *UpgradeLindormV2StreamEngineResponse
	GetBody() *UpgradeLindormV2StreamEngineResponseBody
}

type UpgradeLindormV2StreamEngineResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeLindormV2StreamEngineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeLindormV2StreamEngineResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeLindormV2StreamEngineResponse) GoString() string {
	return s.String()
}

func (s *UpgradeLindormV2StreamEngineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeLindormV2StreamEngineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeLindormV2StreamEngineResponse) GetBody() *UpgradeLindormV2StreamEngineResponseBody {
	return s.Body
}

func (s *UpgradeLindormV2StreamEngineResponse) SetHeaders(v map[string]*string) *UpgradeLindormV2StreamEngineResponse {
	s.Headers = v
	return s
}

func (s *UpgradeLindormV2StreamEngineResponse) SetStatusCode(v int32) *UpgradeLindormV2StreamEngineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineResponse) SetBody(v *UpgradeLindormV2StreamEngineResponseBody) *UpgradeLindormV2StreamEngineResponse {
	s.Body = v
	return s
}

func (s *UpgradeLindormV2StreamEngineResponse) Validate() error {
	return dara.Validate(s)
}
