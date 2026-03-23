// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNewJobOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *NewJobOrderRequest
	GetAppCode() *string
	SetAppName(v string) *NewJobOrderRequest
	GetAppName() *string
	SetCallbackUrl(v string) *NewJobOrderRequest
	GetCallbackUrl() *string
	SetChangeType(v string) *NewJobOrderRequest
	GetChangeType() *string
	SetClientSystem(v string) *NewJobOrderRequest
	GetClientSystem() *string
	SetClientUniqueId(v string) *NewJobOrderRequest
	GetClientUniqueId() *string
	SetCreator(v string) *NewJobOrderRequest
	GetCreator() *string
	SetParams(v map[string]interface{}) *NewJobOrderRequest
	GetParams() map[string]interface{}
	SetReferUrl(v string) *NewJobOrderRequest
	GetReferUrl() *string
	SetRequestId(v string) *NewJobOrderRequest
	GetRequestId() *string
	SetTitle(v string) *NewJobOrderRequest
	GetTitle() *string
}

type NewJobOrderRequest struct {
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
	Params   map[string]interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
	ReferUrl *string                `json:"ReferUrl,omitempty" xml:"ReferUrl,omitempty"`
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s NewJobOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s NewJobOrderRequest) GoString() string {
	return s.String()
}

func (s *NewJobOrderRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *NewJobOrderRequest) GetAppName() *string {
	return s.AppName
}

func (s *NewJobOrderRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *NewJobOrderRequest) GetChangeType() *string {
	return s.ChangeType
}

func (s *NewJobOrderRequest) GetClientSystem() *string {
	return s.ClientSystem
}

func (s *NewJobOrderRequest) GetClientUniqueId() *string {
	return s.ClientUniqueId
}

func (s *NewJobOrderRequest) GetCreator() *string {
	return s.Creator
}

func (s *NewJobOrderRequest) GetParams() map[string]interface{} {
	return s.Params
}

func (s *NewJobOrderRequest) GetReferUrl() *string {
	return s.ReferUrl
}

func (s *NewJobOrderRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *NewJobOrderRequest) GetTitle() *string {
	return s.Title
}

func (s *NewJobOrderRequest) SetAppCode(v string) *NewJobOrderRequest {
	s.AppCode = &v
	return s
}

func (s *NewJobOrderRequest) SetAppName(v string) *NewJobOrderRequest {
	s.AppName = &v
	return s
}

func (s *NewJobOrderRequest) SetCallbackUrl(v string) *NewJobOrderRequest {
	s.CallbackUrl = &v
	return s
}

func (s *NewJobOrderRequest) SetChangeType(v string) *NewJobOrderRequest {
	s.ChangeType = &v
	return s
}

func (s *NewJobOrderRequest) SetClientSystem(v string) *NewJobOrderRequest {
	s.ClientSystem = &v
	return s
}

func (s *NewJobOrderRequest) SetClientUniqueId(v string) *NewJobOrderRequest {
	s.ClientUniqueId = &v
	return s
}

func (s *NewJobOrderRequest) SetCreator(v string) *NewJobOrderRequest {
	s.Creator = &v
	return s
}

func (s *NewJobOrderRequest) SetParams(v map[string]interface{}) *NewJobOrderRequest {
	s.Params = v
	return s
}

func (s *NewJobOrderRequest) SetReferUrl(v string) *NewJobOrderRequest {
	s.ReferUrl = &v
	return s
}

func (s *NewJobOrderRequest) SetRequestId(v string) *NewJobOrderRequest {
	s.RequestId = &v
	return s
}

func (s *NewJobOrderRequest) SetTitle(v string) *NewJobOrderRequest {
	s.Title = &v
	return s
}

func (s *NewJobOrderRequest) Validate() error {
	return dara.Validate(s)
}
