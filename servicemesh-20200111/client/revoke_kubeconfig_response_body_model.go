// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeKubeconfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKubeconfig(v string) *RevokeKubeconfigResponseBody
	GetKubeconfig() *string
	SetRequestId(v string) *RevokeKubeconfigResponseBody
	GetRequestId() *string
}

type RevokeKubeconfigResponseBody struct {
	// The new kubeconfig file generated.
	//
	// example:
	//
	// apiVersion: v1 clusters: - cluster:     server: https://121.**.**.**:6443     certificate-authority-data: ****	- name: kubernetes contexts: - context:     cluster: kubernetes     user: "*****"   name: ****	- current-context: ****	- kind: Config preferences: {} users: - name: "*****"   user:     client-certificate-data: ****	- client-key-data: *****
	Kubeconfig *string `json:"Kubeconfig,omitempty" xml:"Kubeconfig,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7CF71C8B-79DD-150F-929E-267C51C5E311
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeKubeconfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeKubeconfigResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeKubeconfigResponseBody) GetKubeconfig() *string {
	return s.Kubeconfig
}

func (s *RevokeKubeconfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeKubeconfigResponseBody) SetKubeconfig(v string) *RevokeKubeconfigResponseBody {
	s.Kubeconfig = &v
	return s
}

func (s *RevokeKubeconfigResponseBody) SetRequestId(v string) *RevokeKubeconfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeKubeconfigResponseBody) Validate() error {
	return dara.Validate(s)
}
