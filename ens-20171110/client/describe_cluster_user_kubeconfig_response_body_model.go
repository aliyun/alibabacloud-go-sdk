// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterUserKubeconfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v interface{}) *DescribeClusterUserKubeconfigResponseBody
	GetConfig() interface{}
	SetRequestId(v string) *DescribeClusterUserKubeconfigResponseBody
	GetRequestId() *string
}

type DescribeClusterUserKubeconfigResponseBody struct {
	// example:
	//
	// apiVersion: v1
	//
	// clusters:
	//
	// - cluster:
	//
	//     certificate-authority-data: x
	//
	//     server: https://111.111.111.111:6443
	//
	//   name: kubernetes
	//
	// contexts:
	//
	// - context:
	//
	//     cluster: kubernetes
	//
	//     user: user
	//
	//   name: eck-xxxxx
	//
	// current-context: eck-xxxx
	//
	// kind: Config
	//
	// preferences: {}
	//
	// users:
	//
	// - name: user
	//
	//   user:
	//
	//     client-certificate-data: x
	//
	//     client-key-data: x
	Config interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterUserKubeconfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterUserKubeconfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterUserKubeconfigResponseBody) GetConfig() interface{} {
	return s.Config
}

func (s *DescribeClusterUserKubeconfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterUserKubeconfigResponseBody) SetConfig(v interface{}) *DescribeClusterUserKubeconfigResponseBody {
	s.Config = v
	return s
}

func (s *DescribeClusterUserKubeconfigResponseBody) SetRequestId(v string) *DescribeClusterUserKubeconfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterUserKubeconfigResponseBody) Validate() error {
	return dara.Validate(s)
}
