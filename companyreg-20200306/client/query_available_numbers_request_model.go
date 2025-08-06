// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvailableNumbersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *QueryAvailableNumbersRequest
	GetBizType() *string
}

type QueryAvailableNumbersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s QueryAvailableNumbersRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailableNumbersRequest) GoString() string {
	return s.String()
}

func (s *QueryAvailableNumbersRequest) GetBizType() *string {
	return s.BizType
}

func (s *QueryAvailableNumbersRequest) SetBizType(v string) *QueryAvailableNumbersRequest {
	s.BizType = &v
	return s
}

func (s *QueryAvailableNumbersRequest) Validate() error {
	return dara.Validate(s)
}
