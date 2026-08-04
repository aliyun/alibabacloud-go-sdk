// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPAL7ConfigRewriteOp interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *PAL7ConfigRewriteOp
	GetKey() *string
	SetOldValue(v string) *PAL7ConfigRewriteOp
	GetOldValue() *string
	SetOp(v string) *PAL7ConfigRewriteOp
	GetOp() *string
	SetValue(v string) *PAL7ConfigRewriteOp
	GetValue() *string
	SetValueVariable(v string) *PAL7ConfigRewriteOp
	GetValueVariable() *string
}

type PAL7ConfigRewriteOp struct {
	// HTTP header or query parameter name. Required.
	//
	// example:
	//
	// X-Test-Param
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Value to find and replace. Used only for the replace operation.
	//
	// example:
	//
	// old_value
	OldValue *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
	// Operation type. Required.
	//
	// - **add**: Add an item.
	//
	// - **set**: Set a value.
	//
	// - **delete**: Delete an item.
	//
	// - **replace**: Replace a value.
	//
	// example:
	//
	// add
	Op *string `json:"Op,omitempty" xml:"Op,omitempty"`
	// Target value as a string.
	//
	// example:
	//
	// new_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// Target value as a string. Valid values:
	//
	// - **sase_app_name**: Application name.
	//
	// - **sase_app_id**: Application ID.
	//
	// - **sase_policy_name**: Zero Trust policy name.
	//
	// - **sase_user_username**: Username.
	//
	// - **sase_user_department**: User department.
	//
	// - **sase_user_group_infos**: User organizational structure information.
	//
	// - **sase_user_matched_user_groups**: User group information.
	//
	// - **sase_client_addr**: Client address.
	//
	// - **sase_client_ip**: Client IP address.
	//
	// - **sase_client_port**: Client port.
	//
	// example:
	//
	// sase_app_name
	ValueVariable *string `json:"ValueVariable,omitempty" xml:"ValueVariable,omitempty"`
}

func (s PAL7ConfigRewriteOp) String() string {
	return dara.Prettify(s)
}

func (s PAL7ConfigRewriteOp) GoString() string {
	return s.String()
}

func (s *PAL7ConfigRewriteOp) GetKey() *string {
	return s.Key
}

func (s *PAL7ConfigRewriteOp) GetOldValue() *string {
	return s.OldValue
}

func (s *PAL7ConfigRewriteOp) GetOp() *string {
	return s.Op
}

func (s *PAL7ConfigRewriteOp) GetValue() *string {
	return s.Value
}

func (s *PAL7ConfigRewriteOp) GetValueVariable() *string {
	return s.ValueVariable
}

func (s *PAL7ConfigRewriteOp) SetKey(v string) *PAL7ConfigRewriteOp {
	s.Key = &v
	return s
}

func (s *PAL7ConfigRewriteOp) SetOldValue(v string) *PAL7ConfigRewriteOp {
	s.OldValue = &v
	return s
}

func (s *PAL7ConfigRewriteOp) SetOp(v string) *PAL7ConfigRewriteOp {
	s.Op = &v
	return s
}

func (s *PAL7ConfigRewriteOp) SetValue(v string) *PAL7ConfigRewriteOp {
	s.Value = &v
	return s
}

func (s *PAL7ConfigRewriteOp) SetValueVariable(v string) *PAL7ConfigRewriteOp {
	s.ValueVariable = &v
	return s
}

func (s *PAL7ConfigRewriteOp) Validate() error {
	return dara.Validate(s)
}
