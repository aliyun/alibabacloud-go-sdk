// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReviewRecordResult interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *ReviewRecordResult
	GetAddress() *string
	SetApplyType(v string) *ReviewRecordResult
	GetApplyType() *string
	SetContactName(v string) *ReviewRecordResult
	GetContactName() *string
	SetPhone(v string) *ReviewRecordResult
	GetPhone() *string
	SetSceneDesc(v string) *ReviewRecordResult
	GetSceneDesc() *string
	SetScopes(v []*string) *ReviewRecordResult
	GetScopes() []*string
	SetServiceType(v string) *ReviewRecordResult
	GetServiceType() *string
}

type ReviewRecordResult struct {
	Address     *string   `json:"address,omitempty" xml:"address,omitempty"`
	ApplyType   *string   `json:"applyType,omitempty" xml:"applyType,omitempty"`
	ContactName *string   `json:"contactName,omitempty" xml:"contactName,omitempty"`
	Phone       *string   `json:"phone,omitempty" xml:"phone,omitempty"`
	SceneDesc   *string   `json:"sceneDesc,omitempty" xml:"sceneDesc,omitempty"`
	Scopes      []*string `json:"scopes,omitempty" xml:"scopes,omitempty" type:"Repeated"`
	ServiceType *string   `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s ReviewRecordResult) String() string {
	return dara.Prettify(s)
}

func (s ReviewRecordResult) GoString() string {
	return s.String()
}

func (s *ReviewRecordResult) GetAddress() *string {
	return s.Address
}

func (s *ReviewRecordResult) GetApplyType() *string {
	return s.ApplyType
}

func (s *ReviewRecordResult) GetContactName() *string {
	return s.ContactName
}

func (s *ReviewRecordResult) GetPhone() *string {
	return s.Phone
}

func (s *ReviewRecordResult) GetSceneDesc() *string {
	return s.SceneDesc
}

func (s *ReviewRecordResult) GetScopes() []*string {
	return s.Scopes
}

func (s *ReviewRecordResult) GetServiceType() *string {
	return s.ServiceType
}

func (s *ReviewRecordResult) SetAddress(v string) *ReviewRecordResult {
	s.Address = &v
	return s
}

func (s *ReviewRecordResult) SetApplyType(v string) *ReviewRecordResult {
	s.ApplyType = &v
	return s
}

func (s *ReviewRecordResult) SetContactName(v string) *ReviewRecordResult {
	s.ContactName = &v
	return s
}

func (s *ReviewRecordResult) SetPhone(v string) *ReviewRecordResult {
	s.Phone = &v
	return s
}

func (s *ReviewRecordResult) SetSceneDesc(v string) *ReviewRecordResult {
	s.SceneDesc = &v
	return s
}

func (s *ReviewRecordResult) SetScopes(v []*string) *ReviewRecordResult {
	s.Scopes = v
	return s
}

func (s *ReviewRecordResult) SetServiceType(v string) *ReviewRecordResult {
	s.ServiceType = &v
	return s
}

func (s *ReviewRecordResult) Validate() error {
	return dara.Validate(s)
}
