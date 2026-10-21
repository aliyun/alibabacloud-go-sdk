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
	// The pricing request body, including popCode, popVersion, apiName, and parameters. All calls return a unified pricing envelope. The language parameter specifies the server-side message language, such as en or zh. Target resource operations are not executed for cost-irrelevant or unsupported pricing queries.
	//
	// example:
	//
	// {"popCode":"ecs","popVersion":"2014-05-26","apiName":"CreateDisk","parameters":{"RegionId":"cn-hangzhou","ZoneId":"cn-hangzhou-i","DiskCategory":"cloud_essd","Size":40},"language":"zh"}
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
