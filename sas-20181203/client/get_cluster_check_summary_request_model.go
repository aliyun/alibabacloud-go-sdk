// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterCheckSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetClusterCheckSummaryRequest
	GetClusterId() *string
}

type GetClusterCheckSummaryRequest struct {
	// ID of the queried cluster
	//
	// This parameter is required.
	//
	// example:
	//
	// c3aaf6c8085f84791882eef200cd2****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetClusterCheckSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClusterCheckSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetClusterCheckSummaryRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterCheckSummaryRequest) SetClusterId(v string) *GetClusterCheckSummaryRequest {
	s.ClusterId = &v
	return s
}

func (s *GetClusterCheckSummaryRequest) Validate() error {
	return dara.Validate(s)
}
