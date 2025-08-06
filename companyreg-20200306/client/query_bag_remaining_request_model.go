// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBagRemainingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *QueryBagRemainingRequest
	GetBizType() *string
}

type QueryBagRemainingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// esp.hightech
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s QueryBagRemainingRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryBagRemainingRequest) GoString() string {
	return s.String()
}

func (s *QueryBagRemainingRequest) GetBizType() *string {
	return s.BizType
}

func (s *QueryBagRemainingRequest) SetBizType(v string) *QueryBagRemainingRequest {
	s.BizType = &v
	return s
}

func (s *QueryBagRemainingRequest) Validate() error {
	return dara.Validate(s)
}
