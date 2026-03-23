// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNewJobOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *NewJobOrderShrinkRequest
	GetAppCode() *string
	SetAppName(v string) *NewJobOrderShrinkRequest
	GetAppName() *string
	SetCallbackUrl(v string) *NewJobOrderShrinkRequest
	GetCallbackUrl() *string
	SetChangeType(v string) *NewJobOrderShrinkRequest
	GetChangeType() *string
	SetClientSystem(v string) *NewJobOrderShrinkRequest
	GetClientSystem() *string
	SetClientUniqueId(v string) *NewJobOrderShrinkRequest
	GetClientUniqueId() *string
	SetCreator(v string) *NewJobOrderShrinkRequest
	GetCreator() *string
	SetParamsShrink(v string) *NewJobOrderShrinkRequest
	GetParamsShrink() *string
	SetReferUrl(v string) *NewJobOrderShrinkRequest
	GetReferUrl() *string
	SetRequestId(v string) *NewJobOrderShrinkRequest
	GetRequestId() *string
	SetTitle(v string) *NewJobOrderShrinkRequest
	GetTitle() *string
}

type NewJobOrderShrinkRequest struct {
	AppCode     *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// This parameter is required.
	ChangeType *string `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// This parameter is required.
	ClientSystem *string `json:"ClientSystem,omitempty" xml:"ClientSystem,omitempty"`
	// This parameter is required.
	ClientUniqueId *string `json:"ClientUniqueId,omitempty" xml:"ClientUniqueId,omitempty"`
	Creator        *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// This parameter is required.
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
	ReferUrl     *string `json:"ReferUrl,omitempty" xml:"ReferUrl,omitempty"`
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s NewJobOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s NewJobOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *NewJobOrderShrinkRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *NewJobOrderShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *NewJobOrderShrinkRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *NewJobOrderShrinkRequest) GetChangeType() *string {
	return s.ChangeType
}

func (s *NewJobOrderShrinkRequest) GetClientSystem() *string {
	return s.ClientSystem
}

func (s *NewJobOrderShrinkRequest) GetClientUniqueId() *string {
	return s.ClientUniqueId
}

func (s *NewJobOrderShrinkRequest) GetCreator() *string {
	return s.Creator
}

func (s *NewJobOrderShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *NewJobOrderShrinkRequest) GetReferUrl() *string {
	return s.ReferUrl
}

func (s *NewJobOrderShrinkRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *NewJobOrderShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *NewJobOrderShrinkRequest) SetAppCode(v string) *NewJobOrderShrinkRequest {
	s.AppCode = &v
	return s
}

func (s *NewJobOrderShrinkRequest) SetAppName(v string) *NewJobOrderShrinkRequest {
	s.AppName = &v
	return s
}

func (s *NewJobOrderShrinkRequest) SetCallbackUrl(v string) *NewJobOrderShrinkRequest {
	s.CallbackUrl = &v
	return s
}

func (s *NewJobOrderShrinkRequest) SetChangeType(v string) *NewJobOrderShrinkRequest {
	s.ChangeType = &v
	return s
}

func (s *NewJobOrderShrinkRequest) SetClientSystem(v string) *NewJobOrderShrinkRequest {
	s.ClientSystem = &v
	return s
}

func (s *NewJobOrderShrinkRequest) SetClientUniqueId(v string) *NewJobOrderShrinkRequest {
	s.ClientUniqueId = &v
	return s
}

func (s *NewJobOrderShrinkRequest) SetCreator(v string) *NewJobOrderShrinkRequest {
	s.Creator = &v
	return s
}

func (s *NewJobOrderShrinkRequest) SetParamsShrink(v string) *NewJobOrderShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *NewJobOrderShrinkRequest) SetReferUrl(v string) *NewJobOrderShrinkRequest {
	s.ReferUrl = &v
	return s
}

func (s *NewJobOrderShrinkRequest) SetRequestId(v string) *NewJobOrderShrinkRequest {
	s.RequestId = &v
	return s
}

func (s *NewJobOrderShrinkRequest) SetTitle(v string) *NewJobOrderShrinkRequest {
	s.Title = &v
	return s
}

func (s *NewJobOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
