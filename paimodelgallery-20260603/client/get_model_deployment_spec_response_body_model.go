// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelDeploymentSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInferenceSpec(v map[string]interface{}) *GetModelDeploymentSpecResponseBody
	GetInferenceSpec() map[string]interface{}
	SetRequestId(v string) *GetModelDeploymentSpecResponseBody
	GetRequestId() *string
}

type GetModelDeploymentSpecResponseBody struct {
	// example:
	//
	// {
	//
	//     "containers": [
	//
	//       {
	//
	//         "image": "eas-registry-vpc.cn-hangzhou.cr.aliyuncs.com/pai-eas/sglang:v0.5.17",
	//
	//         "port": 8000,
	//
	//         "script": "python -m sglang.launch_server **	- --port 8000"
	//
	//       }
	//
	//     ],
	//
	//     "metadata": {
	//
	//       "cpu": 248,
	//
	//       "disk": 850,
	//
	//       "gpu": 8,
	//
	//       "instance": 1,
	//
	//       "memory": 2744000,
	//
	//       "shm_size": 512
	//
	//     }
	//
	//   }
	InferenceSpec map[string]interface{} `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	// example:
	//
	// B6B54325-C98C-5937-87A3-2F96C07652EC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetModelDeploymentSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelDeploymentSpecResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelDeploymentSpecResponseBody) GetInferenceSpec() map[string]interface{} {
	return s.InferenceSpec
}

func (s *GetModelDeploymentSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelDeploymentSpecResponseBody) SetInferenceSpec(v map[string]interface{}) *GetModelDeploymentSpecResponseBody {
	s.InferenceSpec = v
	return s
}

func (s *GetModelDeploymentSpecResponseBody) SetRequestId(v string) *GetModelDeploymentSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelDeploymentSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
