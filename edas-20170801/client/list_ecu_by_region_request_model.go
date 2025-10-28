// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEcuByRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAct(v string) *ListEcuByRegionRequest
	GetAct() *string
	SetLogicalRegionId(v string) *ListEcuByRegionRequest
	GetLogicalRegionId() *string
}

type ListEcuByRegionRequest struct {
	// Set the value to `pop-query`.
	//
	// This parameter is required.
	//
	// example:
	//
	// pop-query
	Act *string `json:"Act,omitempty" xml:"Act,omitempty"`
	// The ID of the namespace.
	//
	// 	- The ID of a custom namespace is in the `region ID:namespace identifier` format. Example: cn-beijing:tdy218.
	//
	// 	- The ID of the default namespace is in the `region ID` format. Example: cn-beijing.
	//
	// example:
	//
	// cn-beijing or cn-beijing:tdy218
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
}

func (s ListEcuByRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEcuByRegionRequest) GoString() string {
	return s.String()
}

func (s *ListEcuByRegionRequest) GetAct() *string {
	return s.Act
}

func (s *ListEcuByRegionRequest) GetLogicalRegionId() *string {
	return s.LogicalRegionId
}

func (s *ListEcuByRegionRequest) SetAct(v string) *ListEcuByRegionRequest {
	s.Act = &v
	return s
}

func (s *ListEcuByRegionRequest) SetLogicalRegionId(v string) *ListEcuByRegionRequest {
	s.LogicalRegionId = &v
	return s
}

func (s *ListEcuByRegionRequest) Validate() error {
	return dara.Validate(s)
}
