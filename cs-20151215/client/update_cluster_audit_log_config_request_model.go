// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterAuditLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisable(v bool) *UpdateClusterAuditLogConfigRequest
	GetDisable() *bool
	SetSlsProjectName(v string) *UpdateClusterAuditLogConfigRequest
	GetSlsProjectName() *string
}

type UpdateClusterAuditLogConfigRequest struct {
	// Enable or disable audit logging.
	//
	// 	- false: enables audit logging or updates the audit logging configurations.
	//
	// 	- true: disables audit logging.
	//
	// example:
	//
	// false
	Disable *bool `json:"disable,omitempty" xml:"disable,omitempty"`
	// The [Simple Log Service project](https://help.aliyun.com/document_detail/48873.html) to which the [Logstore](https://help.aliyun.com/document_detail/48873.html) storing the cluster audit logs belongs.
	//
	// 	- Default value: k8s-log-{clusterid}.
	//
	// 	- After the cluster audit log feature is enabled, a Logstore is created in the specified Simple Log Service project to store cluster audit logs.
	//
	// 	- If you want to change the project after audit logging is enabled for the cluster, you can use this parameter to specify another project. You can perform this operation only in ACK managed clusters.
	//
	// example:
	//
	// k8s-log-c82e6987e2961451182edacd74faf****
	SlsProjectName *string `json:"sls_project_name,omitempty" xml:"sls_project_name,omitempty"`
}

func (s UpdateClusterAuditLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterAuditLogConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterAuditLogConfigRequest) GetDisable() *bool {
	return s.Disable
}

func (s *UpdateClusterAuditLogConfigRequest) GetSlsProjectName() *string {
	return s.SlsProjectName
}

func (s *UpdateClusterAuditLogConfigRequest) SetDisable(v bool) *UpdateClusterAuditLogConfigRequest {
	s.Disable = &v
	return s
}

func (s *UpdateClusterAuditLogConfigRequest) SetSlsProjectName(v string) *UpdateClusterAuditLogConfigRequest {
	s.SlsProjectName = &v
	return s
}

func (s *UpdateClusterAuditLogConfigRequest) Validate() error {
	return dara.Validate(s)
}
