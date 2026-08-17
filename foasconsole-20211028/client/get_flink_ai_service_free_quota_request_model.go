// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlinkAiServiceFreeQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegion(v string) *GetFlinkAiServiceFreeQuotaRequest
	GetRegion() *string
}

type GetFlinkAiServiceFreeQuotaRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetFlinkAiServiceFreeQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFlinkAiServiceFreeQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetFlinkAiServiceFreeQuotaRequest) GetRegion() *string {
	return s.Region
}

func (s *GetFlinkAiServiceFreeQuotaRequest) SetRegion(v string) *GetFlinkAiServiceFreeQuotaRequest {
	s.Region = &v
	return s
}

func (s *GetFlinkAiServiceFreeQuotaRequest) Validate() error {
	return dara.Validate(s)
}
