// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmonMonitorDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *GetEmonMonitorDataRequest
	GetBody() *string
}

type GetEmonMonitorDataRequest struct {
	// The Grafana metric monitoring query conditions, including the query start time, metric expression, and metric query list.
	//
	// example:
	//
	// {"start":1522127381471,"queries":[{"metric":"elasticbuild.elasticsearch.source.total_doc_count","aggregator":"sum","downsample":"avg","tags":{"taskName":"et-testtsk","userId":"123456"},"granularity":"1m"}],"limit":"","end":1522129151000}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEmonMonitorDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEmonMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *GetEmonMonitorDataRequest) GetBody() *string {
	return s.Body
}

func (s *GetEmonMonitorDataRequest) SetBody(v string) *GetEmonMonitorDataRequest {
	s.Body = &v
	return s
}

func (s *GetEmonMonitorDataRequest) Validate() error {
	return dara.Validate(s)
}
