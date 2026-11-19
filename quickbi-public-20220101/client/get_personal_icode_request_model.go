// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPersonalICodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQuickBIUserId(v string) *GetPersonalICodeRequest
	GetQuickBIUserId() *string
}

type GetPersonalICodeRequest struct {
	// The Quick BI user ID.
	//
	// example:
	//
	// asdasda
	QuickBIUserId *string `json:"QuickBIUserId,omitempty" xml:"QuickBIUserId,omitempty"`
}

func (s GetPersonalICodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPersonalICodeRequest) GoString() string {
	return s.String()
}

func (s *GetPersonalICodeRequest) GetQuickBIUserId() *string {
	return s.QuickBIUserId
}

func (s *GetPersonalICodeRequest) SetQuickBIUserId(v string) *GetPersonalICodeRequest {
	s.QuickBIUserId = &v
	return s
}

func (s *GetPersonalICodeRequest) Validate() error {
	return dara.Validate(s)
}
