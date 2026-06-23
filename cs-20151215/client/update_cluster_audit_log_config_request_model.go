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
	// Specifies whether to disable the cluster audit log feature. Valid values:
	//
	// - false: enables the audit log feature or updates the audit log configuration.
	//
	// - true: disables the audit log feature.
	//
	// example:
	//
	// false
	Disable *bool `json:"disable,omitempty" xml:"disable,omitempty"`
	// The [SLS Project](https://help.aliyun.com/document_detail/48873.html) that contains the [Logstore](https://help.aliyun.com/document_detail/48874.html) for cluster audit logs.
	//
	// - Default value: k8s-log-{clusterid}.
	//
	// - After you enable the cluster audit log feature, a Logstore for cluster audit logs is created in the specified SLS Project.
	//
	// - If you need to change the SLS Project after enabling the cluster audit log feature, use this parameter to specify a new SLS Project. Only ACK managed clusters support changing the SLS Project.
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
