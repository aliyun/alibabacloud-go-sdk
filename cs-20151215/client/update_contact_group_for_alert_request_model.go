// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactGroupForAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertRuleGroupName(v string) *UpdateContactGroupForAlertRequest
	GetAlertRuleGroupName() *string
	SetContactGroupIds(v []*int64) *UpdateContactGroupForAlertRequest
	GetContactGroupIds() []*int64
	SetCrName(v string) *UpdateContactGroupForAlertRequest
	GetCrName() *string
	SetNamespace(v string) *UpdateContactGroupForAlertRequest
	GetNamespace() *string
}

type UpdateContactGroupForAlertRequest struct {
	// The name of the alert contact group.
	//
	// example:
	//
	// sample
	AlertRuleGroupName *string `json:"alert_rule_group_name,omitempty" xml:"alert_rule_group_name,omitempty"`
	// The list of contact group IDs.
	ContactGroupIds []*int64 `json:"contact_group_ids,omitempty" xml:"contact_group_ids,omitempty" type:"Repeated"`
	// The name of the container registry instance.
	//
	// example:
	//
	// sample
	CrName *string `json:"cr_name,omitempty" xml:"cr_name,omitempty"`
	// The namespace in which the resource resides.
	//
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s UpdateContactGroupForAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactGroupForAlertRequest) GoString() string {
	return s.String()
}

func (s *UpdateContactGroupForAlertRequest) GetAlertRuleGroupName() *string {
	return s.AlertRuleGroupName
}

func (s *UpdateContactGroupForAlertRequest) GetContactGroupIds() []*int64 {
	return s.ContactGroupIds
}

func (s *UpdateContactGroupForAlertRequest) GetCrName() *string {
	return s.CrName
}

func (s *UpdateContactGroupForAlertRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateContactGroupForAlertRequest) SetAlertRuleGroupName(v string) *UpdateContactGroupForAlertRequest {
	s.AlertRuleGroupName = &v
	return s
}

func (s *UpdateContactGroupForAlertRequest) SetContactGroupIds(v []*int64) *UpdateContactGroupForAlertRequest {
	s.ContactGroupIds = v
	return s
}

func (s *UpdateContactGroupForAlertRequest) SetCrName(v string) *UpdateContactGroupForAlertRequest {
	s.CrName = &v
	return s
}

func (s *UpdateContactGroupForAlertRequest) SetNamespace(v string) *UpdateContactGroupForAlertRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateContactGroupForAlertRequest) Validate() error {
	return dara.Validate(s)
}
