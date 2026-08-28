// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerPatrolHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TriggerPatrolHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *TriggerPatrolHeaders
	GetWorkspace() *string
}

type TriggerPatrolHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s TriggerPatrolHeaders) String() string {
	return dara.Prettify(s)
}

func (s TriggerPatrolHeaders) GoString() string {
	return s.String()
}

func (s *TriggerPatrolHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TriggerPatrolHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *TriggerPatrolHeaders) SetCommonHeaders(v map[string]*string) *TriggerPatrolHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TriggerPatrolHeaders) SetWorkspace(v string) *TriggerPatrolHeaders {
	s.Workspace = &v
	return s
}

func (s *TriggerPatrolHeaders) Validate() error {
	return dara.Validate(s)
}
