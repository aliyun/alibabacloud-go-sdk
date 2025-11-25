// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchSlsDispatchStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeBatchSlsDispatchStatusRequest
	GetLang() *string
}

type DescribeBatchSlsDispatchStatusRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeBatchSlsDispatchStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchSlsDispatchStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeBatchSlsDispatchStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeBatchSlsDispatchStatusRequest) SetLang(v string) *DescribeBatchSlsDispatchStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusRequest) Validate() error {
	return dara.Validate(s)
}
