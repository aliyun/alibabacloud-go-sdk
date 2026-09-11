// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v map[string]interface{}) *GetApiPriceRequest
	GetBody() map[string]interface{}
}

type GetApiPriceRequest struct {
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApiPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApiPriceRequest) GoString() string {
	return s.String()
}

func (s *GetApiPriceRequest) GetBody() map[string]interface{} {
	return s.Body
}

func (s *GetApiPriceRequest) SetBody(v map[string]interface{}) *GetApiPriceRequest {
	s.Body = v
	return s
}

func (s *GetApiPriceRequest) Validate() error {
	return dara.Validate(s)
}
