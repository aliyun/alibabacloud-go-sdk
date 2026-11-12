// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateServiceRequest
	GetClientToken() *string
	SetDevelop(v string) *CreateServiceRequest
	GetDevelop() *string
	SetDryRun(v bool) *CreateServiceRequest
	GetDryRun() *bool
	SetLabels(v map[string]*string) *CreateServiceRequest
	GetLabels() map[string]*string
	SetWorkspaceId(v string) *CreateServiceRequest
	GetWorkspaceId() *string
	SetBody(v string) *CreateServiceRequest
	GetBody() *string
}

type CreateServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 5db780a0-9d3e-4f35-a84a-xxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enter development mode.
	//
	// example:
	//
	// true
	Develop *string `json:"Develop,omitempty" xml:"Develop,omitempty"`
	// Specifies whether to only perform a dry run for this request.
	//
	// - true: A check request is sent, and no actual operation is performed. If the check fails, the corresponding error is returned. If the check succeeds, the error code DryRunOperation is returned.
	//
	// - false (default): A normal request is sent.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The user-defined labels.
	//
	// example:
	//
	// {"key":"value"}
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123456
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The request body. For instructions on key request parameter settings, see **Table 1. Request body parameters*	- and **Table 2. Metadata parameters*	- below. For more details about parameter settings, see [All related metric description for service models](https://help.aliyun.com/document_detail/450525.html).
	//
	// example:
	//
	// Service deployed by using an image:
	//
	// {
	//
	//   "name": "foo",
	//
	//   "metadata": {
	//
	//     "instance": 2,
	//
	//     "memory": 7000,
	//
	//     "cpu": 4
	//
	//     },
	//
	//   "containers": [
	//
	//     {
	//
	//       "image": "****",
	//
	//       "script": "***	- --listen=0.0.0.0 --server_port=8000 --headless",
	//
	//       "port": 8000
	//
	//     }
	//
	//   ],
	//
	//   "storage": [
	//
	//     {
	//
	//       "oss": {
	//
	//         "path": "oss://examplebuket/data111/",
	//
	//         "readOnly": false
	//
	//       },
	//
	//       "properties": {
	//
	//         "resource_type": "model"
	//
	//       },
	//
	//       "mount_path": "/data"
	//
	//     }
	//
	//   ]
	//
	// }
	//
	// AI-Web application deployed by using an image:
	//
	// {
	//
	//   "name": "foo",
	//
	//   "metadata": {
	//
	//     "instance": 1,
	//
	//     "memory": 7000,
	//
	//     "cpu": 4,
	//
	//     "enable_webservice": true
	//
	//   },
	//
	//   "containers": [
	//
	//     {
	//
	//       "image": "****",
	//
	//       "script": "***	- --listen=0.0.0.0 --server_port=8000 --headless",
	//
	//       "port": 8000
	//
	//     }
	//
	//   ],
	//
	//   "storage": [
	//
	//     {
	//
	//       "oss": {
	//
	//         "path": "oss://examplebucket/data111/",
	//
	//         "readOnly": false
	//
	//       },
	//
	//       "properties": {
	//
	//       "resource_type": "model"
	//
	//       },
	//
	//       "mount_path": "/data"
	//
	//     }
	//
	//   ]
	//
	// }
	//
	// Service deployed by using a model and a processor:
	//
	// {
	//
	//   "metadata": {
	//
	//     "instance": 1,
	//
	//     "memory": 7000,
	//
	//     "cpu": 4
	//
	//   },
	//
	//   "name": "foo",
	//
	//   "model_config": {},
	//
	//   "processor_type": "python",
	//
	//   "processor_path": "oss://****",
	//
	//   "processor_entry": "a.py",
	//
	//   "model_path": "oss://****"
	//
	// }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServiceRequest) GetDevelop() *string {
	return s.Develop
}

func (s *CreateServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateServiceRequest) GetLabels() map[string]*string {
	return s.Labels
}

func (s *CreateServiceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateServiceRequest) GetBody() *string {
	return s.Body
}

func (s *CreateServiceRequest) SetClientToken(v string) *CreateServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceRequest) SetDevelop(v string) *CreateServiceRequest {
	s.Develop = &v
	return s
}

func (s *CreateServiceRequest) SetDryRun(v bool) *CreateServiceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceRequest) SetLabels(v map[string]*string) *CreateServiceRequest {
	s.Labels = v
	return s
}

func (s *CreateServiceRequest) SetWorkspaceId(v string) *CreateServiceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateServiceRequest) SetBody(v string) *CreateServiceRequest {
	s.Body = &v
	return s
}

func (s *CreateServiceRequest) Validate() error {
	return dara.Validate(s)
}
