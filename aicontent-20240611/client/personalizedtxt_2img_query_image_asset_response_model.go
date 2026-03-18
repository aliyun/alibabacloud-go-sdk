// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgQueryImageAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Personalizedtxt2imgQueryImageAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Personalizedtxt2imgQueryImageAssetResponse
	GetStatusCode() *int32
	SetBody(v interface{}) *Personalizedtxt2imgQueryImageAssetResponse
	GetBody() interface{}
}

type Personalizedtxt2imgQueryImageAssetResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       interface{}        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Personalizedtxt2imgQueryImageAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgQueryImageAssetResponse) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryImageAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Personalizedtxt2imgQueryImageAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Personalizedtxt2imgQueryImageAssetResponse) GetBody() interface{} {
	return s.Body
}

func (s *Personalizedtxt2imgQueryImageAssetResponse) SetHeaders(v map[string]*string) *Personalizedtxt2imgQueryImageAssetResponse {
	s.Headers = v
	return s
}

func (s *Personalizedtxt2imgQueryImageAssetResponse) SetStatusCode(v int32) *Personalizedtxt2imgQueryImageAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryImageAssetResponse) SetBody(v interface{}) *Personalizedtxt2imgQueryImageAssetResponse {
	s.Body = v
	return s
}

func (s *Personalizedtxt2imgQueryImageAssetResponse) Validate() error {
	return dara.Validate(s)
}
