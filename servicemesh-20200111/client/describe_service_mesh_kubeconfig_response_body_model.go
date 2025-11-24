// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshKubeconfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExpireTime(v string) *DescribeServiceMeshKubeconfigResponseBody
	GetExpireTime() *string
	SetKubeconfig(v string) *DescribeServiceMeshKubeconfigResponseBody
	GetKubeconfig() *string
	SetRequestId(v string) *DescribeServiceMeshKubeconfigResponseBody
	GetRequestId() *string
}

type DescribeServiceMeshKubeconfigResponseBody struct {
	// The expiration time of the kubeconfig certificate. The format is: YYYY-MM-DD hh: mm: ss.
	//
	// example:
	//
	// 2024-05-28 16:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The content of the kubeconfig file of the cluster.
	//
	// example:
	//
	// apiVersion: v1 clusters: - cluster:     server: https://47.110.xx.xx:6443     certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURUakNDQWphZ0F3SUJBZ0lVYzBQVy82ejR1aHlxYkRRdnNsV1htSmpJeFdNd0RRWUpLb1pJaHZjTkFRRUwKQlFBd1BqRW5NQThHQTFVRUNoTUlhR0Z1WjNwb2IzVXdGQVlEVlFRS0V3MWhiR2xpWVdKaElHTnNiM1ZrTVJNdwpFUVlEVlFRREV3cHJkV0psY201bGRHVnpNQ0FYRFRJeU1EUXdOekExTVRnd01Gb1lEekl3TlRJd016TXdNRFV4Ck9EQXdXakErTVNjd0R3WURWUVFLRXdob1lXNW5lbWh2ZFRBVUJnTlZCQW9URFdGc2FXSmhZbUVnWTJ4dmRXUXgKRXpBUkJnTlZCQU1UQ210MVltVnlibVYwWlhNd2dnRWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUJE****
	Kubeconfig *string `json:"Kubeconfig,omitempty" xml:"Kubeconfig,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceMeshKubeconfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshKubeconfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshKubeconfigResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeServiceMeshKubeconfigResponseBody) GetKubeconfig() *string {
	return s.Kubeconfig
}

func (s *DescribeServiceMeshKubeconfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceMeshKubeconfigResponseBody) SetExpireTime(v string) *DescribeServiceMeshKubeconfigResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeServiceMeshKubeconfigResponseBody) SetKubeconfig(v string) *DescribeServiceMeshKubeconfigResponseBody {
	s.Kubeconfig = &v
	return s
}

func (s *DescribeServiceMeshKubeconfigResponseBody) SetRequestId(v string) *DescribeServiceMeshKubeconfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceMeshKubeconfigResponseBody) Validate() error {
	return dara.Validate(s)
}
