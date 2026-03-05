// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMonthAmountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneType(v string) *GetMonthAmountRequest
	GetSceneType() *string
}

type GetMonthAmountRequest struct {
	// This parameter is required.
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
}

func (s GetMonthAmountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMonthAmountRequest) GoString() string {
	return s.String()
}

func (s *GetMonthAmountRequest) GetSceneType() *string {
	return s.SceneType
}

func (s *GetMonthAmountRequest) SetSceneType(v string) *GetMonthAmountRequest {
	s.SceneType = &v
	return s
}

func (s *GetMonthAmountRequest) Validate() error {
	return dara.Validate(s)
}
