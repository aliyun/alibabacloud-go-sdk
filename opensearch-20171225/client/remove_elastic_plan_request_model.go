// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveElasticPlanRequest interface {
	dara.Model
	String() string
	GoString() string
}

type RemoveElasticPlanRequest struct {
}

func (s RemoveElasticPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *RemoveElasticPlanRequest) Validate() error {
	return dara.Validate(s)
}
