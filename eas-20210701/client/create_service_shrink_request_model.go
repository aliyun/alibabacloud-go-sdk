// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateServiceShrinkRequest
	GetClientToken() *string
	SetDevelop(v string) *CreateServiceShrinkRequest
	GetDevelop() *string
	SetDryRun(v bool) *CreateServiceShrinkRequest
	GetDryRun() *bool
	SetLabelsShrink(v string) *CreateServiceShrinkRequest
	GetLabelsShrink() *string
	SetWorkspaceId(v string) *CreateServiceShrinkRequest
	GetWorkspaceId() *string
	SetBody(v string) *CreateServiceShrinkRequest
	GetBody() *string
}

type CreateServiceShrinkRequest struct {
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
	LabelsShrink *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
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

func (s CreateServiceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServiceShrinkRequest) GetDevelop() *string {
	return s.Develop
}

func (s *CreateServiceShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateServiceShrinkRequest) GetLabelsShrink() *string {
	return s.LabelsShrink
}

func (s *CreateServiceShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateServiceShrinkRequest) GetBody() *string {
	return s.Body
}

func (s *CreateServiceShrinkRequest) SetClientToken(v string) *CreateServiceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetDevelop(v string) *CreateServiceShrinkRequest {
	s.Develop = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetDryRun(v bool) *CreateServiceShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetLabelsShrink(v string) *CreateServiceShrinkRequest {
	s.LabelsShrink = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetWorkspaceId(v string) *CreateServiceShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetBody(v string) *CreateServiceShrinkRequest {
	s.Body = &v
	return s
}

func (s *CreateServiceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
