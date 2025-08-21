// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJiangSuTelecomDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v string) *GetJiangSuTelecomDataRequest
	GetDate() *string
}

type GetJiangSuTelecomDataRequest struct {
	// example:
	//
	// 2024-11-09
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetJiangSuTelecomDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJiangSuTelecomDataRequest) GoString() string {
	return s.String()
}

func (s *GetJiangSuTelecomDataRequest) GetDate() *string {
	return s.Date
}

func (s *GetJiangSuTelecomDataRequest) SetDate(v string) *GetJiangSuTelecomDataRequest {
	s.Date = &v
	return s
}

func (s *GetJiangSuTelecomDataRequest) Validate() error {
	return dara.Validate(s)
}
