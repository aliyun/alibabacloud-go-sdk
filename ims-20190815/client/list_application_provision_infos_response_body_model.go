// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationProvisionInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationProvisionInfos(v *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos) *ListApplicationProvisionInfosResponseBody
	GetApplicationProvisionInfos() *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos
	SetRequestId(v string) *ListApplicationProvisionInfosResponseBody
	GetRequestId() *string
}

type ListApplicationProvisionInfosResponseBody struct {
	ApplicationProvisionInfos *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos `json:"ApplicationProvisionInfos,omitempty" xml:"ApplicationProvisionInfos,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E403EBFD-C997-489D-BFC7-37C05E66D67C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApplicationProvisionInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationProvisionInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponseBody) GetApplicationProvisionInfos() *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos {
	return s.ApplicationProvisionInfos
}

func (s *ListApplicationProvisionInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationProvisionInfosResponseBody) SetApplicationProvisionInfos(v *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos) *ListApplicationProvisionInfosResponseBody {
	s.ApplicationProvisionInfos = v
	return s
}

func (s *ListApplicationProvisionInfosResponseBody) SetRequestId(v string) *ListApplicationProvisionInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBody) Validate() error {
	if s.ApplicationProvisionInfos != nil {
		if err := s.ApplicationProvisionInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos struct {
	ApplicationProvisionInfo []*ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo `json:"ApplicationProvisionInfo,omitempty" xml:"ApplicationProvisionInfo,omitempty" type:"Repeated"`
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos) GetApplicationProvisionInfo() []*ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	return s.ApplicationProvisionInfo
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos) SetApplicationProvisionInfo(v []*ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos {
	s.ApplicationProvisionInfo = v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos) Validate() error {
	if s.ApplicationProvisionInfo != nil {
		for _, item := range s.ApplicationProvisionInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo struct {
	AccountId      *string                                                                                                   `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AppId          *string                                                                                                   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName        *string                                                                                                   `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateDate     *string                                                                                                   `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DelegatedScope *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	DisplayName    *string                                                                                                   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	UpdateDate     *string                                                                                                   `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) GetAccountId() *string {
	return s.AccountId
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) GetAppId() *string {
	return s.AppId
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) GetAppName() *string {
	return s.AppName
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) GetDelegatedScope() *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope {
	return s.DelegatedScope
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetAccountId(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.AccountId = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetAppId(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.AppId = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetAppName(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.AppName = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetCreateDate(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.CreateDate = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetDelegatedScope(v *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.DelegatedScope = v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetDisplayName(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.DisplayName = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetUpdateDate(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.UpdateDate = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) Validate() error {
	if s.DelegatedScope != nil {
		if err := s.DelegatedScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope struct {
	PredefinedScopes *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope) GetPredefinedScopes() *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes {
	return s.PredefinedScopes
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope) SetPredefinedScopes(v *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope {
	s.PredefinedScopes = v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope) Validate() error {
	if s.PredefinedScopes != nil {
		if err := s.PredefinedScopes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes struct {
	PredefinedScope []*ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes) GetPredefinedScope() []*ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope {
	return s.PredefinedScope
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes) SetPredefinedScope(v []*ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes) Validate() error {
	if s.PredefinedScope != nil {
		for _, item := range s.PredefinedScope {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) GetName() *string {
	return s.Name
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) Validate() error {
	return dara.Validate(s)
}
