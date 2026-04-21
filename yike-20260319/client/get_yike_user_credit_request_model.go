// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeUserCreditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetYikeUserId(v string) *GetYikeUserCreditRequest
	GetYikeUserId() *string
}

type GetYikeUserCreditRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// id
	YikeUserId *string `json:"YikeUserId,omitempty" xml:"YikeUserId,omitempty"`
}

func (s GetYikeUserCreditRequest) String() string {
	return dara.Prettify(s)
}

func (s GetYikeUserCreditRequest) GoString() string {
	return s.String()
}

func (s *GetYikeUserCreditRequest) GetYikeUserId() *string {
	return s.YikeUserId
}

func (s *GetYikeUserCreditRequest) SetYikeUserId(v string) *GetYikeUserCreditRequest {
	s.YikeUserId = &v
	return s
}

func (s *GetYikeUserCreditRequest) Validate() error {
	return dara.Validate(s)
}
