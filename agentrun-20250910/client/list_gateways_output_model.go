// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewaysOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *Gateway) *ListGatewaysOutput
	GetItems() *Gateway
	SetPageNumber(v int32) *ListGatewaysOutput
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewaysOutput
	GetPageSize() *int32
	SetTotal(v int32) *ListGatewaysOutput
	GetTotal() *int32
}

type ListGatewaysOutput struct {
	Items      *Gateway `json:"items,omitempty" xml:"items,omitempty"`
	PageNumber *int32   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *int32   `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListGatewaysOutput) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaysOutput) GoString() string {
	return s.String()
}

func (s *ListGatewaysOutput) GetItems() *Gateway {
	return s.Items
}

func (s *ListGatewaysOutput) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewaysOutput) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewaysOutput) GetTotal() *int32 {
	return s.Total
}

func (s *ListGatewaysOutput) SetItems(v *Gateway) *ListGatewaysOutput {
	s.Items = v
	return s
}

func (s *ListGatewaysOutput) SetPageNumber(v int32) *ListGatewaysOutput {
	s.PageNumber = &v
	return s
}

func (s *ListGatewaysOutput) SetPageSize(v int32) *ListGatewaysOutput {
	s.PageSize = &v
	return s
}

func (s *ListGatewaysOutput) SetTotal(v int32) *ListGatewaysOutput {
	s.Total = &v
	return s
}

func (s *ListGatewaysOutput) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}
