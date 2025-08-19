// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMonitorDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *GetMonitorDataRequest
	GetBody() *string
}

type GetMonitorDataRequest struct {
	// example:
	//
	// {"start":1689245180581,"end":1689246950582,"queries":[{"metric":"aliyunes.elasticsearch.index.docs.count","aggregator":"sum","downsample":"avg","tags":{"resource":"{appName}"},"filters":[],"granularity":"auto"}]}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMonitorDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *GetMonitorDataRequest) GetBody() *string {
	return s.Body
}

func (s *GetMonitorDataRequest) SetBody(v string) *GetMonitorDataRequest {
	s.Body = &v
	return s
}

func (s *GetMonitorDataRequest) Validate() error {
	return dara.Validate(s)
}
