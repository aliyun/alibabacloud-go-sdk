// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceType(v string) *GetBasicStatisticsRequest
	GetSourceType() *string
}

type GetBasicStatisticsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s GetBasicStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBasicStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetBasicStatisticsRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *GetBasicStatisticsRequest) SetSourceType(v string) *GetBasicStatisticsRequest {
	s.SourceType = &v
	return s
}

func (s *GetBasicStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
