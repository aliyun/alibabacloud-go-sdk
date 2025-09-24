// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryName(v string) *DeleteConfigRequest
	GetCategoryName() *string
	SetLabels(v string) *DeleteConfigRequest
	GetLabels() *string
}

type DeleteConfigRequest struct {
	// The category of the configuration item. Valid values:
	//
	// 	- CommonResourceConfig
	//
	// 	- DLCAutoRecycle - DLCPriorityConfig
	//
	// 	- DSWPriorityConfig
	//
	// 	- QuotaMaximumDuration
	//
	// 	- CommonTagConfig
	//
	// example:
	//
	// CommonResourceConfig
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The filter conditions. Separate multiple conditions with commas (,). The conditions have an AND relationship.
	//
	// example:
	//
	// key1=value;key2=value2
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s DeleteConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigRequest) GetCategoryName() *string {
	return s.CategoryName
}

func (s *DeleteConfigRequest) GetLabels() *string {
	return s.Labels
}

func (s *DeleteConfigRequest) SetCategoryName(v string) *DeleteConfigRequest {
	s.CategoryName = &v
	return s
}

func (s *DeleteConfigRequest) SetLabels(v string) *DeleteConfigRequest {
	s.Labels = &v
	return s
}

func (s *DeleteConfigRequest) Validate() error {
	return dara.Validate(s)
}
