// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterAuditProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuditEnabled(v bool) *GetClusterAuditProjectResponseBody
	GetAuditEnabled() *bool
	SetSlsProjectName(v string) *GetClusterAuditProjectResponseBody
	GetSlsProjectName() *string
}

type GetClusterAuditProjectResponseBody struct {
	// Indicates whether the cluster auditing feature is enabled for the cluster.
	//
	// 	- `true`: The cluster auditing feature is enabled for the cluster.
	//
	// 	- `false`: The cluster auditing feature is disabled for the cluster.
	//
	// example:
	//
	// true
	AuditEnabled *bool `json:"audit_enabled,omitempty" xml:"audit_enabled,omitempty"`
	// The SLS project in which the audit logs of the API server are stored.
	//
	// example:
	//
	// k8s-log-cad1230511cbb4db4a488e58518******
	SlsProjectName *string `json:"sls_project_name,omitempty" xml:"sls_project_name,omitempty"`
}

func (s GetClusterAuditProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterAuditProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterAuditProjectResponseBody) GetAuditEnabled() *bool {
	return s.AuditEnabled
}

func (s *GetClusterAuditProjectResponseBody) GetSlsProjectName() *string {
	return s.SlsProjectName
}

func (s *GetClusterAuditProjectResponseBody) SetAuditEnabled(v bool) *GetClusterAuditProjectResponseBody {
	s.AuditEnabled = &v
	return s
}

func (s *GetClusterAuditProjectResponseBody) SetSlsProjectName(v string) *GetClusterAuditProjectResponseBody {
	s.SlsProjectName = &v
	return s
}

func (s *GetClusterAuditProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
